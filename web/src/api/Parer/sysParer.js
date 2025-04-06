import service from '@/utils/request'
// @Tags Paper
// @Summary 创建试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Paper true "创建试卷管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /P/createPaper [post]
export const createPaper = (data) => {
  return service({
    url: '/P/createPaper',
    method: 'post',
    data
  })
}

// @Tags Paper
// @Summary 删除试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Paper true "删除试卷管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /P/deletePaper [delete]
export const deletePaper = (params) => {
  return service({
    url: '/P/deletePaper',
    method: 'delete',
    params
  })
}

// @Tags Paper
// @Summary 批量删除试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除试卷管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /P/deletePaper [delete]
export const deletePaperByIds = (params) => {
  return service({
    url: '/P/deletePaperByIds',
    method: 'delete',
    params
  })
}

// @Tags Paper
// @Summary 更新试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Paper true "更新试卷管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /P/updatePaper [put]
export const updatePaper = (data) => {
  return service({
    url: '/P/updatePaper',
    method: 'put',
    data
  })
}

// @Tags Paper
// @Summary 用id查询试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Paper true "用id查询试卷管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /P/findPaper [get]
export const findPaper = (params) => {
  return service({
    url: '/P/findPaper',
    method: 'get',
    params
  })
}

// @Tags Paper
// @Summary 分页获取试卷管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取试卷管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /P/getPaperList [get]
export const getPaperList = (params) => {
  return service({
    url: '/P/getPaperList',
    method: 'get',
    params
  })
}

// @Tags Paper
// @Summary 不需要鉴权的试卷管理接口
// @accept application/json
// @Produce application/json
// @Param data query ParerReq.PaperSearch true "分页获取试卷管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /P/getPaperPublic [get]
export const getPaperPublic = () => {
  return service({
    url: '/P/getPaperPublic',
    method: 'get',
  })
}


// @Tags Paper
// @Summary 分页获取试卷管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取试卷管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /P/getPaperList [get]
export const getPaperListAll = (params) => {
  return service({
    url: '/P/getPaperListAll',
    method: 'get',
    params
  })
}


// @Tags Paper
// @Summary 用id查询试卷管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Paper true "用id查询试卷管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /P/findPaper [get]
export const findPaperExam = (params) => {
  return service({
    url: '/P/findPaperExam',
    method: 'get',
    params
  })
}
