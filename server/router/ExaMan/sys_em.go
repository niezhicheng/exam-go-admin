package ExaMan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExamManagementRouter struct{}

// InitExamManagementRouter 初始化 考试管理 路由信息
func (s *ExamManagementRouter) InitExamManagementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	EmRouter := Router.Group("Em").Use(middleware.OperationRecord())
	EmRouterWithoutRecord := Router.Group("Em")
	EmRouterWithoutAuth := PublicRouter.Group("Em")
	{
		EmRouter.POST("createExamManagement", EmApi.CreateExamManagement)             // 新建考试管理
		EmRouter.DELETE("deleteExamManagement", EmApi.DeleteExamManagement)           // 删除考试管理
		EmRouter.DELETE("deleteExamManagementByIds", EmApi.DeleteExamManagementByIds) // 批量删除考试管理
		EmRouter.PUT("updateExamManagement", EmApi.UpdateExamManagement)              // 更新考试管理
	}
	{
		EmRouterWithoutRecord.GET("findExamManagement", EmApi.FindExamManagement)               // 根据ID获取考试管理
		EmRouterWithoutRecord.GET("getExamManagementList", EmApi.GetExamManagementList)         // 获取考试管理列表
		EmRouterWithoutRecord.GET("getUserExamManagementList", EmApi.GetUserExamManagementList) // 学生获取考试管理列表
	}
	{
		EmRouterWithoutAuth.GET("getExamManagementPublic", EmApi.GetExamManagementPublic) // 考试管理开放接口
	}
}
