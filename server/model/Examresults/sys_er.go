// 自动生成模板ExamResults
package Examresults

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 考试成绩 结构体  ExamResults
type ExamResults struct {
	global.GVA_MODEL
	UserId     uint    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`             //用户id
	ExamId     uint    `json:"examId" form:"examId" gorm:"column:exam_id;comment:考试id;"`             //考试id
	FinalScore float64 `json:"finalScore" form:"finalScore" gorm:"column:final_score;comment:最终成绩;"` //最终成绩
}

// TableName 考试成绩 ExamResults自定义表名 exam_results
func (ExamResults) TableName() string {
	return "exam_results"
}
