package ExaMan

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ExamManagementApi }

var EmService = service.ServiceGroupApp.ExaManServiceGroup.ExamManagementService
