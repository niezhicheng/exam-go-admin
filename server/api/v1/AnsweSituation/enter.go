package AnsweSituation

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AnsweringSituationApi }

var AsService = service.ServiceGroupApp.AnsweSituationServiceGroup.AnsweringSituationService
