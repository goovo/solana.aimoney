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

type SysUserAibotApi struct{}

// CreateSysUserAibot 创建授权交易
// @Tags SysUserAibot
// @Summary 创建授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserAibot true "创建授权交易"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysUserAibot/createSysUserAibot [post]
func (sysUserAibotApi *SysUserAibotApi) CreateSysUserAibot(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var sysUserAibot running.SysUserAibot
	err := c.ShouldBindJSON(&sysUserAibot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysUserAibotService.CreateSysUserAibot(ctx, &sysUserAibot)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysUserAibot 删除授权交易
// @Tags SysUserAibot
// @Summary 删除授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserAibot true "删除授权交易"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysUserAibot/deleteSysUserAibot [delete]
func (sysUserAibotApi *SysUserAibotApi) DeleteSysUserAibot(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	userId := c.Query("userId")
	err := sysUserAibotService.DeleteSysUserAibot(ctx, userId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysUserAibotByIds 批量删除授权交易
// @Tags SysUserAibot
// @Summary 批量删除授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysUserAibot/deleteSysUserAibotByIds [delete]
func (sysUserAibotApi *SysUserAibotApi) DeleteSysUserAibotByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	userIds := c.QueryArray("userIds[]")
	err := sysUserAibotService.DeleteSysUserAibotByIds(ctx, userIds)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysUserAibot 更新授权交易
// @Tags SysUserAibot
// @Summary 更新授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserAibot true "更新授权交易"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysUserAibot/updateSysUserAibot [put]
func (sysUserAibotApi *SysUserAibotApi) UpdateSysUserAibot(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var sysUserAibot running.SysUserAibot
	err := c.ShouldBindJSON(&sysUserAibot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysUserAibotService.UpdateSysUserAibot(ctx, sysUserAibot)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysUserAibot 用id查询授权交易
// @Tags SysUserAibot
// @Summary 用id查询授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param userId query int true "用id查询授权交易"
// @Success 200 {object} response.Response{data=running.SysUserAibot,msg=string} "查询成功"
// @Router /sysUserAibot/findSysUserAibot [get]
func (sysUserAibotApi *SysUserAibotApi) FindSysUserAibot(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	userId := c.Query("userId")
	resysUserAibot, err := sysUserAibotService.GetSysUserAibot(ctx, userId)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(resysUserAibot, c)
}

// GetSysUserAibotList 分页获取授权交易列表
// @Tags SysUserAibot
// @Summary 分页获取授权交易列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserAibotSearch true "分页获取授权交易列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysUserAibot/getSysUserAibotList [get]
func (sysUserAibotApi *SysUserAibotApi) GetSysUserAibotList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo runningReq.SysUserAibotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)             // 当前登录用户ID（来自 JWT）
	authId := utils.GetUserAuthorityId(c) // 当前登录用户角色ID（来自 JWT）
	var list []running.SysUserAibot
	var total int64
	if authId == 888 || authId == 9528 {
		list, total, err = sysUserAibotService.GetSysUserAibotInfoListWithUid(ctx, pageInfo, int64(uid))
	} else {
		list, total, err = sysUserAibotService.GetSysUserAibotInfoList(ctx, pageInfo)
	}
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetSysUserAibotPublic 不需要鉴权的授权交易接口
// @Tags SysUserAibot
// @Summary 不需要鉴权的授权交易接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserAibot/getSysUserAibotPublic [get]
func (sysUserAibotApi *SysUserAibotApi) GetSysUserAibotPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sysUserAibotService.GetSysUserAibotPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的授权交易接口信息",
	}, "获取成功", c)
}
