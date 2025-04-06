<script setup>
import { ref } from 'vue';
import { formatDate } from '@/utils/format'
import { getExamManagementList, getUserExamManagementList } from "@/api/ExaMan/sysEm";
const type = ref([
  '全部考试',
  '未开始',
  '报名中',
  '正在进行',
  '已结束'
])
const tagType = ref('全部考试')
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
    if (searchInfo.value.registrationRequired === ""){
      searchInfo.value.registrationRequired=null
    }
    getTableData()
  })
}


const getTableData = async() => {
  const table = await getUserExamManagementList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
</script>

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

        <el-form-item label="考试名称" prop="examName">
          <el-input v-model="searchInfo.examName" placeholder="搜索条件" />

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

      <el-segmented v-model="tagType" :options="type" class="mr-5" />
    </div>
    <div class="gva-table-box" style="margin-left: 200px;margin-right: 200px">
<!--      <div style="display: flex;flex-direction: row;justify-content: center;align-items: center">-->
<!--        <el-tag type="primary" effect="light" size="large">考试中心</el-tag>-->

<!--      </div>-->
      <el-row :gutter="20">
        <el-col :span="12" v-for="(item,index) in tableData" :key="index">
          <el-card style="max-width: 100%;margin-top: 10px;margin-bottom: 10px" shadow="hover" >
            <div style="color: #1890ff;font-weight: 700;font-size: 20px">
              {{ item.examName }}
            </div>
            <div>
              <span style="color: #555;font-weight: 700;">考试时间</span>
              <span style="margin-left: 20px"></span>
              <span style="color: #F56C6C;font-size: 14px;line-height: 26px;">{{ formatDate(item.examStartTime) }} ~ {{ formatDate(item.examEndTime) }}</span>
            </div>
            <div>
              <span style="color: #555;font-weight: 700;">考试状态</span>
              <span style="margin-left: 20px"></span>
              <span style="color: #67C23A">进行中</span>
            </div>
            <div>
              <span style="color: #555;font-weight: 700;">需要报名</span>
              <span style="margin-left: 20px"></span>
              <span style="color: #E6A23C;font-size: 14px;line-height: 26px;">否</span>
            </div>
            <div>
              <span style="color: #555;font-weight: 700;">考试总分</span>
              <span style="margin-left: 20px"></span>
              <span style="color: #888;font-size: 14px;line-height: 26px;">{{ item.scope }}分</span>
            </div>
            <div style="display: flex;flex-direction: row;float: right;margin-bottom: 20px">
              <el-button size="default">进入考试</el-button>
            </div>
          </el-card>
        </el-col>
      </el-row>
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
  </div>
</template>

<style scoped lang="scss">

</style>
