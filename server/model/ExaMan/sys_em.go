// 自动生成模板ExamManagement
package ExaMan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Parer"
	"time"
)

// 考试管理 结构体  ExamManagement
type ExamManagement struct {
	global.GVA_MODEL
	ExamName              string      `json:"examName" form:"examName" gorm:"column:exam_name;comment:考试名称;"`                                           //考试名称
	ExamCoverPicture      string      `json:"examCoverPicture" form:"examCoverPicture" gorm:"column:exam_cover_picture;comment:考试封面图;"`                 //考试封面图
	ExamDescription       string      `json:"examDescription" form:"examDescription" gorm:"column:exam_description;comment:考试描述;"`                      //考试描述
	RegistrationRequired  bool        `json:"registrationRequired" form:"registrationRequired" gorm:"column:registration_required;comment:是否需要报名;"`     //是否需要报名
	RegistrationStartTime *time.Time  `json:"registrationStartTime" form:"registrationStartTime" gorm:"column:registration_start_time;comment:报名开始时间;"` //报名开始时间
	RegistrationDeadline  *time.Time  `json:"registrationDeadline" form:"registrationDeadline" gorm:"column:registration_deadline;comment:报名结束时间;"`     //报名结束时间
	ExamStartTime         *time.Time  `json:"examStartTime" form:"examStartTime" gorm:"column:exam_start_time;comment:考试开始时间;"`                         //考试开始时间
	ExamEndTime           *time.Time  `json:"examEndTime" form:"examEndTime" gorm:"column:exam_end_time;comment:考试结束时间;"`                               //考试结束时间
	PaperId               uint        `json:"paper_id" form:"paper_id" gorm:"column:paper_id;comment:关联试卷;"`                                            // 关联试卷
	PaperModel            Parer.Paper `json:"paper_model" gorm:"foreignKey:ID;references:PaperId;"`                                                     // 关联试卷的model
	Status                uint        `json:"status" gorm:"column:status;comment:状态;"`                                                                  // 状态 1 未开始 2 报名中 3 进行中 4 已结束
	Scope                 float64     `json:"scope" gorm:"column:scope;comment:总分;"`                                                                    //总分
}

// TableName 考试管理 ExamManagement自定义表名 exam_management
func (ExamManagement) TableName() string {
	return "exam_management"
}
