package strategy

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/strategy"
    strategyReq "github.com/flipped-aurora/gin-vue-admin/server/model/strategy/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type KeyvaluestoreApi struct {}



// CreateKeyvaluestore 创建keyvaluestore表
// @Tags Keyvaluestore
// @Summary 创建keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body strategy.Keyvaluestore true "创建keyvaluestore表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /keyvaluestore/createKeyvaluestore [post]
func (keyvaluestoreApi *KeyvaluestoreApi) CreateKeyvaluestore(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var keyvaluestore strategy.Keyvaluestore
	err := c.ShouldBindJSON(&keyvaluestore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = keyvaluestoreService.CreateKeyvaluestore(ctx,&keyvaluestore)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteKeyvaluestore 删除keyvaluestore表
// @Tags Keyvaluestore
// @Summary 删除keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body strategy.Keyvaluestore true "删除keyvaluestore表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /keyvaluestore/deleteKeyvaluestore [delete]
func (keyvaluestoreApi *KeyvaluestoreApi) DeleteKeyvaluestore(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := keyvaluestoreService.DeleteKeyvaluestore(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteKeyvaluestoreByIds 批量删除keyvaluestore表
// @Tags Keyvaluestore
// @Summary 批量删除keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /keyvaluestore/deleteKeyvaluestoreByIds [delete]
func (keyvaluestoreApi *KeyvaluestoreApi) DeleteKeyvaluestoreByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := keyvaluestoreService.DeleteKeyvaluestoreByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateKeyvaluestore 更新keyvaluestore表
// @Tags Keyvaluestore
// @Summary 更新keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body strategy.Keyvaluestore true "更新keyvaluestore表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /keyvaluestore/updateKeyvaluestore [put]
func (keyvaluestoreApi *KeyvaluestoreApi) UpdateKeyvaluestore(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var keyvaluestore strategy.Keyvaluestore
	err := c.ShouldBindJSON(&keyvaluestore)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = keyvaluestoreService.UpdateKeyvaluestore(ctx,keyvaluestore)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindKeyvaluestore 用id查询keyvaluestore表
// @Tags Keyvaluestore
// @Summary 用id查询keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询keyvaluestore表"
// @Success 200 {object} response.Response{data=strategy.Keyvaluestore,msg=string} "查询成功"
// @Router /keyvaluestore/findKeyvaluestore [get]
func (keyvaluestoreApi *KeyvaluestoreApi) FindKeyvaluestore(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	rekeyvaluestore, err := keyvaluestoreService.GetKeyvaluestore(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rekeyvaluestore, c)
}
// GetKeyvaluestoreList 分页获取keyvaluestore表列表
// @Tags Keyvaluestore
// @Summary 分页获取keyvaluestore表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query strategyReq.KeyvaluestoreSearch true "分页获取keyvaluestore表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /keyvaluestore/getKeyvaluestoreList [get]
func (keyvaluestoreApi *KeyvaluestoreApi) GetKeyvaluestoreList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo strategyReq.KeyvaluestoreSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := keyvaluestoreService.GetKeyvaluestoreInfoList(ctx,pageInfo)
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

// GetKeyvaluestorePublic 不需要鉴权的keyvaluestore表接口
// @Tags Keyvaluestore
// @Summary 不需要鉴权的keyvaluestore表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /keyvaluestore/getKeyvaluestorePublic [get]
func (keyvaluestoreApi *KeyvaluestoreApi) GetKeyvaluestorePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    keyvaluestoreService.GetKeyvaluestorePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的keyvaluestore表接口信息",
    }, "获取成功", c)
}
