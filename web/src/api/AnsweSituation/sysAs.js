import service from '@/utils/request'
// @Tags AnsweringSituation
// @Summary 创建答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnsweringSituation true "创建答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /As/createAnsweringSituation [post]
export const createAnsweringSituation = (data) => {
  return service({
    url: '/As/createAnsweringSituation',
    method: 'post',
    data
  })
}

// @Tags AnsweringSituation
// @Summary 删除答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnsweringSituation true "删除答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /As/deleteAnsweringSituation [delete]
export const deleteAnsweringSituation = (params) => {
  return service({
    url: '/As/deleteAnsweringSituation',
    method: 'delete',
    params
  })
}

// @Tags AnsweringSituation
// @Summary 批量删除答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /As/deleteAnsweringSituation [delete]
export const deleteAnsweringSituationByIds = (params) => {
  return service({
    url: '/As/deleteAnsweringSituationByIds',
    method: 'delete',
    params
  })
}

// @Tags AnsweringSituation
// @Summary 更新答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnsweringSituation true "更新答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /As/updateAnsweringSituation [put]
export const updateAnsweringSituation = (data) => {
  return service({
    url: '/As/updateAnsweringSituation',
    method: 'put',
    data
  })
}

// @Tags AnsweringSituation
// @Summary 用id查询答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AnsweringSituation true "用id查询答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /As/findAnsweringSituation [get]
export const findAnsweringSituation = (params) => {
  return service({
    url: '/As/findAnsweringSituation',
    method: 'get',
    params
  })
}

// @Tags AnsweringSituation
// @Summary 分页获取答题情况列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取答题情况列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /As/getAnsweringSituationList [get]
export const getAnsweringSituationList = (params) => {
  return service({
    url: '/As/getAnsweringSituationList',
    method: 'get',
    params
  })
}

// @Tags AnsweringSituation
// @Summary 不需要鉴权的答题情况接口
// @accept application/json
// @Produce application/json
// @Param data query AnsweSituationReq.AnsweringSituationSearch true "分页获取答题情况列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /As/getAnsweringSituationPublic [get]
export const getAnsweringSituationPublic = () => {
  return service({
    url: '/As/getAnsweringSituationPublic',
    method: 'get',
  })
}


// @Tags AnsweringSituation
// @Summary 用id查询答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AnsweringSituation true "用id查询答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /As/findAnsweringSituation [get]
export const findExamAnsweringSituation = (params) => {
  return service({
    url: '/As/findExamAnsweringSituation',
    method: 'get',
    params
  })
}


// @Tags AnsweringSituation
// @Summary 创建答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnsweringSituation true "创建答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /As/createAnsweringSituation [post]
export const createUserAnsweringSituation = (data) => {
  return service({
    url: '/As/createUserAnsweringSituation',
    method: 'post',
    data
  })
}



// @Tags AnsweringSituation
// @Summary 创建答题情况
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnsweringSituation true "创建答题情况"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /As/createAnsweringSituation [post]
export const submitUserAnsweringSituation = (data) => {
  return service({
    url: '/As/submitUserAnsweringSituation',
    method: 'post',
    data
  })
}
