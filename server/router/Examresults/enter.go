package Examresults

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ExamResultsRouter }

var ErApi = api.ApiGroupApp.ExamresultsApiGroup.ExamResultsApi
