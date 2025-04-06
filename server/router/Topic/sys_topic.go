package Topic

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TopicRouter struct {}

// InitTopicRouter 初始化 题目管理 路由信息
func (s *TopicRouter) InitTopicRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	TRouter := Router.Group("T").Use(middleware.OperationRecord())
	TRouterWithoutRecord := Router.Group("T")
	TRouterWithoutAuth := PublicRouter.Group("T")
	{
		TRouter.POST("createTopic", TApi.CreateTopic)   // 新建题目管理
		TRouter.DELETE("deleteTopic", TApi.DeleteTopic) // 删除题目管理
		TRouter.DELETE("deleteTopicByIds", TApi.DeleteTopicByIds) // 批量删除题目管理
		TRouter.PUT("updateTopic", TApi.UpdateTopic)    // 更新题目管理
	}
	{
		TRouterWithoutRecord.GET("findTopic", TApi.FindTopic)        // 根据ID获取题目管理
		TRouterWithoutRecord.GET("getTopicList", TApi.GetTopicList)  // 获取题目管理列表
	}
	{
	    TRouterWithoutAuth.GET("getTopicPublic", TApi.GetTopicPublic)  // 题目管理开放接口
	}
}
