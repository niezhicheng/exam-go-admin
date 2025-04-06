package ParerDetail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaperDetailRouter struct{}

// InitPaperDetailRouter 初始化 试卷详细 路由信息
func (s *PaperDetailRouter) InitPaperDetailRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	DpRouter := Router.Group("Dp").Use(middleware.OperationRecord())
	DpRouterWithoutRecord := Router.Group("Dp")
	DpRouterWithoutAuth := PublicRouter.Group("Dp")
	{
		DpRouter.POST("createPaperDetail", DpApi.CreatePaperDetail)             // 新建试卷详细
		DpRouter.DELETE("deletePaperDetail", DpApi.DeletePaperDetail)           // 删除试卷详细
		DpRouter.DELETE("deletePaperDetailByIds", DpApi.DeletePaperDetailByIds) // 批量删除试卷详细
		DpRouter.PUT("updatePaperDetail", DpApi.UpdatePaperDetail)              // 更新试卷详细
	}
	{
		DpRouterWithoutRecord.GET("findPaperDetail", DpApi.FindPaperDetail)           // 根据ID获取试卷详细
		DpRouterWithoutRecord.GET("findPaperDetailAll", DpApi.FindPaperDetailAll)     // 根据试卷id获取试题列表
		DpRouterWithoutRecord.GET("findPaperNoDetailAll", DpApi.FindPaperNoDetailAll) // 根据试卷id获取没有的试题列表
		DpRouterWithoutRecord.GET("getPaperDetailList", DpApi.GetPaperDetailList)     // 获取试卷详细列表
	}
	{
		DpRouterWithoutAuth.GET("getPaperDetailPublic", DpApi.GetPaperDetailPublic) // 试卷详细开放接口
	}
}
