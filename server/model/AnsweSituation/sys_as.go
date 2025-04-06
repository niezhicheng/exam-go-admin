// 自动生成模板AnsweringSituation
package AnsweSituation

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 答题情况 结构体  AnsweringSituation
type AnsweringSituation struct {
	global.GVA_MODEL
	UserId  uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`    //用户id
	ExamId  uint   `json:"examId" form:"examId" gorm:"column:exam_id;comment:考试id;"`    //考试id
	TitleId uint   `json:"titleId" form:"titleId" gorm:"column:title_id;comment:题目id;"` //题目id
	Answer  string `json:"answer" form:"answer" gorm:"column:answer;comment:答案;"`       //答案
}

// TableName 答题情况 AnsweringSituation自定义表名 answering_situation
func (AnsweringSituation) TableName() string {
	return "answering_situation"
}
