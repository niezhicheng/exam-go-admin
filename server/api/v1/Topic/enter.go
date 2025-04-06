package Topic

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ TopicApi }

var TService = service.ServiceGroupApp.TopicServiceGroup.TopicService
