package Examresults

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Examresults"
    ExamresultsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Examresults/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ExamResultsApi struct {}



// CreateExamResults 创建考试成绩
// @Tags ExamResults
// @Summary 创建考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Examresults.ExamResults true "创建考试成绩"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /Er/createExamResults [post]
func (ErApi *ExamResultsApi) CreateExamResults(c *gin.Context) {
	var Er Examresults.ExamResults
	err := c.ShouldBindJSON(&Er)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ErService.CreateExamResults(&Er)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteExamResults 删除考试成绩
// @Tags ExamResults
// @Summary 删除考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Examresults.ExamResults true "删除考试成绩"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /Er/deleteExamResults [delete]
func (ErApi *ExamResultsApi) DeleteExamResults(c *gin.Context) {
	ID := c.Query("ID")
	err := ErService.DeleteExamResults(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteExamResultsByIds 批量删除考试成绩
// @Tags ExamResults
// @Summary 批量删除考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /Er/deleteExamResultsByIds [delete]
func (ErApi *ExamResultsApi) DeleteExamResultsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := ErService.DeleteExamResultsByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateExamResults 更新考试成绩
// @Tags ExamResults
// @Summary 更新考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Examresults.ExamResults true "更新考试成绩"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /Er/updateExamResults [put]
func (ErApi *ExamResultsApi) UpdateExamResults(c *gin.Context) {
	var Er Examresults.ExamResults
	err := c.ShouldBindJSON(&Er)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ErService.UpdateExamResults(Er)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindExamResults 用id查询考试成绩
// @Tags ExamResults
// @Summary 用id查询考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Examresults.ExamResults true "用id查询考试成绩"
// @Success 200 {object} response.Response{data=Examresults.ExamResults,msg=string} "查询成功"
// @Router /Er/findExamResults [get]
func (ErApi *ExamResultsApi) FindExamResults(c *gin.Context) {
	ID := c.Query("ID")
	reEr, err := ErService.GetExamResults(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reEr, c)
}

// GetExamResultsList 分页获取考试成绩列表
// @Tags ExamResults
// @Summary 分页获取考试成绩列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ExamresultsReq.ExamResultsSearch true "分页获取考试成绩列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /Er/getExamResultsList [get]
func (ErApi *ExamResultsApi) GetExamResultsList(c *gin.Context) {
	var pageInfo ExamresultsReq.ExamResultsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ErService.GetExamResultsInfoList(pageInfo)
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

// GetExamResultsPublic 不需要鉴权的考试成绩接口
// @Tags ExamResults
// @Summary 不需要鉴权的考试成绩接口
// @accept application/json
// @Produce application/json
// @Param data query ExamresultsReq.ExamResultsSearch true "分页获取考试成绩列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Er/getExamResultsPublic [get]
func (ErApi *ExamResultsApi) GetExamResultsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    ErService.GetExamResultsPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的考试成绩接口信息",
    }, "获取成功", c)
}
