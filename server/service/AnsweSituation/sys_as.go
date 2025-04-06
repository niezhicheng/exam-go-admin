package AnsweSituation

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/AnsweSituation"
	AnsweSituationReq "github.com/flipped-aurora/gin-vue-admin/server/model/AnsweSituation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Examresults"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Topic"
	"gorm.io/gorm"
)

type AnsweringSituationService struct{}

// CreateAnsweringSituation 创建答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) CreateAnsweringSituation(As *AnsweSituation.AnsweringSituation) (err error) {
	err = global.GVA_DB.Create(As).Error
	return err
}

// DeleteAnsweringSituation 删除答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) DeleteAnsweringSituation(ID string) (err error) {
	err = global.GVA_DB.Delete(&AnsweSituation.AnsweringSituation{}, "id = ?", ID).Error
	return err
}

// DeleteAnsweringSituationByIds 批量删除答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) DeleteAnsweringSituationByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]AnsweSituation.AnsweringSituation{}, "id in ?", IDs).Error
	return err
}

// UpdateAnsweringSituation 更新答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) UpdateAnsweringSituation(As AnsweSituation.AnsweringSituation) (err error) {
	err = global.GVA_DB.Model(&AnsweSituation.AnsweringSituation{}).Where("id = ?", As.ID).Updates(&As).Error
	return err
}

// GetAnsweringSituation 根据ID获取答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) GetAnsweringSituation(ID string) (As AnsweSituation.AnsweringSituation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&As).Error
	return
}

// GetAnsweringSituationInfoList 分页获取答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) GetAnsweringSituationInfoList(info AnsweSituationReq.AnsweringSituationSearch) (list []AnsweSituation.AnsweringSituation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&AnsweSituation.AnsweringSituation{})
	var Ass []AnsweSituation.AnsweringSituation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.ExamId != nil {
		db = db.Where("exam_id = ?", info.ExamId)
	}
	if info.TitleId != nil {
		db = db.Where("title_id = ?", info.TitleId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Ass).Error
	return Ass, total, err
}
func (AsService *AnsweringSituationService) GetAnsweringSituationPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetAnsweringSituation 根据ID获取答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) GetExamAnsweringSituation(ID string, userid uint) (resp []uint, err error) {
	var As []AnsweSituation.AnsweringSituation
	err = global.GVA_DB.Where("user_id = ? and exam_id = ?", userid, ID).Find(&As).Error
	for _, a := range As {
		resp = append(resp, a.TitleId)
	}
	if len(resp) == 0 {
		resp = []uint{}
	}
	return
}

// CreateAnsweringSituation 创建答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) CreateUserAnsweringSituation(As *AnsweSituation.AnsweringSituation) (err error) {
	// 开始事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚
		}
	}()
	var existingSituation AnsweSituation.AnsweringSituation
	// 查询是否已存在相同的记录
	if err := tx.Where("user_id = ? AND exam_id = ? AND title_id = ?", As.UserId, As.ExamId, As.TitleId).First(&existingSituation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果没有找到，则插入新记录
			if err := tx.Create(&As).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err // 查询错误
		}
	} else {
		// 如果找到，则更新记录
		existingSituation.Answer = As.Answer
		if err := tx.Save(&existingSituation).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return err
}

// CreateAnsweringSituation 创建答题情况记录
// Author [yourname](https://github.com/yourname)
func (AsService *AnsweringSituationService) SubmitUserAnsweringSituation(As *AnsweSituation.AnsweringSituation) (err error) {
	var exam Examresults.ExamResults
	err = global.GVA_DB.Model(&Examresults.ExamResults{}).Where("user_id = ? and exam_id = ?", As.UserId, As.ExamId).First(&exam).Error
	if err == nil {
		return errors.New("不可重复提交试卷")
	}
	var answer []AnsweSituation.AnsweringSituation
	global.GVA_DB.Model(&AnsweSituation.AnsweringSituation{}).Where("user_id = ? and exam_id = ?", As.UserId, As.ExamId).Find(&answer)
	var topic []Topic.Topic
	global.GVA_DB.Model(&Topic.Topic{}).Find(&topic)
	var scope float64 = 0
	for _, t := range topic {
		for _, situation := range answer {
			if t.ID == situation.TitleId {
				if t.Answer == situation.Answer {
					scope = scope + t.Scope
				}
			}
		}
	}
	err = global.GVA_DB.Create(&Examresults.ExamResults{
		UserId:     As.UserId,
		ExamId:     As.ExamId,
		FinalScore: scope,
	}).Error
	return err
}
