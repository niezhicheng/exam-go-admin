package AnsweSituation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnsweringSituationRouter struct{}

// InitAnsweringSituationRouter 初始化 答题情况 路由信息
func (s *AnsweringSituationRouter) InitAnsweringSituationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	AsRouter := Router.Group("As").Use(middleware.OperationRecord())
	AsRouterWithoutRecord := Router.Group("As")
	AsRouterWithoutAuth := PublicRouter.Group("As")
	{
		AsRouter.POST("createAnsweringSituation", AsApi.CreateAnsweringSituation)             // 新建答题情况
		AsRouter.DELETE("deleteAnsweringSituation", AsApi.DeleteAnsweringSituation)           // 删除答题情况
		AsRouter.DELETE("deleteAnsweringSituationByIds", AsApi.DeleteAnsweringSituationByIds) // 批量删除答题情况
		AsRouter.PUT("updateAnsweringSituation", AsApi.UpdateAnsweringSituation)              // 更新答题情况
		AsRouter.POST("createUserAnsweringSituation", AsApi.CreateUserAnsweringSituation)     // 用户提交考试答题
		AsRouter.POST("submitUserAnsweringSituation", AsApi.SumbitUserAnsweringSituation)     // 用户提交考试赛卷
	}
	{
		AsRouterWithoutRecord.GET("findAnsweringSituation", AsApi.FindAnsweringSituation)         // 根据ID获取答题情况
		AsRouterWithoutRecord.GET("findExamAnsweringSituation", AsApi.FindExamAnsweringSituation) // 根据考试获取用户考试答题情况
		AsRouterWithoutRecord.GET("getAnsweringSituationList", AsApi.GetAnsweringSituationList)   // 获取答题情况列表
	}
	{
		AsRouterWithoutAuth.GET("getAnsweringSituationPublic", AsApi.GetAnsweringSituationPublic) // 答题情况开放接口
	}
}
