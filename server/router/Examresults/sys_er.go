package Examresults

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExamResultsRouter struct {}

// InitExamResultsRouter 初始化 考试成绩 路由信息
func (s *ExamResultsRouter) InitExamResultsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	ErRouter := Router.Group("Er").Use(middleware.OperationRecord())
	ErRouterWithoutRecord := Router.Group("Er")
	ErRouterWithoutAuth := PublicRouter.Group("Er")
	{
		ErRouter.POST("createExamResults", ErApi.CreateExamResults)   // 新建考试成绩
		ErRouter.DELETE("deleteExamResults", ErApi.DeleteExamResults) // 删除考试成绩
		ErRouter.DELETE("deleteExamResultsByIds", ErApi.DeleteExamResultsByIds) // 批量删除考试成绩
		ErRouter.PUT("updateExamResults", ErApi.UpdateExamResults)    // 更新考试成绩
	}
	{
		ErRouterWithoutRecord.GET("findExamResults", ErApi.FindExamResults)        // 根据ID获取考试成绩
		ErRouterWithoutRecord.GET("getExamResultsList", ErApi.GetExamResultsList)  // 获取考试成绩列表
	}
	{
	    ErRouterWithoutAuth.GET("getExamResultsPublic", ErApi.GetExamResultsPublic)  // 考试成绩开放接口
	}
}
