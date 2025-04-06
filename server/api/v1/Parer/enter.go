package Parer

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PaperApi }

var PService = service.ServiceGroupApp.ParerServiceGroup.PaperService
