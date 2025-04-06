package Examresults

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ExamResultsApi }

var ErService = service.ServiceGroupApp.ExamresultsServiceGroup.ExamResultsService
