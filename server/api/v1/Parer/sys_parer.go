package Parer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Parer"
	ParerReq "github.com/flipped-aurora/gin-vue-admin/server/model/Parer/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaperApi struct{}

// CreatePaper 创建试卷管理
// @Tags Paper
// @Summary 创建试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Parer.Paper true "创建试卷管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /P/createPaper [post]
func (PApi *PaperApi) CreatePaper(c *gin.Context) {
	var P Parer.Paper
	err := c.ShouldBindJSON(&P)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PService.CreatePaper(&P)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePaper 删除试卷管理
// @Tags Paper
// @Summary 删除试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Parer.Paper true "删除试卷管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /P/deletePaper [delete]
func (PApi *PaperApi) DeletePaper(c *gin.Context) {
	ID := c.Query("ID")
	err := PService.DeletePaper(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePaperByIds 批量删除试卷管理
// @Tags Paper
// @Summary 批量删除试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /P/deletePaperByIds [delete]
func (PApi *PaperApi) DeletePaperByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := PService.DeletePaperByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePaper 更新试卷管理
// @Tags Paper
// @Summary 更新试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Parer.Paper true "更新试卷管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /P/updatePaper [put]
func (PApi *PaperApi) UpdatePaper(c *gin.Context) {
	var P Parer.Paper
	err := c.ShouldBindJSON(&P)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PService.UpdatePaper(P)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPaper 用id查询试卷管理
// @Tags Paper
// @Summary 用id查询试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Parer.Paper true "用id查询试卷管理"
// @Success 200 {object} response.Response{data=Parer.Paper,msg=string} "查询成功"
// @Router /P/findPaper [get]
func (PApi *PaperApi) FindPaper(c *gin.Context) {
	ID := c.Query("ID")
	reP, err := PService.GetPaper(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reP, c)
}

// GetPaperList 分页获取试卷管理列表
// @Tags Paper
// @Summary 分页获取试卷管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ParerReq.PaperSearch true "分页获取试卷管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /P/getPaperList [get]
func (PApi *PaperApi) GetPaperList(c *gin.Context) {
	var pageInfo ParerReq.PaperSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PService.GetPaperInfoList(pageInfo)
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

// GetPaperPublic 不需要鉴权的试卷管理接口
// @Tags Paper
// @Summary 不需要鉴权的试卷管理接口
// @accept application/json
// @Produce application/json
// @Param data query ParerReq.PaperSearch true "分页获取试卷管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /P/getPaperPublic [get]
func (PApi *PaperApi) GetPaperPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	PService.GetPaperPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的试卷管理接口信息",
	}, "获取成功", c)
}

// GetPaperList 分页获取试卷管理列表
// @Tags Paper
// @Summary 分页获取试卷管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ParerReq.PaperSearch true "分页获取试卷管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /P/getPaperList [get]
func (PApi *PaperApi) GetPaperListAll(c *gin.Context) {
	var pageInfo ParerReq.PaperSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PService.GetPaperInfoListAll(pageInfo)
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

// FindPaper 用id查询试卷管理
// @Tags Paper
// @Summary 用id查询试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Parer.Paper true "用id查询试卷管理"
// @Success 200 {object} response.Response{data=Parer.Paper,msg=string} "查询成功"
// @Router /P/findPaper [get]
func (PApi *PaperApi) FindPaperExam(c *gin.Context) {
	ID := c.Query("ID")
	reP, err := PService.GetPaperExam(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reP, c)
}
