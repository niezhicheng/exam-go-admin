
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="题目:" prop="subject">
          <el-input v-model="formData.subject" :clearable="true"  placeholder="请输入题目" />
       </el-form-item>
        <el-form-item label="内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="题目类型:" prop="topicType">
          <el-input v-model.number="formData.topicType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="答案:" prop="answer">
          <el-input v-model="formData.answer" :clearable="true"  placeholder="请输入答案" />
       </el-form-item>
        <el-form-item label="解析:" prop="analysis">
          <RichEdit v-model="formData.analysis"/>
       </el-form-item>
        <el-form-item label="禁用:" prop="disable">
          <el-switch v-model="formData.disable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createTopic,
  updateTopic,
  findTopic
} from '@/api/Topic/sysTopic'

defineOptions({
    name: 'TopicForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            subject: '',
            content: '',
            topicType: undefined,
            answer: '',
            analysis: '',
            disable: false,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTopic({ ID: route.query.id })
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
               res = await createTopic(formData.value)
               break
             case 'update':
               res = await updateTopic(formData.value)
               break
             default:
               res = await createTopic(formData.value)
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
