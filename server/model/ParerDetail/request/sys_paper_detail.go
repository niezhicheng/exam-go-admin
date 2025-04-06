package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PaperDetailSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PaperId        uint       `json:"paperId" form:"paperId" `
	TestId         uint       `json:"testId" form:"testId" `
	request.PageInfo
}
