
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="考试名称:" prop="examName">
          <el-input v-model="formData.examName" :clearable="true"  placeholder="请输入考试名称" />
       </el-form-item>
        <el-form-item label="考试封面图:" prop="examCoverPicture">
          <SelectImage v-model="formData.examCoverPicture" file-type="image"/>
       </el-form-item>
        <el-form-item label="考试描述:" prop="examDescription">
          <el-input v-model="formData.examDescription" :clearable="true"  placeholder="请输入考试描述" />
       </el-form-item>
        <el-form-item label="是否需要报名:" prop="registrationRequired">
          <el-switch v-model="formData.registrationRequired" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="报名开始时间:" prop="registrationStartTime">
          <el-date-picker v-model="formData.registrationStartTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="报名结束时间:" prop="registrationDeadline">
          <el-date-picker v-model="formData.registrationDeadline" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="考试开始时间:" prop="examStartTime">
          <el-date-picker v-model="formData.examStartTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="考试结束时间:" prop="examEndTime">
          <el-date-picker v-model="formData.examEndTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createExamManagement,
  updateExamManagement,
  findExamManagement
} from '@/api/ExaMan/sysEm'

defineOptions({
    name: 'ExamManagementForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            examName: '',
            examCoverPicture: "",
            examDescription: '',
            registrationRequired: false,
            registrationStartTime: new Date(),
            registrationDeadline: new Date(),
            examStartTime: new Date(),
            examEndTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findExamManagement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createExamManagement(formData.value)
               break
             case 'update':
               res = await updateExamManagement(formData.value)
               break
             default:
               res = await createExamManagement(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
