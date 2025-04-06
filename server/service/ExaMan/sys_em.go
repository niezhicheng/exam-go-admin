package ExaMan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ExaMan"
	ExaManReq "github.com/flipped-aurora/gin-vue-admin/server/model/ExaMan/request"
	"gorm.io/gorm/clause"
)

type ExamManagementService struct{}

// CreateExamManagement 创建考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) CreateExamManagement(Em *ExaMan.ExamManagement) (err error) {
	err = global.GVA_DB.Create(Em).Error
	return err
}

// DeleteExamManagement 删除考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) DeleteExamManagement(ID string) (err error) {
	err = global.GVA_DB.Delete(&ExaMan.ExamManagement{}, "id = ?", ID).Error
	return err
}

// DeleteExamManagementByIds 批量删除考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) DeleteExamManagementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ExaMan.ExamManagement{}, "id in ?", IDs).Error
	return err
}

// UpdateExamManagement 更新考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) UpdateExamManagement(Em ExaMan.ExamManagement) (err error) {
	err = global.GVA_DB.Model(&ExaMan.ExamManagement{}).Where("id = ?", Em.ID).Updates(&Em).Error
	return err
}

// GetExamManagement 根据ID获取考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) GetExamManagement(ID string) (Em ExaMan.ExamManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Em).Error
	return
}

// GetExamManagementInfoList 分页获取考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) GetExamManagementInfoList(info ExaManReq.ExamManagementSearch) (list []ExaMan.ExamManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ExaMan.ExamManagement{}).Preload(clause.Associations)
	var Ems []ExaMan.ExamManagement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ExamName != "" {
		db = db.Where("exam_name LIKE ?", "%"+info.ExamName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Ems).Error
	return Ems, total, err
}
func (EmService *ExamManagementService) GetExamManagementPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetExamManagementInfoList 分页获取考试管理记录
// Author [yourname](https://github.com/yourname)
func (EmService *ExamManagementService) GetUserExamManagementInfoList(info ExaManReq.ExamManagementSearch) (list []ExaMan.ExamManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ExaMan.ExamManagement{}).Preload(clause.Associations)
	var Ems []ExaMan.ExamManagement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ExamName != "" {
		db = db.Where("exam_name LIKE ?", "%"+info.ExamName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Ems).Error
	return Ems, total, err
}
