import service from '@/utils/request'
// @Tags PaperDetail
// @Summary 创建试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperDetail true "创建试卷详细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /Dp/createPaperDetail [post]
export const createPaperDetail = (data) => {
  return service({
    url: '/Dp/createPaperDetail',
    method: 'post',
    data
  })
}

// @Tags PaperDetail
// @Summary 删除试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperDetail true "删除试卷详细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Dp/deletePaperDetail [delete]
export const deletePaperDetail = (params) => {
  return service({
    url: '/Dp/deletePaperDetail',
    method: 'delete',
    params
  })
}

// @Tags PaperDetail
// @Summary 批量删除试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除试卷详细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Dp/deletePaperDetail [delete]
export const deletePaperDetailByIds = (params) => {
  return service({
    url: '/Dp/deletePaperDetailByIds',
    method: 'delete',
    params
  })
}

// @Tags PaperDetail
// @Summary 更新试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperDetail true "更新试卷详细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Dp/updatePaperDetail [put]
export const updatePaperDetail = (data) => {
  return service({
    url: '/Dp/updatePaperDetail',
    method: 'put',
    data
  })
}

// @Tags PaperDetail
// @Summary 用id查询试卷详细
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PaperDetail true "用id查询试卷详细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Dp/findPaperDetail [get]
export const findPaperDetail = (params) => {
  return service({
    url: '/Dp/findPaperDetail',
    method: 'get',
    params
  })
}

// @Tags PaperDetail
// @Summary 分页获取试卷详细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取试卷详细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dp/getPaperDetailList [get]
export const getPaperDetailList = (params) => {
  return service({
    url: '/Dp/getPaperDetailList',
    method: 'get',
    params
  })
}

// @Tags PaperDetail
// @Summary 不需要鉴权的试卷详细接口
// @accept application/json
// @Produce application/json
// @Param data query ParerDetailReq.PaperDetailSearch true "分页获取试卷详细列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Dp/getPaperDetailPublic [get]
export const getPaperDetailPublic = () => {
  return service({
    url: '/Dp/getPaperDetailPublic',
    method: 'get',
  })
}


// @Tags PaperDetail
// @Summary 分页获取试卷详细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取试卷详细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dp/getPaperDetailList [get]
export const FindPaperDetailAll = (params) => {
  return service({
    url: '/Dp/findPaperDetailAll',
    method: 'get',
    params
  })
}


// @Tags PaperDetail
// @Summary 分页获取试卷详细列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取试卷详细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Dp/getPaperDetailList [get]
export const GetQuestionListTableData = (params) => {
  return service({
    url: '/Dp/findPaperNoDetailAll',
    method: 'get',
    params
  })
}
