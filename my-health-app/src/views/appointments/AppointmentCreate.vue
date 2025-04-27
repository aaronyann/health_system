<template>
  <div class="appointment-create">
    <el-card>
      <h2>新建预约</h2>
      <el-form :model="appointmentForm" ref="appointmentFormRef" label-width="120px">
        <el-form-item label="患者姓名">
          <el-input v-model="appointmentForm.patientName" placeholder="请输入患者姓名"></el-input>
        </el-form-item>
        <el-form-item label="医生姓名">
          <el-input v-model="appointmentForm.doctorName" placeholder="请输入医生姓名"></el-input>
        </el-form-item>
        <el-form-item label="科室">
          <el-input v-model="appointmentForm.department" placeholder="请输入科室"></el-input>
        </el-form-item>
        <el-form-item label="预约时间">
          <el-date-picker
            v-model="appointmentForm.appointmentTime"
            type="datetime"
            placeholder="选择预约时间"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="appointmentForm.notes"
            placeholder="请输入备注"
            type="textarea"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSubmit">提交预约</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'AppointmentCreate',
  data() {
    return {
      appointmentForm: {
        patientName: '',
        doctorName: '',
        department: '',
        appointmentTime: '',
        notes: '',
      },
    }
  },
  methods: {
    handleSubmit() {
      // 模拟提交预约数据，实际项目中应调用 API
      console.log('提交的预约信息：', this.appointmentForm)
      this.$message.success('预约提交成功')
      this.$router.push({ name: 'AppointmentList' })
    },
    handleReset() {
      this.appointmentForm = {
        patientName: '',
        doctorName: '',
        department: '',
        appointmentTime: '',
        notes: '',
      }
    },
  },
}
</script>

<style scoped>
.appointment-create {
  padding: 20px;
}
</style>
