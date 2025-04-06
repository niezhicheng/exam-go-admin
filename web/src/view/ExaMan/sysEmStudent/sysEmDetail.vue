<script setup>
import {ref, watch, onMounted } from 'vue';
import {findPaperExam} from "@/api/Parer/sysParer";
import {
  createUserAnsweringSituation,
  findExamAnsweringSituation,
  submitUserAnsweringSituation
} from "@/api/AnsweSituation/sysAs"
import {ElMessage} from "element-plus";

const checkedItems = ref([]); // 存储选中的项


// 这个是多选的话
const onChange = (index) => {
  let ischeck = list_arab.value[index]
  if(currentData.value.testModel.topicType !==2){
    checkedItems.value = [];
    checkedItems.value.push(ischeck)
    return
  }
  if (checkedItems.value.includes(ischeck)) {
    // 如果已经选中，则移除
    checkedItems.value = checkedItems.value.filter(i => i !== ischeck);
  } else {
    // 否则添加
    checkedItems.value.push(ischeck);
  }
};

const data = ref([])
// 按 topicType 分组每种题型
const singleChoiceQuestions = ref([]);
const multipleChoiceQuestions = ref([]);
const fillInTheBlankQuestions = ref([]);
const trueFalseQuestions = ref([]);


const alreadyQuestion = ref([])

const underway = ref('')

const FindPaperExam = async () => {
  const res = await findPaperExam({"ID": 1})
  data.value = res.data
  const resp = await findExamAnsweringSituation({"ID": 1})
  alreadyQuestion.value = resp.data
  data.value.forEach(question => {
    switch (question.testModel.topicType) {
      case 1:
        singleChoiceQuestions.value.push(question);
        break;
      case 2:
        multipleChoiceQuestions.value.push(question);
        break;
      case 3:
        trueFalseQuestions.value.push(question);
        break;
      case 4:
        fillInTheBlankQuestions.value.push(question);
        break;
    }
  });
  if(alreadyQuestion.value.length === data.value.length){
    isall.value = true
    return
  }
  let istue = false; // 确保在循环外部定义此变量

// 检查单选题
  singleChoiceQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });

  if (istue) {
    return; // 如果找到，直接返回
  }

// 检查多选题
  multipleChoiceQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });

  if (istue) {
    return; // 如果找到，直接返回
  }
// 检查判断题
  trueFalseQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });
  if (istue) {
    return; // 如果找到，直接返回
  }
// 检查填空题
  fillInTheBlankQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });




}

onMounted(() => {
  FindPaperExam()
})


// 根据id获取对应状态颜色（也就是el-tag的type）的函数
const getStatusColor = (id) => {
  if(underway.value === id){
    return "primary"
  }
  const index = alreadyQuestion.value.findIndex((qId) => qId === id);
  if (index > -1) {
    return "danger";
  }
  // 同样可以根据更多情况返回不同type值来控制颜色，比如正在作答返回"primary"等
  return "info";
};

// 监视 underway.value 的变化
watch(underway, (newValue) => {
  findCorrespondingData(newValue);
});


const currentData = ref()

// 查找对应的值
const findCorrespondingData = (value) => {
  if (value) {
    currentData.value = data.value.find(item => item.testId === value)
    if(currentData.value.testModel.topicType === 1 || currentData.value.testModel.topicType === 2 ){
      parsedOptions.value = JSON.parse(currentData.value.testModel.option)
    }
    if(currentData.value.testModel.topicType === 3){
      parsedOptions.value = [{ value: '正确'},{ value: '错误'}]
    }
    if(currentData.value.testModel.topicType === 4){
      parsedOptions.value = ''
    }

  }
};

const parsedOptions = ref('');

const list_arab = ref(Array.from({ length: 26 }, (_, i) => String.fromCharCode(65 + i)))

// 检查项是否被选中的函数
const isChecked = (item,index) => {
  let ischeck = list_arab.value[index]
  return checkedItems.value.includes(ischeck);
};

const nextId = async () => {
  input.value = ""
  let checkedItemsString = checkedItems.value.join(',');
  const index = data.value.findIndex(jsonData => jsonData.testId === underway.value)
  if (index !== -1) {
    if(data.value[index]['testModel']['topicType'] === 3){
      if(checkedItemsString === 'A'){
        checkedItemsString = "正确"
      }
      if(checkedItemsString === 'B'){
        checkedItemsString = "错误"
      }
    }
  }
  console.log("题目id",underway.value)
  const res = await createUserAnsweringSituation({"examId": 1, "titleId": data.value[index]['testId'],"answer": checkedItemsString })
  if(res.code === 0){
    alreadyQuestion.value.push(underway.value)
    checkedItems.value = []

    updateNextId()
  }

}


// 上一步
const Previousstep = () => {
}


const updateNextId = () => {
  let istue = false; // 确保在循环外部定义此变量

// 检查单选题
  singleChoiceQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });

  if (istue) {
    return; // 如果找到，直接返回
  }

// 检查多选题
  multipleChoiceQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });

  if (istue) {
    return; // 如果找到，直接返回
  }

  // 检查判断题
  trueFalseQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });
  if (istue) {
    return; // 如果找到，直接返回
  }
// 检查填空题
  fillInTheBlankQuestions.value.some(item => {
    if (!alreadyQuestion.value.includes(item.testId)) {
      underway.value = item.testId; // 赋值
      istue = true; // 标记为找到
      return true; // 退出循环
    }
    return false; // 继续遍历
  });


  if(alreadyQuestion.value.length === data.value.length){
    isall.value = true
  }
}

const input = ref('')
const isall = ref(false)

const changeinput = (value) => {
  checkedItems.value[0] = value
}

const goToQuestion = (testId) => {
  checkedItems.value = []
  input.value = ''
  underway.value = testId; // 更新 underway
  findCorrespondingData(testId); // 更新当前题目数据
};


const submitpaper = async () => {
  const res = await submitUserAnsweringSituation({"examId": 1})
  if(res.code === 0){
    ElMessage({
      message: '交卷成功',
      type: 'success',
    })
  }
}

</script>

<template>
  <div>
    <div class="gva-search-box" style="display: flex;flex-direction: row;justify-content: space-between;align-items: center">
      <div style="font-size: 20px;font-weight: 400">
        <span>距离考试结束还有:</span>
        <span style="margin-left: 20px"></span>
        <span style="color: #F56C6C">44分钟48秒</span>
      </div>
      <div>
        <el-button type="primary" icon="arrow-down" @click="submitpaper()">交卷</el-button>
      </div>
    </div>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="gva-table-box">
          <el-alert title="答题卡" type="info" :closable="false" center/>
          <div style="display: flex;flex-direction: row;margin-top: 20px;margin-bottom: 20px">
            <el-tag type="info" size="large">未答题</el-tag>
            <el-tag type="primary" size="large" style="margin-left: 20px">作答中</el-tag>
            <el-tag type="danger" size="large" style="margin-left: 20px">已作答</el-tag>
          </div>
          <el-alert title="单选题" type="info" :closable="false" center/>
          <div style="margin-top: 20px; margin-bottom: 20px">
            <el-row :gutter="10">
              <el-col :span="2" v-for="(item,index) in singleChoiceQuestions" :key="index" style="margin: 10px;">
                <el-tag :type="getStatusColor(item.testId)" size="large" style="cursor: pointer;" @click="goToQuestion(item.testId)">{{ index + 1 }}</el-tag>
              </el-col>
            </el-row>
          </div>
          <el-alert title="多选题" type="info" :closable="false" center/>
          <div style="margin-top: 20px; margin-bottom: 20px">
            <el-row :gutter="10">
              <el-col :span="2" v-for="(item,index)  in multipleChoiceQuestions" :key="index" style="margin: 10px;">
                <el-tag :type="getStatusColor(item.testId)" size="large" style="cursor: pointer;" @click="goToQuestion(item.testId)">{{ singleChoiceQuestions.length + index + 1}}</el-tag>
              </el-col>
            </el-row>
          </div>
          <el-alert title="判断题" type="info" :closable="false" center/>
          <div style="margin-top: 20px; margin-bottom: 20px">
            <el-row :gutter="10">
              <el-col :span="2" v-for="(item,index) in trueFalseQuestions" :key="index" style="margin: 10px;">
                <el-tag :type="getStatusColor(item.testId)" size="large" style="cursor: pointer;" @click="goToQuestion(item.testId)">{{  singleChoiceQuestions.length + multipleChoiceQuestions.length + fillInTheBlankQuestions.length + index + 1 }}</el-tag>
              </el-col>
            </el-row>
          </div>
          <el-alert title="填空题" type="info" :closable="false" center/>
          <div style="margin-top: 20px; margin-bottom: 20px">
            <el-row :gutter="10">
              <el-col :span="2" v-for="(item,index)  in fillInTheBlankQuestions" :key="index" style="margin: 10px;">
                <el-tag :type="getStatusColor(item.testId)" size="large" style="cursor: pointer;" @click="goToQuestion(item.testId)">{{ singleChoiceQuestions.length + multipleChoiceQuestions.length + index + 1 }}</el-tag>
              </el-col>
            </el-row>
          </div>

        </div>
      </el-col>
      <el-col :span="18">
        <div class="gva-table-box">
          <div style="padding: 10px;font-weight: bold;font-size: 20px">
            {{ currentData ? currentData.testModel.subject : '未选择任何题目' }}
          </div>
          <div v-for="(item,index) in parsedOptions" :key="index" v-if="parsedOptions !== ''">
            <el-check-tag
                :checked="isChecked(item,index)"
                type="primary"
                :hit="true"
                @change="() => onChange(index)"
                effect="plain"
                style="width: 500px; margin: 10px; height: 50px; text-align: left; display: flex; align-items: center;"
            >
              {{ list_arab[index] }}. {{ item.value }}
            </el-check-tag>
          </div>
          <div style="padding: 10px" v-if="parsedOptions === '' && isall === false">
            <el-input v-model="input" style="width: 240px" @change="changeinput" placeholder="请输入答案" />
          </div>
          <div style="padding: 10px" v-if="isall === true">
            已经答题完全部 请提交交卷
          </div>
          <div style="padding: 10px">
<!--            <el-button type="success" @click="Previousstep()" v-if="alreadyQuestion.length >= 1">上一题</el-button>-->
            <el-button type="primary" @click="nextId()" v-if="isall === false">下一题</el-button>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">

</style>
