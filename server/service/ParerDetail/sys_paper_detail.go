package ParerDetail

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ParerDetail"
	ParerDetailReq "github.com/flipped-aurora/gin-vue-admin/server/model/ParerDetail/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Topic"
	"gorm.io/gorm/clause"
)

type PaperDetailService struct{}

// CreatePaperDetail 创建试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) CreatePaperDetail(Dp *ParerDetail.PaperDetail) (err error) {
	err = global.GVA_DB.Create(Dp).Error
	return err
}

// DeletePaperDetail 删除试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) DeletePaperDetail(ID string) (err error) {
	err = global.GVA_DB.Delete(&ParerDetail.PaperDetail{}, "id = ?", ID).Error
	return err
}

// DeletePaperDetailByIds 批量删除试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) DeletePaperDetailByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]ParerDetail.PaperDetail{}, "id in ?", IDs).Error
	return err
}

// UpdatePaperDetail 更新试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) UpdatePaperDetail(Dp ParerDetail.PaperDetail) (err error) {
	err = global.GVA_DB.Model(&ParerDetail.PaperDetail{}).Where("id = ?", Dp.ID).Updates(&Dp).Error
	return err
}

// GetPaperDetail 根据ID获取试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) GetPaperDetail(ID string) (Dp ParerDetail.PaperDetail, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Dp).Error
	return
}

// GetPaperDetailInfoList 分页获取试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) GetPaperDetailInfoList(info ParerDetailReq.PaperDetailSearch) (list []ParerDetail.PaperDetail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ParerDetail.PaperDetail{})
	var Dps []ParerDetail.PaperDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PaperId != 0 {
		db = db.Where("paper_id = ?", info.PaperId)
	}
	if info.TestId != 0 {
		db = db.Where("test_id = ?", info.TestId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Dps).Error
	return Dps, total, err
}
func (DpService *PaperDetailService) GetPaperDetailPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetPaperDetailInfoList 分页获取试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) GetPaperInfoList(info ParerDetailReq.PaperDetailSearch) (list []ParerDetail.PaperDetail, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&ParerDetail.PaperDetail{}).Preload(clause.Associations)
	var Dps []ParerDetail.PaperDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PaperId != 0 {
		db = db.Where("paper_id = ?", info.PaperId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&Dps).Error
	return Dps, total, err
}

// GetPaperDetailInfoList 分页获取试卷详细记录
// Author [yourname](https://github.com/yourname)
func (DpService *PaperDetailService) GetPaperNoInfoList(info ParerDetailReq.PaperDetailSearch) (list []Topic.Topic, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	var Dps []ParerDetail.PaperDetail
	global.GVA_DB.Model(&ParerDetail.PaperDetail{}).Where("paper_id = ?", info.PaperId).Find(&Dps)

	// 如果有条件搜索 下方会自动创建搜索语句

	var testId []uint
	for _, dp := range Dps {
		testId = append(testId, dp.TestId)
	}
	var topic []Topic.Topic
	db := global.GVA_DB.Model(&Topic.Topic{})
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if len(testId) != 0 {
		db = db.Not("id in ?", testId)
	}
	db = db.Find(&topic).Count(&total)

	return topic, total, err
}
