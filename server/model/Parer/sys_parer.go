// 自动生成模板Paper
package Parer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 试卷管理 结构体  Paper
type Paper struct {
	global.GVA_MODEL
	PaperName        string  `json:"paperName" form:"paperName" gorm:"column:paper_name;comment:试卷名称;"`                    //试卷名称
	PaperDescription string  `json:"paperDescription" form:"paperDescription" gorm:"column:paper_description;comment:描述;"` //描述
	Total            float64 `json:"total" form:"total" gorm:"column:total;comment:总分;"`                                   // 总分
}

// TableName 试卷管理 Paper自定义表名 paper
func (Paper) TableName() string {
	return "paper"
}
