package ParerDetail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ParerDetail"
	ParerDetailReq "github.com/flipped-aurora/gin-vue-admin/server/model/ParerDetail/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaperDetailApi struct{}

// CreatePaperDetail 创建试卷详细
// @Tags PaperDetail
// @Summary 创建试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ParerDetail.PaperDetail true "创建试卷详细"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Dp/createPaperDetail [post]
func (DpApi *PaperDetailApi) CreatePaperDetail(c *gin.Context) {
	var Dp ParerDetail.PaperDetail
	err := c.ShouldBindJSON(&Dp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DpService.CreatePaperDetail(&Dp)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePaperDetail 删除试卷详细
// @Tags PaperDetail
// @Summary 删除试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ParerDetail.PaperDetail true "删除试卷详细"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Dp/deletePaperDetail [delete]
func (DpApi *PaperDetailApi) DeletePaperDetail(c *gin.Context) {
	ID := c.Query("ID")
	err := DpService.DeletePaperDetail(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePaperDetailByIds 批量删除试卷详细
// @Tags PaperDetail
// @Summary 批量删除试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Dp/deletePaperDetailByIds [delete]
func (DpApi *PaperDetailApi) DeletePaperDetailByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := DpService.DeletePaperDetailByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePaperDetail 更新试卷详细
// @Tags PaperDetail
// @Summary 更新试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ParerDetail.PaperDetail true "更新试卷详细"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Dp/updatePaperDetail [put]
func (DpApi *PaperDetailApi) UpdatePaperDetail(c *gin.Context) {
	var Dp ParerDetail.PaperDetail
	err := c.ShouldBindJSON(&Dp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DpService.UpdatePaperDetail(Dp)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPaperDetail 用id查询试卷详细
// @Tags PaperDetail
// @Summary 用id查询试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ParerDetail.PaperDetail true "用id查询试卷详细"
// @Success 200 {object} response.Response{data=ParerDetail.PaperDetail,msg=string} "查询成功"
// @Router /Dp/findPaperDetail [get]
func (DpApi *PaperDetailApi) FindPaperDetail(c *gin.Context) {
	ID := c.Query("ID")
	reDp, err := DpService.GetPaperDetail(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reDp, c)
}

// GetPaperDetailList 分页获取试卷详细列表
// @Tags PaperDetail
// @Summary 分页获取试卷详细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ParerDetailReq.PaperDetailSearch true "分页获取试卷详细列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Dp/getPaperDetailList [get]
func (DpApi *PaperDetailApi) GetPaperDetailList(c *gin.Context) {
	var pageInfo ParerDetailReq.PaperDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DpService.GetPaperDetailInfoList(pageInfo)
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

// GetPaperDetailPublic 不需要鉴权的试卷详细接口
// @Tags PaperDetail
// @Summary 不需要鉴权的试卷详细接口
// @accept application/json
// @Produce application/json
// @Param data query ParerDetailReq.PaperDetailSearch true "分页获取试卷详细列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Dp/getPaperDetailPublic [get]
func (DpApi *PaperDetailApi) GetPaperDetailPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	DpService.GetPaperDetailPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的试卷详细接口信息",
	}, "获取成功", c)
}

// FindPaperDetailAll 根据试卷id获取试题列表
// @Tags PaperDetail
// @Summary 根据试卷id获取试题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ParerDetailReq.PaperDetailSearch true "根据试卷id获取试题列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Dp/getPaperDetailList [get]
func (DpApi *PaperDetailApi) FindPaperDetailAll(c *gin.Context) {
	var pageInfo ParerDetailReq.PaperDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DpService.GetPaperInfoList(pageInfo)
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

// FindPaperDetailAll 根据试卷id获取试题列表
// @Tags PaperDetail
// @Summary 根据试卷id获取试题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ParerDetailReq.PaperDetailSearch true "根据试卷id获取试题列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Dp/getPaperDetailList [get]
func (DpApi *PaperDetailApi) FindPaperNoDetailAll(c *gin.Context) {
	var pageInfo ParerDetailReq.PaperDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DpService.GetPaperNoInfoList(pageInfo)
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
