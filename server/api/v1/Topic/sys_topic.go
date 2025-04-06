package Topic

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Topic"
    TopicReq "github.com/flipped-aurora/gin-vue-admin/server/model/Topic/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TopicApi struct {}



// CreateTopic 创建题目管理
// @Tags Topic
// @Summary 创建题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Topic.Topic true "创建题目管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /T/createTopic [post]
func (TApi *TopicApi) CreateTopic(c *gin.Context) {
	var T Topic.Topic
	err := c.ShouldBindJSON(&T)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TService.CreateTopic(&T)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteTopic 删除题目管理
// @Tags Topic
// @Summary 删除题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Topic.Topic true "删除题目管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /T/deleteTopic [delete]
func (TApi *TopicApi) DeleteTopic(c *gin.Context) {
	ID := c.Query("ID")
	err := TService.DeleteTopic(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTopicByIds 批量删除题目管理
// @Tags Topic
// @Summary 批量删除题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /T/deleteTopicByIds [delete]
func (TApi *TopicApi) DeleteTopicByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := TService.DeleteTopicByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTopic 更新题目管理
// @Tags Topic
// @Summary 更新题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Topic.Topic true "更新题目管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /T/updateTopic [put]
func (TApi *TopicApi) UpdateTopic(c *gin.Context) {
	var T Topic.Topic
	err := c.ShouldBindJSON(&T)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = TService.UpdateTopic(T)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTopic 用id查询题目管理
// @Tags Topic
// @Summary 用id查询题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Topic.Topic true "用id查询题目管理"
// @Success 200 {object} response.Response{data=Topic.Topic,msg=string} "查询成功"
// @Router /T/findTopic [get]
func (TApi *TopicApi) FindTopic(c *gin.Context) {
	ID := c.Query("ID")
	reT, err := TService.GetTopic(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reT, c)
}

// GetTopicList 分页获取题目管理列表
// @Tags Topic
// @Summary 分页获取题目管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TopicReq.TopicSearch true "分页获取题目管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /T/getTopicList [get]
func (TApi *TopicApi) GetTopicList(c *gin.Context) {
	var pageInfo TopicReq.TopicSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := TService.GetTopicInfoList(pageInfo)
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

// GetTopicPublic 不需要鉴权的题目管理接口
// @Tags Topic
// @Summary 不需要鉴权的题目管理接口
// @accept application/json
// @Produce application/json
// @Param data query TopicReq.TopicSearch true "分页获取题目管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /T/getTopicPublic [get]
func (TApi *TopicApi) GetTopicPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    TService.GetTopicPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的题目管理接口信息",
    }, "获取成功", c)
}
