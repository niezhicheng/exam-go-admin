package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/AnsweSituation"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ExaMan"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Examresults"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Parer"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ParerDetail"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Topic"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	TopicServiceGroup          Topic.ServiceGroup
	ParerServiceGroup          Parer.ServiceGroup
	ParerDetailServiceGroup    ParerDetail.ServiceGroup
	ExaManServiceGroup         ExaMan.ServiceGroup
	AnsweSituationServiceGroup AnsweSituation.ServiceGroup
	ExamresultsServiceGroup    Examresults.ServiceGroup
}
