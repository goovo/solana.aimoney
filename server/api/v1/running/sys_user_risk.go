package running

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/running"
    runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysUserRiskApi struct {}

// SetUserRisk 设置用户风险等级
// @Tags SetUserRisk
// @Summary 创建用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserRisk true "设置用户风险等级"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysUserRisk/SetUserRisk [post]
func (sysUserRiskApi *SysUserRiskApi) SetUserRisk(c *gin.Context) {
    // 从 token 中获取当前用户ID（中文注释）
    uid := utils.GetUserID(c)

    // 创建业务用Context
    ctx := c.Request.Context()

    // 绑定请求体中的风险等级数据（中文注释）
    var sysUserRisk running.SysUserRisk
    if err := c.ShouldBindJSON(&sysUserRisk); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }

    // 调用服务层：按 userId 进行插入/更新（中文注释）
    if err := sysUserRiskService.SetUserRisk(ctx, &sysUserRisk, int(uid)); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败:"+err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}

// CreateSysUserRisk 创建用户风险等级
// @Tags SysUserRisk
// @Summary 创建用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserRisk true "创建用户风险等级"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysUserRisk/createSysUserRisk [post]
func (sysUserRiskApi *SysUserRiskApi) CreateSysUserRisk(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var sysUserRisk running.SysUserRisk
    err := c.ShouldBindJSON(&sysUserRisk)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    err = sysUserRiskService.CreateSysUserRisk(ctx,&sysUserRisk)
    if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}

// DeleteSysUserRisk 删除用户风险等级
// @Tags SysUserRisk
// @Summary 删除用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserRisk true "删除用户风险等级"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysUserRisk/deleteSysUserRisk [delete]
func (sysUserRiskApi *SysUserRiskApi) DeleteSysUserRisk(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    id := c.Query("id")
    err := sysUserRiskService.DeleteSysUserRisk(ctx,id)
    if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("删除成功", c)
}

// DeleteSysUserRiskByIds 批量删除用户风险等级
// @Tags SysUserRisk
// @Summary 批量删除用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysUserRisk/deleteSysUserRiskByIds [delete]
func (sysUserRiskApi *SysUserRiskApi) DeleteSysUserRiskByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    ids := c.QueryArray("ids[]")
    err := sysUserRiskService.DeleteSysUserRiskByIds(ctx,ids)
    if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSysUserRisk 更新用户风险等级
// @Tags SysUserRisk
// @Summary 更新用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserRisk true "更新用户风险等级"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysUserRisk/updateSysUserRisk [put]
func (sysUserRiskApi *SysUserRiskApi) UpdateSysUserRisk(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

    var sysUserRisk running.SysUserRisk
    err := c.ShouldBindJSON(&sysUserRisk)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    err = sysUserRiskService.UpdateSysUserRisk(ctx,sysUserRisk)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("更新成功", c)
}

// FindSysUserRisk 用id查询用户风险等级
// @Tags SysUserRisk
// @Summary 用id查询用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询用户风险等级"
// @Success 200 {object} response.Response{data=running.SysUserRisk,msg=string} "查询成功"
// @Router /sysUserRisk/findSysUserRisk [get]
func (sysUserRiskApi *SysUserRiskApi) FindSysUserRisk(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    id := c.Query("id")
    resysUserRisk, err := sysUserRiskService.GetSysUserRisk(ctx,id)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
        return
    }
    response.OkWithData(resysUserRisk, c)
}
// GetSysUserRiskList 分页获取用户风险等级列表
// @Tags SysUserRisk
// @Summary 分页获取用户风险等级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserRiskSearch true "分页获取用户风险等级列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysUserRisk/getSysUserRiskList [get]
func (sysUserRiskApi *SysUserRiskApi) GetSysUserRiskList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var pageInfo runningReq.SysUserRiskSearch
    err := c.ShouldBindQuery(&pageInfo)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    list, total, err := sysUserRiskService.GetSysUserRiskInfoList(ctx,pageInfo)
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

// GetSysUserRiskPublic 不需要鉴权的用户风险等级接口
// @Tags SysUserRisk
// @Summary 不需要鉴权的用户风险等级接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserRisk/getSysUserRiskPublic [get]
func (sysUserRiskApi *SysUserRiskApi) GetSysUserRiskPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    sysUserRiskService.GetSysUserRiskPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户风险等级接口信息",
    }, "获取成功", c)
}
