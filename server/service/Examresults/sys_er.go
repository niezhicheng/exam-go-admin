package Examresults

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Examresults"
    ExamresultsReq "github.com/flipped-aurora/gin-vue-admin/server/model/Examresults/request"
)

type ExamResultsService struct {}
// CreateExamResults 创建考试成绩记录
// Author [yourname](https://github.com/yourname)
func (ErService *ExamResultsService) CreateExamResults(Er *Examresults.ExamResults) (err error) {
	err = global.GVA_DB.Create(Er).Error
	return err
}

// DeleteExamResults 删除考试成绩记录
// Author [yourname](https://github.com/yourname)
func (ErService *ExamResultsService)DeleteExamResults(ID string) (err error) {
	err = global.GVA_DB.Delete(&Examresults.ExamResults{},"id = ?",ID).Error
	return err
}

// DeleteExamResultsByIds 批量删除考试成绩记录
// Author [yourname](https://github.com/yourname)
func (ErService *ExamResultsService)DeleteExamResultsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Examresults.ExamResults{},"id in ?",IDs).Error
	return err
}

// UpdateExamResults 更新考试成绩记录
// Author [yourname](https://github.com/yourname)
func (ErService *ExamResultsService)UpdateExamResults(Er Examresults.ExamResults) (err error) {
	err = global.GVA_DB.Model(&Examresults.ExamResults{}).Where("id = ?",Er.ID).Updates(&Er).Error
	return err
}

// GetExamResults 根据ID获取考试成绩记录
// Author [yourname](https://github.com/yourname)
func (ErService *ExamResultsService)GetExamResults(ID string) (Er Examresults.ExamResults, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Er).Error
	return
}

// GetExamResultsInfoList 分页获取考试成绩记录
// Author [yourname](https://github.com/yourname)
func (ErService *ExamResultsService)GetExamResultsInfoList(info ExamresultsReq.ExamResultsSearch) (list []Examresults.ExamResults, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Examresults.ExamResults{})
    var Ers []Examresults.ExamResults
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
    if info.ExamId != nil {
        db = db.Where("exam_id = ?",info.ExamId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&Ers).Error
	return  Ers, total, err
}
func (ErService *ExamResultsService)GetExamResultsPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
