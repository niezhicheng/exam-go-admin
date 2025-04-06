package AnsweSituation

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AnsweringSituationRouter }

var AsApi = api.ApiGroupApp.AnsweSituationApiGroup.AnsweringSituationApi
