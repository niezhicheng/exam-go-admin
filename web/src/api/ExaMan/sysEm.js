import service from '@/utils/request'
// @Tags ExamManagement
// @Summary 创建考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamManagement true "创建考试管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Em/createExamManagement [post]
export const createExamManagement = (data) => {
  return service({
    url: '/Em/createExamManagement',
    method: 'post',
    data
  })
}

// @Tags ExamManagement
// @Summary 删除考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamManagement true "删除考试管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Em/deleteExamManagement [delete]
export const deleteExamManagement = (params) => {
  return service({
    url: '/Em/deleteExamManagement',
    method: 'delete',
    params
  })
}

// @Tags ExamManagement
// @Summary 批量删除考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除考试管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Em/deleteExamManagement [delete]
export const deleteExamManagementByIds = (params) => {
  return service({
    url: '/Em/deleteExamManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags ExamManagement
// @Summary 更新考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamManagement true "更新考试管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Em/updateExamManagement [put]
export const updateExamManagement = (data) => {
  return service({
    url: '/Em/updateExamManagement',
    method: 'put',
    data
  })
}

// @Tags ExamManagement
// @Summary 用id查询考试管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExamManagement true "用id查询考试管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Em/findExamManagement [get]
export const findExamManagement = (params) => {
  return service({
    url: '/Em/findExamManagement',
    method: 'get',
    params
  })
}

// @Tags ExamManagement
// @Summary 分页获取考试管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取考试管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Em/getExamManagementList [get]
export const getExamManagementList = (params) => {
  return service({
    url: '/Em/getExamManagementList',
    method: 'get',
    params
  })
}

// @Tags ExamManagement
// @Summary 不需要鉴权的考试管理接口
// @accept application/json
// @Produce application/json
// @Param data query ExaManReq.ExamManagementSearch true "分页获取考试管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Em/getExamManagementPublic [get]
export const getExamManagementPublic = () => {
  return service({
    url: '/Em/getExamManagementPublic',
    method: 'get',
  })
}


// @Tags ExamManagement
// @Summary 分页获取考试管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取考试管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Em/getExamManagementList [get]
export const getUserExamManagementList = (params) => {
  return service({
    url: '/Em/getUserExamManagementList',
    method: 'get',
    params
  })
}

