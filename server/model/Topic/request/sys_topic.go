package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TopicSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Subject  string `json:"subject" form:"subject" `
    TopicType  *int `json:"topicType" form:"topicType" `
    Disable  *bool `json:"disable" form:"disable" `
    request.PageInfo
}
