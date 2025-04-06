package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct{}

// Author represents the author of a commit
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Commit represents a single commit in the push event
type Commit struct {
	ID        string   `json:"id"`
	Message   string   `json:"message"`
	Title     string   `json:"title"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
	Author    Author   `json:"author"`
	Added     []string `json:"added"`
	Modified  []string `json:"modified"`
	Removed   []string `json:"removed"`
}

// Repository represents the GitLab repository
type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

// Project represents the GitLab project
type Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}

// PushEvent represents the structure of a GitLab push event
type PushEvent struct {
	ObjectKind        string                 `json:"object_kind"`
	EventName         string                 `json:"event_name"`
	Before            string                 `json:"before"`
	After             string                 `json:"after"`
	Ref               string                 `json:"ref"`
	RefProtected      bool                   `json:"ref_protected"`
	CheckoutSHA       string                 `json:"checkout_sha"`
	Message           *string                `json:"message"` // 使用指针以处理可能为 null 的情况
	UserID            int                    `json:"user_id"`
	UserName          string                 `json:"user_name"`
	UserUsername      string                 `json:"user_username"`
	UserEmail         *string                `json:"user_email"` // 使用指针以处理可能为 null 的情况
	UserAvatar        string                 `json:"user_avatar"`
	ProjectID         int                    `json:"project_id"`
	Project           Project                `json:"project"`
	Commits           []Commit               `json:"commits"`
	TotalCommitsCount int                    `json:"total_commits_count"`
	PushOptions       map[string]interface{} `json:"push_options"` // 使用 interface{} 以支持任意 JSON 结构
	Repository        Repository             `json:"repository"`
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("ceshi", func(c *gin.Context) {
			// 读取请求体
			fmt.Println(c.Request.Header)
			fmt.Println(c.Request.Form)
			fmt.Println(c.Request.Body)
			// 获取原始请求体
			//body, err := ioutil.ReadAll(c.Request.Body)
			//if err != nil {
			//	c.JSON(http.StatusBadRequest, gin.H{"error": "Could not read request body"})
			//	return
			//}

			// 打印请求体
			//fmt.Println("Received body:", string(body))

			// 重新设置请求体，以便后续可以继续读取
			//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			// 打印请求头
			//fmt.Println("Request Headers:", c.Request.Header)

			// 绑定 JSON 数据到结构体
			var pushEvents PushEvent
			if err := c.ShouldBindJSON(&pushEvents); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 打印解析后的数据
			fmt.Printf("Parsed Push Event: %+v\n", pushEvents)

			//fmt.Println(c.Request)
			// Commit represents a single commit in the push event
			// Author represents the author of a commit
			//type Author struct {
			//	Name  string `json:"name"`
			//	Email string `json:"email"`
			//}
			//type Commit struct {
			//	ID        string `json:"id"`
			//	Message   string `json:"message"`
			//	Timestamp string `json:"timestamp"`
			//	URL       string `json:"url"`
			//	Author    Author `json:"author"`
			//}
			//
			//// Repository represents the GitLab repository
			//type Repository struct {
			//	Name            string `json:"name"`
			//	URL             string `json:"url"`
			//	Description     string `json:"description"`
			//	Homepage        string `json:"homepage"`
			//	GitHTTPURL      string `json:"git_http_url"`
			//	GitSSHURL       string `json:"git_ssh_url"`
			//	VisibilityLevel int    `json:"visibility_level"`
			//}
			//
			//// PushEvent represents the structure of a GitLab push event
			//type PushEvent struct {
			//	ObjectKind        string     `json:"object_kind"`
			//	EventName         string     `json:"event_name"`
			//	Before            string     `json:"before"`
			//	After             string     `json:"after"`
			//	Ref               string     `json:"ref"`
			//	CheckoutSHA       string     `json:"checkout_sha"`
			//	Message           string     `json:"message"`
			//	UserID            int        `json:"user_id"`
			//	UserName          string     `json:"user_name"`
			//	Repository        Repository `json:"repository"`
			//	Commits           []Commit   `json:"commits"`
			//	TotalCommitsCount int        `json:"total_commits_count"`
			//}
			//var pushEvent PushEvent
			//
			//// Bind the JSON to the PushEvent struct
			//if err := c.ShouldBindJSON(&pushEvent); err != nil {
			//	fmt.Println(err.Error(), "这是")
			//	c.JSON(400, gin.H{"error": err.Error()})
			//	return
			//}

			// Print the parsed data
			//fmt.Printf("Received Push Event: %+v\n", pushEvent)

			// Respond to the request
			c.JSON(200, gin.H{"status": "received"})
			//fmt.Println(context.Request)
		})
	}
	return baseRouter
}
