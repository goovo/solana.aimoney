package running

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/running"
    runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysUserAssetsApi struct {}



// CreateSysUserAssets 创建用户资产
// @Tags SysUserAssets
// @Summary 创建用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserAssets true "创建用户资产"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysUserAssets/createSysUserAssets [post]
func (sysUserAssetsApi *SysUserAssetsApi) CreateSysUserAssets(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var sysUserAssets running.SysUserAssets
	err := c.ShouldBindJSON(&sysUserAssets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysUserAssetsService.CreateSysUserAssets(ctx,&sysUserAssets)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSysUserAssets 删除用户资产
// @Tags SysUserAssets
// @Summary 删除用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserAssets true "删除用户资产"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysUserAssets/deleteSysUserAssets [delete]
func (sysUserAssetsApi *SysUserAssetsApi) DeleteSysUserAssets(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := sysUserAssetsService.DeleteSysUserAssets(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysUserAssetsByIds 批量删除用户资产
// @Tags SysUserAssets
// @Summary 批量删除用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysUserAssets/deleteSysUserAssetsByIds [delete]
func (sysUserAssetsApi *SysUserAssetsApi) DeleteSysUserAssetsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := sysUserAssetsService.DeleteSysUserAssetsByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysUserAssets 更新用户资产
// @Tags SysUserAssets
// @Summary 更新用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body running.SysUserAssets true "更新用户资产"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysUserAssets/updateSysUserAssets [put]
func (sysUserAssetsApi *SysUserAssetsApi) UpdateSysUserAssets(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var sysUserAssets running.SysUserAssets
	err := c.ShouldBindJSON(&sysUserAssets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysUserAssetsService.UpdateSysUserAssets(ctx,sysUserAssets)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysUserAssets 用id查询用户资产
// @Tags SysUserAssets
// @Summary 用id查询用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询用户资产"
// @Success 200 {object} response.Response{data=running.SysUserAssets,msg=string} "查询成功"
// @Router /sysUserAssets/findSysUserAssets [get]
func (sysUserAssetsApi *SysUserAssetsApi) FindSysUserAssets(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	resysUserAssets, err := sysUserAssetsService.GetSysUserAssets(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(resysUserAssets, c)
}
// GetSysUserAssetsList 分页获取用户资产列表
// @Tags SysUserAssets
// @Summary 分页获取用户资产列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserAssetsSearch true "分页获取用户资产列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysUserAssets/getSysUserAssetsList [get]
func (sysUserAssetsApi *SysUserAssetsApi) GetSysUserAssetsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo runningReq.SysUserAssetsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysUserAssetsService.GetSysUserAssetsInfoList(ctx,pageInfo)
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

// GetSysUserAssetsPublic 不需要鉴权的用户资产接口
// @Tags SysUserAssets
// @Summary 不需要鉴权的用户资产接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserAssets/getSysUserAssetsPublic [get]
func (sysUserAssetsApi *SysUserAssetsApi) GetSysUserAssetsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    sysUserAssetsService.GetSysUserAssetsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户资产接口信息",
    }, "获取成功", c)
}
