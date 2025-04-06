import service from '@/utils/request'
// @Tags ExamResults
// @Summary 创建考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamResults true "创建考试成绩"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Er/createExamResults [post]
export const createExamResults = (data) => {
  return service({
    url: '/Er/createExamResults',
    method: 'post',
    data
  })
}

// @Tags ExamResults
// @Summary 删除考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamResults true "删除考试成绩"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Er/deleteExamResults [delete]
export const deleteExamResults = (params) => {
  return service({
    url: '/Er/deleteExamResults',
    method: 'delete',
    params
  })
}

// @Tags ExamResults
// @Summary 批量删除考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除考试成绩"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Er/deleteExamResults [delete]
export const deleteExamResultsByIds = (params) => {
  return service({
    url: '/Er/deleteExamResultsByIds',
    method: 'delete',
    params
  })
}

// @Tags ExamResults
// @Summary 更新考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamResults true "更新考试成绩"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Er/updateExamResults [put]
export const updateExamResults = (data) => {
  return service({
    url: '/Er/updateExamResults',
    method: 'put',
    data
  })
}

// @Tags ExamResults
// @Summary 用id查询考试成绩
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExamResults true "用id查询考试成绩"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Er/findExamResults [get]
export const findExamResults = (params) => {
  return service({
    url: '/Er/findExamResults',
    method: 'get',
    params
  })
}

// @Tags ExamResults
// @Summary 分页获取考试成绩列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取考试成绩列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Er/getExamResultsList [get]
export const getExamResultsList = (params) => {
  return service({
    url: '/Er/getExamResultsList',
    method: 'get',
    params
  })
}

// @Tags ExamResults
// @Summary 不需要鉴权的考试成绩接口
// @accept application/json
// @Produce application/json
// @Param data query ExamresultsReq.ExamResultsSearch true "分页获取考试成绩列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Er/getExamResultsPublic [get]
export const getExamResultsPublic = () => {
  return service({
    url: '/Er/getExamResultsPublic',
    method: 'get',
  })
}
