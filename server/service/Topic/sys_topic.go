package Topic

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Topic"
    TopicReq "github.com/flipped-aurora/gin-vue-admin/server/model/Topic/request"
)

type TopicService struct {}
// CreateTopic 创建题目管理记录
// Author [yourname](https://github.com/yourname)
func (TService *TopicService) CreateTopic(T *Topic.Topic) (err error) {
	err = global.GVA_DB.Create(T).Error
	return err
}

// DeleteTopic 删除题目管理记录
// Author [yourname](https://github.com/yourname)
func (TService *TopicService)DeleteTopic(ID string) (err error) {
	err = global.GVA_DB.Delete(&Topic.Topic{},"id = ?",ID).Error
	return err
}

// DeleteTopicByIds 批量删除题目管理记录
// Author [yourname](https://github.com/yourname)
func (TService *TopicService)DeleteTopicByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Topic.Topic{},"id in ?",IDs).Error
	return err
}

// UpdateTopic 更新题目管理记录
// Author [yourname](https://github.com/yourname)
func (TService *TopicService)UpdateTopic(T Topic.Topic) (err error) {
	err = global.GVA_DB.Model(&Topic.Topic{}).Where("id = ?",T.ID).Updates(&T).Error
	return err
}

// GetTopic 根据ID获取题目管理记录
// Author [yourname](https://github.com/yourname)
func (TService *TopicService)GetTopic(ID string) (T Topic.Topic, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&T).Error
	return
}

// GetTopicInfoList 分页获取题目管理记录
// Author [yourname](https://github.com/yourname)
func (TService *TopicService)GetTopicInfoList(info TopicReq.TopicSearch) (list []Topic.Topic, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Topic.Topic{})
    var Ts []Topic.Topic
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Subject != "" {
        db = db.Where("subject LIKE ?","%"+ info.Subject+"%")
    }
    if info.TopicType != nil {
        db = db.Where("topic_type = ?",info.TopicType)
    }
    if info.Disable != nil {
        db = db.Where("disable = ?",info.Disable)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&Ts).Error
	return  Ts, total, err
}
func (TService *TopicService)GetTopicPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
