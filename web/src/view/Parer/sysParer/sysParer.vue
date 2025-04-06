
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>

        <el-form-item label="试卷名称" prop="paperName">
         <el-input v-model="searchInfo.paperName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="描述" prop="paperDescription">
         <el-input v-model="searchInfo.paperDescription" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>

        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

          <el-table-column align="left" label="试卷名称" prop="paperName" width="120" />
          <el-table-column align="left" label="描述" prop="paperDescription" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
              <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
              <el-button  type="primary" link icon="edit" class="table-button" @click="updatePaperFunc(scope.row)">变更</el-button>
              <el-button  type="primary" link icon="edit" class="table-button" @click="PaperDetail(scope.row)">编辑试题</el-button>
              <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="试卷名称:"  prop="paperName" >
              <el-input v-model="formData.paperName" :clearable="true"  placeholder="请输入试卷名称" />
            </el-form-item>
            <el-form-item label="描述:"  prop="paperDescription" >
              <el-input v-model="formData.paperDescription" :clearable="true"  placeholder="请输入描述" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="试卷名称">
                        {{ detailFrom.paperName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="描述">
                        {{ detailFrom.paperDescription }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>



    <el-dialog v-model="dialogVisible" title="试卷详情">
      <el-button @click="addRow">添加</el-button>
      <el-table :data="paperDetails" width="500">
        <el-table-column
            prop="testModel.subject"
            label="题目"
            width="300">
        </el-table-column>
        <el-table-column align="left" label="题目类型" prop="topicType" >
          <template #default="scope">
            {{ getTopicTypeLabel(scope.row.testModel.topicType)}}
          </template>
        </el-table-column>
        <el-table-column align="left" label="分数" prop="scope"  />
        <el-table-column label="操作" width="100">
          <template #default="scope">
            <el-button @click="deleteRowDetail(scope.row)" type="danger" size="mini">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <span class="dialog-footer">
        <div style="display: flex;flex-direction: row;justify-content: space-between">
          <div></div>
          <div style="margin-top: 20px">
            <el-button @click="dialogVisible = false">关闭</el-button>
          </div>

        </div>

      </span>
    </el-dialog>


    <!-- 添加试卷详情的弹窗 -->
    <el-dialog v-model="addRowDialogVisible" title="添加试卷详情">
      <el-table
          style="width: 100%"
          tooltip-effect="dark"
          :data="questionList"
      >
        <el-table-column align="left" label="题目" prop="subject" />
        <el-table-column align="left" label="题目类型" prop="topicType" >
          <template #default="scope">
            {{ getTopicTypeLabel(scope.row.topicType)}}
          </template>
        </el-table-column>
        <el-table-column align="left" label="分数" prop="scope"  />
        <el-table-column fixed="right" label="操作" min-width="120">
          <template #default="scope">
            <el-button  type="primary" link icon="document-add" @click="addRowTopic(scope.row)">添加</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="QuestionTotal"
            @current-change="handleQuestionCurrentChange"
            @size-change="handleQuestionSizeChange"
        />
      </div>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createPaper,
  deletePaper,
  deletePaperByIds,
  updatePaper,
  findPaper,
  getPaperList
} from '@/api/Parer/sysParer'
import {
  createPaperDetail,
  deletePaperDetail,
  FindPaperDetailAll,
  GetQuestionListTableData
} from '@/api/ParerDetail/sysPaperDetail'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'Paper'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
    paperName: '',
    paperDescription: '',
})

const addData = ref({
  paperId: 0,
  testId: 0,
  scope: 0
})


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getPaperList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deletePaperFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deletePaperByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePaperFunc = async(row) => {
    const res = await findPaper({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePaperFunc = async (row) => {
    const res = await deletePaper({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        paperName: '',
        paperDescription: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createPaper(formData.value)
                  break
                case 'update':
                  res = await updatePaper(formData.value)
                  break
                default:
                  res = await createPaper(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findPaper({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}



const dialogVisible = ref(false);
const paperDetails = ref([]);
const paperId = ref(0)
const PaperDetail = async (row) => {
  paperId.value = row.ID
  const res = await FindPaperDetailAll({"paperId":row.ID})
  console.log()
  if(res.code === 0){
    dialogVisible.value = true
    paperDetails.value = res.data.list
  }
}

// 从函数里面用的
const PaperDetailAll = async (id) => {
  const res = await FindPaperDetailAll({"paperId":id})
  if(res.code === 0){
    paperDetails.value = res.data.list
  }
}

const addRowTopic = async (row) => {
  const res = await createPaperDetail({paperId: paperId.value,testId: row.ID,scope:row.scope})
  if(res.code === 0) {
    PaperDetailAll(paperId.value)
    getQuestionListTableData()
  }
}

const deleteRowDetail = async (row) => {
  const res = await deletePaperDetail({"ID": row.ID})
  console.log(res,"这是res")
  if(res.code === 0) {
    PaperDetailAll(paperId.value)
    getQuestionListTableData()
  }
}

const addRowDialogVisible = ref(false);
const questionList = ref([])
const addRow = () => {
  addRowDialogVisible.value = true
  getQuestionListTableData()
}

const QuestionPage = ref(1)
const QuestionPageSize = ref(10)
const QuestionTotal = ref(0)
// 分页
const handleQuestionSizeChange = (val) => {
  QuestionPageSize.value = val
  getQuestionListTableData()
}

// 修改页面容量
const handleQuestionCurrentChange = (val) => {
  QuestionPage.value = val
  getQuestionListTableData()
}

// 查询
const getQuestionListTableData = async() => {
  const table = await GetQuestionListTableData({ page: page.value, pageSize: pageSize.value,paperId: paperId.value })
  if (table.code === 0) {
    questionList.value = table.data.list
    QuestionTotal.value = table.data.total
    QuestionPage.value = table.data.page
    QuestionPageSize.value = table.data.pageSize
  }
}


const options = ref([{
  value: 1,
  label: '单选题'
}, {
  value: 2,
  label: '多选题'
}, {
  value: 3,
  label: '判断题'
}, {
  value: 4,
  label: '填空题'
}])
// 根据 topicType 返回对应的标签
const getTopicTypeLabel = (type) => {
  const option = options.value.find(opt => opt.value === type);
  return option ? option.label : '未知类型';
};
</script>

<style>

</style>
