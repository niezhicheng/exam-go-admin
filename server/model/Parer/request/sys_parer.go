package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PaperSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    PaperName  string `json:"paperName" form:"paperName" `
    PaperDescription  string `json:"paperDescription" form:"paperDescription" `
    request.PageInfo
}
