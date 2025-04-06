package ExaMan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ExaMan"
	ExaManReq "github.com/flipped-aurora/gin-vue-admin/server/model/ExaMan/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExamManagementApi struct{}

// CreateExamManagement 创建考试管理
// @Tags ExamManagement
// @Summary 创建考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ExaMan.ExamManagement true "创建考试管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Em/createExamManagement [post]
func (EmApi *ExamManagementApi) CreateExamManagement(c *gin.Context) {
	var Em ExaMan.ExamManagement
	err := c.ShouldBindJSON(&Em)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EmService.CreateExamManagement(&Em)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExamManagement 删除考试管理
// @Tags ExamManagement
// @Summary 删除考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ExaMan.ExamManagement true "删除考试管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Em/deleteExamManagement [delete]
func (EmApi *ExamManagementApi) DeleteExamManagement(c *gin.Context) {
	ID := c.Query("ID")
	err := EmService.DeleteExamManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteExamManagementByIds 批量删除考试管理
// @Tags ExamManagement
// @Summary 批量删除考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Em/deleteExamManagementByIds [delete]
func (EmApi *ExamManagementApi) DeleteExamManagementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := EmService.DeleteExamManagementByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateExamManagement 更新考试管理
// @Tags ExamManagement
// @Summary 更新考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ExaMan.ExamManagement true "更新考试管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Em/updateExamManagement [put]
func (EmApi *ExamManagementApi) UpdateExamManagement(c *gin.Context) {
	var Em ExaMan.ExamManagement
	err := c.ShouldBindJSON(&Em)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = EmService.UpdateExamManagement(Em)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindExamManagement 用id查询考试管理
// @Tags ExamManagement
// @Summary 用id查询考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ExaMan.ExamManagement true "用id查询考试管理"
// @Success 200 {object} response.Response{data=ExaMan.ExamManagement,msg=string} "查询成功"
// @Router /Em/findExamManagement [get]
func (EmApi *ExamManagementApi) FindExamManagement(c *gin.Context) {
	ID := c.Query("ID")
	reEm, err := EmService.GetExamManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reEm, c)
}

// GetExamManagementList 分页获取考试管理列表
// @Tags ExamManagement
// @Summary 分页获取考试管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ExaManReq.ExamManagementSearch true "分页获取考试管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Em/getExamManagementList [get]
func (EmApi *ExamManagementApi) GetExamManagementList(c *gin.Context) {
	var pageInfo ExaManReq.ExamManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EmService.GetExamManagementInfoList(pageInfo)
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

// GetExamManagementPublic 不需要鉴权的考试管理接口
// @Tags ExamManagement
// @Summary 不需要鉴权的考试管理接口
// @accept application/json
// @Produce application/json
// @Param data query ExaManReq.ExamManagementSearch true "分页获取考试管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Em/getExamManagementPublic [get]
func (EmApi *ExamManagementApi) GetExamManagementPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	EmService.GetExamManagementPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的考试管理接口信息",
	}, "获取成功", c)
}

// GetExamManagementList 分页获取考试管理列表
// @Tags ExamManagement
// @Summary 分页获取考试管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ExaManReq.ExamManagementSearch true "分页获取考试管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Em/getExamManagementList [get]
func (EmApi *ExamManagementApi) GetUserExamManagementList(c *gin.Context) {
	var pageInfo ExaManReq.ExamManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := EmService.GetUserExamManagementInfoList(pageInfo)
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
