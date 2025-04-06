package ParerDetail

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PaperDetailRouter }

var DpApi = api.ApiGroupApp.ParerDetailApiGroup.PaperDetailApi
