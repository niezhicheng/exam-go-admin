// 自动生成模板PaperDetail
package ParerDetail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Topic"
)

// 试卷详细 结构体  PaperDetail
type PaperDetail struct {
	global.GVA_MODEL
	PaperId   uint        `json:"paperId" form:"paperId" gorm:"column:paper_id;comment:试卷id;" binding:"required"` //试卷id
	TestId    uint        `json:"testId" form:"testId" gorm:"column:test_id;comment:试题id;" binding:"required"`    //试题id
	TestModel Topic.Topic `json:"testModel" gorm:"foreignKey:ID;references:TestId;comment:试题信息;"`                 // 试题信息
	Scope     float64     `json:"scope" form:"scope" gorm:"column:scope;comment:分数;" binding:"required"`          // 分数
}

// TableName 试卷详细 PaperDetail自定义表名 paper_detail
func (PaperDetail) TableName() string {
	return "paper_detail"
}
