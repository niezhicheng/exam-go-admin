package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		TopicRouter := router.RouterGroupApp.Topic
		TopicRouter.InitTopicRouter(privateGroup, publicGroup)
	}
	{
		ParerRouter := router.RouterGroupApp.Parer
		ParerRouter.InitPaperRouter(privateGroup, publicGroup)
	}
	{
		ParerDetailRouter := router.RouterGroupApp.ParerDetail
		ParerDetailRouter.InitPaperDetailRouter(privateGroup, publicGroup)
	}
	{
		ExaManRouter := router.RouterGroupApp.ExaMan
		ExaManRouter.InitExamManagementRouter(privateGroup, publicGroup)
	}
	{
		AnsweSituationRouter := router.RouterGroupApp.AnsweSituation
		AnsweSituationRouter.InitAnsweringSituationRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		ExamresultsRouter := router.RouterGroupApp.Examresults
		ExamresultsRouter.InitExamResultsRouter(privateGroup, publicGroup)
	}
}
