
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="考试id:" prop="examId">
          <el-input v-model.number="formData.examId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="题目id:" prop="titleId">
          <el-input v-model.number="formData.titleId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="答案:" prop="answer">
          <el-input v-model="formData.answer" :clearable="true"  placeholder="请输入答案" />
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
  createAnsweringSituation,
  updateAnsweringSituation,
  findAnsweringSituation
} from '@/api/AnsweSituation/sysAs'

defineOptions({
    name: 'AnsweringSituationForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            userId: undefined,
            examId: undefined,
            titleId: undefined,
            answer: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAnsweringSituation({ ID: route.query.id })
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
               res = await createAnsweringSituation(formData.value)
               break
             case 'update':
               res = await updateAnsweringSituation(formData.value)
               break
             default:
               res = await createAnsweringSituation(formData.value)
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
