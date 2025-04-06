package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ExamResultsSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    UserId  *int `json:"userId" form:"userId" `
    ExamId  *int `json:"examId" form:"examId" `
    request.PageInfo
}
