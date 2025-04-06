package Topic

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ TopicRouter }

var TApi = api.ApiGroupApp.TopicApiGroup.TopicApi
