package ParerDetail

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PaperDetailApi }

var DpService = service.ServiceGroupApp.ParerDetailServiceGroup.PaperDetailService
