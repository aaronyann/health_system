<template>
  <div class="insurance-info">
    <el-card>
      <h2>保险信息</h2>
      <el-form :model="insuranceForm" ref="insuranceFormRef" label-width="120px">
        <el-form-item label="保险公司">
          <el-input v-model="insuranceForm.provider" placeholder="请输入保险公司名称"></el-input>
        </el-form-item>
        <el-form-item label="保单号码">
          <el-input v-model="insuranceForm.policyNumber" placeholder="请输入保单号码"></el-input>
        </el-form-item>
        <el-form-item label="生效日期">
          <el-date-picker v-model="insuranceForm.validFrom" type="date" placeholder="选择生效日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="失效日期">
          <el-date-picker v-model="insuranceForm.validTo" type="date" placeholder="选择失效日期">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="保险范围">
          <el-input
            v-model="insuranceForm.coverage"
            type="textarea"
            placeholder="请输入保险覆盖范围及说明"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">保存信息</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'InsuranceInfo',
  data() {
    return {
      insuranceForm: {
        provider: '',
        policyNumber: '',
        validFrom: '',
        validTo: '',
        coverage: '',
      },
    }
  },
  methods: {
    handleSubmit() {
      // 模拟提交保险信息，实际开发中请调用 API
      console.log('提交保险信息：', this.insuranceForm)
      this.$message.success('保险信息已保存')
    },
    handleReset() {
      this.insuranceForm = {
        provider: '',
        policyNumber: '',
        validFrom: '',
        validTo: '',
        coverage: '',
      }
    },
  },
}
</script>

<style scoped>
.insurance-info {
  padding: 20px;
}
</style>
