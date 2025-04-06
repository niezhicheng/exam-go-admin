package AnsweSituation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/AnsweSituation"
	AnsweSituationReq "github.com/flipped-aurora/gin-vue-admin/server/model/AnsweSituation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnsweringSituationApi struct{}

// CreateAnsweringSituation 创建答题情况
// @Tags AnsweringSituation
// @Summary 创建答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AnsweSituation.AnsweringSituation true "创建答题情况"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /As/createAnsweringSituation [post]
func (AsApi *AnsweringSituationApi) CreateAnsweringSituation(c *gin.Context) {
	var As AnsweSituation.AnsweringSituation
	err := c.ShouldBindJSON(&As)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = AsService.CreateAnsweringSituation(&As)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAnsweringSituation 删除答题情况
// @Tags AnsweringSituation
// @Summary 删除答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AnsweSituation.AnsweringSituation true "删除答题情况"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /As/deleteAnsweringSituation [delete]
func (AsApi *AnsweringSituationApi) DeleteAnsweringSituation(c *gin.Context) {
	ID := c.Query("ID")
	err := AsService.DeleteAnsweringSituation(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAnsweringSituationByIds 批量删除答题情况
// @Tags AnsweringSituation
// @Summary 批量删除答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /As/deleteAnsweringSituationByIds [delete]
func (AsApi *AnsweringSituationApi) DeleteAnsweringSituationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := AsService.DeleteAnsweringSituationByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAnsweringSituation 更新答题情况
// @Tags AnsweringSituation
// @Summary 更新答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AnsweSituation.AnsweringSituation true "更新答题情况"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /As/updateAnsweringSituation [put]
func (AsApi *AnsweringSituationApi) UpdateAnsweringSituation(c *gin.Context) {
	var As AnsweSituation.AnsweringSituation
	err := c.ShouldBindJSON(&As)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = AsService.UpdateAnsweringSituation(As)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAnsweringSituation 用id查询答题情况
// @Tags AnsweringSituation
// @Summary 用id查询答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AnsweSituation.AnsweringSituation true "用id查询答题情况"
// @Success 200 {object} response.Response{data=AnsweSituation.AnsweringSituation,msg=string} "查询成功"
// @Router /As/findAnsweringSituation [get]
func (AsApi *AnsweringSituationApi) FindAnsweringSituation(c *gin.Context) {
	ID := c.Query("ID")
	reAs, err := AsService.GetAnsweringSituation(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reAs, c)
}

// GetAnsweringSituationList 分页获取答题情况列表
// @Tags AnsweringSituation
// @Summary 分页获取答题情况列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AnsweSituationReq.AnsweringSituationSearch true "分页获取答题情况列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /As/getAnsweringSituationList [get]
func (AsApi *AnsweringSituationApi) GetAnsweringSituationList(c *gin.Context) {
	var pageInfo AnsweSituationReq.AnsweringSituationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := AsService.GetAnsweringSituationInfoList(pageInfo)
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

// GetAnsweringSituationPublic 不需要鉴权的答题情况接口
// @Tags AnsweringSituation
// @Summary 不需要鉴权的答题情况接口
// @accept application/json
// @Produce application/json
// @Param data query AnsweSituationReq.AnsweringSituationSearch true "分页获取答题情况列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /As/getAnsweringSituationPublic [get]
func (AsApi *AnsweringSituationApi) GetAnsweringSituationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	AsService.GetAnsweringSituationPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的答题情况接口信息",
	}, "获取成功", c)
}

// FindAnsweringSituation 用id查询答题情况
// @Tags AnsweringSituation
// @Summary 用id查询答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AnsweSituation.AnsweringSituation true "用id查询答题情况"
// @Success 200 {object} response.Response{data=AnsweSituation.AnsweringSituation,msg=string} "查询成功"
// @Router /As/findAnsweringSituation [get]
func (AsApi *AnsweringSituationApi) FindExamAnsweringSituation(c *gin.Context) {
	ID := c.Query("ID")
	userid := utils.GetUserID(c)
	reAs, err := AsService.GetExamAnsweringSituation(ID, userid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reAs, c)
}

// CreateAnsweringSituation 创建答题情况
// @Tags AnsweringSituation
// @Summary 创建答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AnsweSituation.AnsweringSituation true "创建答题情况"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /As/createAnsweringSituation [post]
func (AsApi *AnsweringSituationApi) CreateUserAnsweringSituation(c *gin.Context) {
	var As AnsweSituation.AnsweringSituation
	err := c.ShouldBindJSON(&As)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	As.UserId = utils.GetUserID(c)
	err = AsService.CreateUserAnsweringSituation(&As)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// CreateAnsweringSituation 创建答题情况
// @Tags AnsweringSituation
// @Summary 创建答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AnsweSituation.AnsweringSituation true "创建答题情况"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /As/createAnsweringSituation [post]
func (AsApi *AnsweringSituationApi) SumbitUserAnsweringSituation(c *gin.Context) {
	var As AnsweSituation.AnsweringSituation
	err := c.ShouldBindJSON(&As)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	As.UserId = utils.GetUserID(c)
	err = AsService.SubmitUserAnsweringSituation(&As)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
