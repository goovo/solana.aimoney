package running

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/running"
    runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
    "github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
    "go.uber.org/zap"
)

type SysUserApiApi struct {}



// CreateSysUserApi 创建用户APIs
// @Tags SysUserApi
// @Summary 创建用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserApi true "创建用户APIs"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysUserApi/createSysUserApi [post]
func (sysUserApiApi *SysUserApiApi) CreateSysUserApi(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var sysUserApi running.SysUserApi
	err := c.ShouldBindJSON(&sysUserApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 如果前端没有传递status字段，设置默认值5
	if sysUserApi.Status == nil {
		defaultStatus := 5
		sysUserApi.Status = &defaultStatus
	}

	err = sysUserApiService.CreateSysUserApi(ctx,&sysUserApi)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSysUserApi 删除用户APIs
// @Tags SysUserApi
// @Summary 删除用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserApi true "删除用户APIs"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysUserApi/deleteSysUserApi [delete]
func (sysUserApiApi *SysUserApiApi) DeleteSysUserApi(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := sysUserApiService.DeleteSysUserApi(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysUserApiByIds 批量删除用户APIs
// @Tags SysUserApi
// @Summary 批量删除用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysUserApi/deleteSysUserApiByIds [delete]
func (sysUserApiApi *SysUserApiApi) DeleteSysUserApiByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := sysUserApiService.DeleteSysUserApiByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysUserApi 更新用户APIs
// @Tags SysUserApi
// @Summary 更新用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserApi true "更新用户APIs"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysUserApi/updateSysUserApi [put]
func (sysUserApiApi *SysUserApiApi) UpdateSysUserApi(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var sysUserApi running.SysUserApi
	err := c.ShouldBindJSON(&sysUserApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysUserApiService.UpdateSysUserApi(ctx,sysUserApi)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysUserApi 用id查询用户APIs
// @Tags SysUserApi
// @Summary 用id查询用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询用户APIs"
// @Success 200 {object} response.Response{data=running.SysUserApi,msg=string} "查询成功"
// @Router /sysUserApi/findSysUserApi [get]
func (sysUserApiApi *SysUserApiApi) FindSysUserApi(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	resysUserApi, err := sysUserApiService.GetSysUserApi(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(resysUserApi, c)
}
// GetSysUserApiList 分页获取用户APIs列表
// @Tags SysUserApi
// @Summary 分页获取用户APIs列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserApiSearch true "分页获取用户APIs列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysUserApi/getSysUserApiList [get]
func (sysUserApiApi *SysUserApiApi) GetSysUserApiList(c *gin.Context) {
	// 创建业务用Context
    ctx := c.Request.Context()
	var pageInfo runningReq.SysUserApiSearch

    // 先绑定查询参数（分页等）
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

    // 根据角色决定查询范围：普通用户(AuthorityId=888)仅能查自己的，管理员可查全部
    // 说明：这里不信任前端传入的 userId，而是根据登录态在后端强制过滤，避免越权
    uid := utils.GetUserID(c)                 // 当前登录用户ID（来自 JWT）
    authId := utils.GetUserAuthorityId(c)     // 当前登录用户角色ID（来自 JWT）

    // 调试：打印当前用户ID与角色ID，以及完整的 claims，帮助定位权限判断不生效的问题
    claimsVal, _ := c.Get("claims") // 从 gin 上下文获取JWT Claims（若存在）
    global.GVA_LOG.Info("GetSysUserApiList 调试", // 使用 zap 记录到控制台
        zap.Uint("uid", uid),
        zap.Uint("authId", authId),
        zap.Any("claims", claimsVal),
    )

    var (
        list  []running.SysUserApi
        total int64
    )

    if authId == 888 {
        // 普通用户：仅查询自己的数据
        global.GVA_LOG.Info("GetSysUserApiList 命中普通用户分支，仅查询自己的数据", // 记录命中分支
            zap.Uint("authId", authId),
            zap.Uint("uid", uid),
        )
        list, total, err = sysUserApiService.GetSysUserApiInfoListWithUid(ctx, pageInfo, uid)
    } else {
        // 管理员：查询所有数据
        global.GVA_LOG.Info("GetSysUserApiList 命中管理员分支，查询全部数据", // 记录命中分支
            zap.Uint("authId", authId),
            zap.Uint("uid", uid),
        )
        list, total, err = sysUserApiService.GetSysUserApiInfoList(ctx, pageInfo)
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

// GetUserApiList 分页获取用户APIs列表--基于当前登录的用户id
// @Tags SysUserApi
// @Summary 分页获取用户APIs列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserApiSearch true "分页获取用户APIs列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysUserApi/getUserApiList [get]
func (sysUserApiApi *SysUserApiApi) GetUserApiList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()
	uid := utils.GetUserID(c)

	var pageInfo runningReq.SysUserApiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysUserApiService.GetSysUserApiInfoListWithUid(ctx,pageInfo,uid)
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

// GetSysUserApiPublic 不需要鉴权的用户APIs接口
// @Tags SysUserApi
// @Summary 不需要鉴权的用户APIs接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserApi/getSysUserApiPublic [get]
func (sysUserApiApi *SysUserApiApi) GetSysUserApiPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    sysUserApiService.GetSysUserApiPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户APIs接口信息",
    }, "获取成功", c)
}
