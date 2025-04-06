package Parer

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PaperRouter }

var PApi = api.ApiGroupApp.ParerApiGroup.PaperApi
