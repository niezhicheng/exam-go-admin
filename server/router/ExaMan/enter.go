package ExaMan

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ExamManagementRouter }

var EmApi = api.ApiGroupApp.ExaManApiGroup.ExamManagementApi
