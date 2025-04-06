package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/AnsweSituation"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ExaMan"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Examresults"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Parer"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ParerDetail"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Topic"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(Topic.Topic{}, Parer.Paper{}, ParerDetail.PaperDetail{}, ExaMan.ExamManagement{}, AnsweSituation.AnsweringSituation{}, Examresults.ExamResults{})
	if err != nil {
		return err
	}
	return nil
}
