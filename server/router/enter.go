package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/AnsweSituation"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ExaMan"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Examresults"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Parer"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ParerDetail"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Topic"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	Topic          Topic.RouterGroup
	Parer          Parer.RouterGroup
	ParerDetail    ParerDetail.RouterGroup
	ExaMan         ExaMan.RouterGroup
	AnsweSituation AnsweSituation.RouterGroup
	Examresults    Examresults.RouterGroup
}
