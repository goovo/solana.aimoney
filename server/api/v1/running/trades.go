package running

import (
    "os/exec"
    "syscall"
    "runtime"      // 新增：用于判断当前操作系统
    "strconv"      // 恢复：GetDryRun 中使用了 FormatUint/Atoi
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/running"
    runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
    "github.com/gin-gonic/gin"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    "go.uber.org/zap"
    "strings" // 使用：处理字符串判空
)

type TradesApi struct {}



// CreateTrades 创建交易报表(模拟盘)
// @Tags Trades
// @Summary 创建交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.Trades true "创建交易报表(模拟盘)"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /trades/createTrades [post]
func (tradesApi *TradesApi) CreateTrades(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var trades running.Trades
    err := c.ShouldBindJSON(&trades)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    err = tradesService.CreateTrades(ctx,&trades)
    if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}
// 为当前用户创建模拟交易机器人
func (tradesApi *TradesApi) CreateDryRun(c *gin.Context) {
    // 注意：此接口仅用于启动频繁交易（freqtrade）模拟盘进程，不涉及数据库事务
    uid := utils.GetUserID(c)  

    // 1. 根据操作系统选择工作目录
    var (
        workDir string
    )
    switch runtime.GOOS {
    case "darwin":
        workDir = "/Users/qinjialiang/Devs/ai/aibot"
    default:
        workDir = "/trader/freqtrade"
    }

    // 2. 从数据库查询该用户 status=1 的一条 sys_user_api 记录，仅选择 config_file 字段（为相对路径）
    type apiCfg struct {
        ConfigFile *string `gorm:"column:config_file"`
    }
    var rec apiCfg
    if err := global.GVA_DB.Model(&running.SysUserApi{}).
        Where("userId = ? AND status = 1", uid). // 注意：列名与模型标签一致，使用 userId
         Order("id desc").
         Select("config_file").
         First(&rec).Error; err != nil {
        // 若没有可用的API配置，则返回 pid=-1
        global.GVA_LOG.Warn("未找到可用的API配置，返回 pid=-1", zap.Uint("uid", uid), zap.Error(err))
        response.OkWithData(-1, c)
        return
    }
    var configPath string
    if rec.ConfigFile != nil {
        configPath = *rec.ConfigFile
    }
    if strings.TrimSpace(configPath) == "" {
        // 若记录存在但未生成配置文件路径，同样返回 pid=-1
        global.GVA_LOG.Warn("config_file 为空，返回 pid=-1", zap.Uint("uid", uid))
        response.OkWithData(-1, c)
        return
    }

    // 3. 组装命令：激活虚拟环境后启动 freqtrade 模拟盘
    // 使用查询到的相对路径 configPath 作为 -c 参数（cmd.Dir 已设置为工作目录）
    cmdStr := `source .venv/bin/activate && exec freqtrade trade \
        -c '` + configPath + `' \
        -s Ai15moreStrategy \
        --user-id ` + strconv.FormatUint(uint64(uid), 10) + ` \
        --dry-run`

    cmd := exec.Command("bash", "-c", cmdStr)
    cmd.Dir = workDir
    cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // 方便整组杀进程

    if err := cmd.Start(); err != nil {
        global.GVA_LOG.Error("freqtrade dry-run 启动失败", zap.Error(err))
        response.FailWithMessage(err.Error(), c)
        return
    }
    pid := cmd.Process.Pid
    // 异步等待退出并记录日志
    go func() {
        _ = cmd.Wait()
        global.GVA_LOG.Info("freqtrade dry-run 已退出", zap.Int("pid", pid))
    }()

    // 5. 立即返回 PID
    response.OkWithData(pid, c) 
}

// DeleteTrades 删除交易报表(模拟盘)
// @Tags Trades
// @Summary 删除交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.Trades true "删除交易报表(模拟盘)"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /trades/deleteTrades [delete]
func (tradesApi *TradesApi) DeleteTrades(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    id := c.Query("id")
    err := tradesService.DeleteTrades(ctx,id)
    if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("删除成功", c)
}

// DeleteTradesByIds 批量删除交易报表(模拟盘)
// @Tags Trades
// @Summary 批量删除交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /trades/deleteTradesByIds [delete]
func (tradesApi *TradesApi) DeleteTradesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    ids := c.QueryArray("ids[]")
    err := tradesService.DeleteTradesByIds(ctx,ids)
    if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("批量删除成功", c)
}

// UpdateTrades 更新交易报表(模拟盘)
// @Tags Trades
// @Summary 更新交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.Trades true "更新交易报表(模拟盘)"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /trades/updateTrades [put]
func (tradesApi *TradesApi) UpdateTrades(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

    var trades running.Trades
    err := c.ShouldBindJSON(&trades)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    err = tradesService.UpdateTrades(ctx,trades)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("更新成功", c)
}

// FindTrades 用id查询交易报表(模拟盘)
// @Tags Trades
// @Summary 用id查询交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询交易报表(模拟盘)"
// @Success 200 {object} response.Response{data=running.Trades,msg=string} "查询成功"
// @Router /trades/findTrades [get]
func (tradesApi *TradesApi) FindTrades(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    id := c.Query("id")
    retrades, err := tradesService.GetTrades(ctx,id)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
        return
    }
    response.OkWithData(retrades, c)
}
// GetTradesList 分页获取交易报表(模拟盘)列表
// @Tags Trades
// @Summary 分页获取交易报表(模拟盘)列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.TradesSearch true "分页获取交易报表(模拟盘)列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /trades/getTradesList [get]
func (tradesApi *TradesApi) GetTradesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var pageInfo runningReq.TradesSearch
    if err := c.ShouldBindQuery(&pageInfo); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }

    // 根据角色决定查询范围：普通用户(AuthorityId=888)仅能查自己的，管理员可查全部
    // 说明：这里不信任前端传入的 userId，而是根据登录态在后端强制过滤，避免越权
    uid := utils.GetUserID(c)                 // 当前登录用户ID（来自 JWT）
    authId := utils.GetUserAuthorityId(c)     // 当前登录用户角色ID（来自 JWT）

    // 调试：打印当前用户ID与角色ID，以及完整的 claims，帮助定位权限判断不生效的问题
    claimsVal, _ := c.Get("claims") // 从 gin 上下文获取JWT Claims（若存在）
    global.GVA_LOG.Info("GetTradesList 调试", // 使用 zap 记录到控制台
        zap.Uint("uid", uid),
        zap.Uint("authId", authId),
        zap.Any("claims", claimsVal),
    )

    // 预先声明用于承接查询结果的变量，避免在分支中用 := 产生局部作用域导致后续不可见
    var (
        list  []running.Trades // 查询结果列表
        total int64            // 结果总数
        err   error            // 错误
    )

    // 普通用户仅能查询自己的，管理员可查全部
    if authId != 888 {
        // 管理员或拥有更高权限的角色，查询全部数据
        list, total, err = tradesService.GetTradesInfoList(ctx, pageInfo)
    } else {
        // 普通用户（AuthorityId=888），仅能查询自己的数据
        list, total, err = tradesService.GetTradesInfoListWithUid(ctx, pageInfo, uid)
    }
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetDryRun 获取当前用户 freqtrade dry-run 进程PID（无则返回0）
// @Tags Trades
// @Summary 获取当前用户 dry-run 进程PID（无则返回0）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=int,msg=string} "查询成功"
// @Router /trades/getDryRun [get]
func (tradesApi *TradesApi) GetDryRun(c *gin.Context) {
    // 中文注释：从登录态获取当前用户ID
    uid := utils.GetUserID(c)
    uidStr := strconv.FormatUint(uint64(uid), 10)

    // 中文注释：通过 ps/grep 管道查询包含 freqtrade trade 和 userId 的进程；过滤掉 grep 本身
    // 说明：此命令适用于 Linux (Ubuntu) 环境；在 macOS 上也可工作，但建议在服务端(Ubuntu)调用
    cmdStr := "ps -ef | grep freqtrade | grep trade | grep " + uidStr + " | grep -v grep"

    // 中文注释：执行命令并获取输出；若无匹配行，grep 退出码非0，此处不将其视为错误
    out, _ := exec.Command("bash", "-c", cmdStr).CombinedOutput()
    lines := strings.Split(strings.TrimSpace(string(out)), "\n")

    pid := 0
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }
        // ps -ef 输出格式：UID PID PPID C STIME TTY TIME CMD
        fields := strings.Fields(line)
        if len(fields) >= 2 {
            if v, err := strconv.Atoi(fields[1]); err == nil {
                pid = v
                break // 取第一条匹配记录即可
            }
        }
    }
    response.OkWithData(pid, c)
}

// GetTradesPublic 不需要鉴权的交易报表(模拟盘)接口
// @Tags Trades
// @Summary 不需要鉴权的交易报表(模拟盘)接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /trades/getTradesPublic [get]
func (tradesApi *TradesApi) GetTradesPublic(c *gin.Context) {
    // 创建业务用Context
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
