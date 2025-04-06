package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/AnsweSituation"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ExaMan"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Examresults"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Parer"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ParerDetail"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Topic"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	TopicApiGroup          Topic.ApiGroup
	ParerApiGroup          Parer.ApiGroup
	ParerDetailApiGroup    ParerDetail.ApiGroup
	ExaManApiGroup         ExaMan.ApiGroup
	AnsweSituationApiGroup AnsweSituation.ApiGroup
	ExamresultsApiGroup    Examresults.ApiGroup
}
