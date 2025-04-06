// 自动生成模板Topic
package Topic

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 题目管理 结构体  Topic
type Topic struct {
	global.GVA_MODEL
	Subject   string  `json:"subject" form:"subject" gorm:"column:subject;comment:题目;"`              //题目
	Content   string  `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`    //内容
	TopicType uint    `json:"topicType" form:"topicType" gorm:"column:topic_type;comment:题目类型;"`     //题目类型 1单选 2多选 3判断 4填空
	Answer    string  `json:"answer" form:"answer" gorm:"column:answer;comment:答案;"`                 //答案
	Analysis  string  `json:"analysis" form:"analysis" gorm:"column:analysis;comment:解析;type:text;"` //解析
	Disable   bool    `json:"disable" form:"disable" gorm:"column:disable;comment:禁用;"`              //禁用
	Option    string  `json:"option" form:"option" gorm:"column:option;comment:选项;"`                 //选项
	Scope     float64 `json:"scope" form:"scope" gorm:"column:scope;comment:分数;"`                    // 分数
}

// TableName 题目管理 Topic自定义表名 topic
func (Topic) TableName() string {
	return "topic"
}
