package Parer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaperRouter struct{}

// InitPaperRouter 初始化 试卷管理 路由信息
func (s *PaperRouter) InitPaperRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PRouter := Router.Group("P").Use(middleware.OperationRecord())
	PRouterWithoutRecord := Router.Group("P")
	PRouterWithoutAuth := PublicRouter.Group("P")
	{
		PRouter.POST("createPaper", PApi.CreatePaper)             // 新建试卷管理
		PRouter.DELETE("deletePaper", PApi.DeletePaper)           // 删除试卷管理
		PRouter.DELETE("deletePaperByIds", PApi.DeletePaperByIds) // 批量删除试卷管理
		PRouter.PUT("updatePaper", PApi.UpdatePaper)              // 更新试卷管理
	}
	{
		PRouterWithoutRecord.GET("findPaper", PApi.FindPaper)             // 根据ID获取试卷管理
		PRouterWithoutRecord.GET("findPaperExam", PApi.FindPaperExam)     // 根据考试获取关联的试卷和考题
		PRouterWithoutRecord.GET("getPaperList", PApi.GetPaperList)       // 获取试卷管理列表
		PRouterWithoutRecord.GET("getPaperListAll", PApi.GetPaperListAll) // 获取试卷全部管理列表
	}
	{
		PRouterWithoutAuth.GET("getPaperPublic", PApi.GetPaperPublic) // 试卷管理开放接口
	}
}
