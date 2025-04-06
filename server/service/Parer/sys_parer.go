package Parer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ExaMan"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Parer"
	ParerReq "github.com/flipped-aurora/gin-vue-admin/server/model/Parer/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ParerDetail"
	"gorm.io/gorm/clause"
)

type PaperService struct{}

// CreatePaper 创建试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) CreatePaper(P *Parer.Paper) (err error) {
	err = global.GVA_DB.Create(P).Error
	return err
}

// DeletePaper 删除试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) DeletePaper(ID string) (err error) {
	err = global.GVA_DB.Delete(&Parer.Paper{}, "id = ?", ID).Error
	return err
}

// DeletePaperByIds 批量删除试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) DeletePaperByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Parer.Paper{}, "id in ?", IDs).Error
	return err
}

// UpdatePaper 更新试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) UpdatePaper(P Parer.Paper) (err error) {
	err = global.GVA_DB.Model(&Parer.Paper{}).Where("id = ?", P.ID).Updates(&P).Error
	return err
}

// GetPaper 根据ID获取试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) GetPaper(ID string) (P Parer.Paper, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&P).Error
	return
}

// GetPaperInfoList 分页获取试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) GetPaperInfoList(info ParerReq.PaperSearch) (list []Parer.Paper, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Parer.Paper{})
	var Ps []Parer.Paper
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PaperName != "" {
		db = db.Where("paper_name LIKE ?", "%"+info.PaperName+"%")
	}
	if info.PaperDescription != "" {
		db = db.Where("paper_description LIKE ?", "%"+info.PaperDescription+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Ps).Error
	return Ps, total, err
}
func (PService *PaperService) GetPaperPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetPaperInfoList 分页获取试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) GetPaperInfoListAll(info ParerReq.PaperSearch) (list []Parer.Paper, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&Parer.Paper{})
	var Ps []Parer.Paper
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&Ps).Error
	return Ps, total, err
}

// GetPaper 根据ID获取试卷管理记录
// Author [yourname](https://github.com/yourname)
func (PService *PaperService) GetPaperExam(ID string) (P []ParerDetail.PaperDetail, err error) {
	var exam ExaMan.ExamManagement
	err = global.GVA_DB.Where("id = ?", ID).First(&exam).Error
	if err != nil {
		return []ParerDetail.PaperDetail{}, err
	}
	var detail []ParerDetail.PaperDetail
	err = global.GVA_DB.Model(&ParerDetail.PaperDetail{}).Where("paper_id = ?", exam.PaperId).Preload(clause.Associations).Find(&detail).Error
	if err != nil {
		return []ParerDetail.PaperDetail{}, err
	}
	//for i, paperDetail := range detail {
	for i, _ := range detail {
		detail[i].TestModel.Answer = ""
		detail[i].TestModel.Analysis = ""
	}
	return detail, nil
}
