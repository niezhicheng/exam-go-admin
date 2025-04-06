import service from '@/utils/request'
// @Tags Topic
// @Summary 创建题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Topic true "创建题目管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /T/createTopic [post]
export const createTopic = (data) => {
  return service({
    url: '/T/createTopic',
    method: 'post',
    data
  })
}

// @Tags Topic
// @Summary 删除题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Topic true "删除题目管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /T/deleteTopic [delete]
export const deleteTopic = (params) => {
  return service({
    url: '/T/deleteTopic',
    method: 'delete',
    params
  })
}

// @Tags Topic
// @Summary 批量删除题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除题目管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /T/deleteTopic [delete]
export const deleteTopicByIds = (params) => {
  return service({
    url: '/T/deleteTopicByIds',
    method: 'delete',
    params
  })
}

// @Tags Topic
// @Summary 更新题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Topic true "更新题目管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /T/updateTopic [put]
export const updateTopic = (data) => {
  return service({
    url: '/T/updateTopic',
    method: 'put',
    data
  })
}

// @Tags Topic
// @Summary 用id查询题目管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Topic true "用id查询题目管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /T/findTopic [get]
export const findTopic = (params) => {
  return service({
    url: '/T/findTopic',
    method: 'get',
    params
  })
}

// @Tags Topic
// @Summary 分页获取题目管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取题目管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /T/getTopicList [get]
export const getTopicList = (params) => {
  return service({
    url: '/T/getTopicList',
    method: 'get',
    params
  })
}

// @Tags Topic
// @Summary 不需要鉴权的题目管理接口
// @accept application/json
// @Produce application/json
// @Param data query TopicReq.TopicSearch true "分页获取题目管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /T/getTopicPublic [get]
export const getTopicPublic = () => {
  return service({
    url: '/T/getTopicPublic',
    method: 'get',
  })
}
