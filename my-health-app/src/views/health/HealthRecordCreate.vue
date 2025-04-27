<template>
  <div class="health-record-create">
    <el-card>
      <h2>新建健康档案</h2>
      <el-form :model="recordForm" ref="recordFormRef" label-width="120px">
        <el-form-item label="患者姓名">
          <el-input v-model="recordForm.patientName" placeholder="请输入患者姓名"></el-input>
        </el-form-item>
        <el-form-item label="数据哈希">
          <el-input v-model="recordForm.dataHash" placeholder="请输入数据哈希或上传文件"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">提交档案</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'HealthRecordCreate',
  data() {
    return {
      recordForm: {
        patientName: '',
        dataHash: '',
      },
    }
  },
  methods: {
    handleSubmit() {
      // 模拟提交健康档案逻辑，实际项目中请替换为 API 调用
      console.log('提交的健康档案：', this.recordForm)
      this.$message.success('健康档案已提交')
      this.$router.push({ name: 'HealthRecords' })
    },
    handleReset() {
      this.recordForm = {
        patientName: '',
        dataHash: '',
      }
    },
  },
}
</script>

<style scoped>
.health-record-create {
  padding: 20px;
}
</style>
