<template>
  <div class="appointment-detail">
    <el-card>
      <h2>预约详情</h2>
      <el-descriptions title="预约信息" :column="1">
        <el-descriptions-item label="预约ID">{{ appointment.id }}</el-descriptions-item>
        <el-descriptions-item label="患者姓名">{{ appointment.patientName }}</el-descriptions-item>
        <el-descriptions-item label="医生姓名">{{ appointment.doctorName }}</el-descriptions-item>
        <el-descriptions-item label="科室">{{ appointment.department }}</el-descriptions-item>
        <el-descriptions-item label="预约时间">{{
          appointment.appointmentTime
        }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ appointment.status }}</el-descriptions-item>
        <el-descriptions-item label="备注">{{ appointment.notes }}</el-descriptions-item>
      </el-descriptions>
      <div style="margin-top: 20px">
        <el-button type="primary" @click="goBack">返回列表</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'AppointmentDetail',
  data() {
    return {
      appointment: {},
    }
  },
  methods: {
    fetchAppointment(id) {
      // 模拟获取预约详情数据，实际项目中请替换为 API 调用
      const appointments = [
        {
          id: 1,
          patientName: 'Alice',
          doctorName: 'Dr. Smith',
          department: '内科',
          appointmentTime: '2025-04-10 09:00:00',
          status: '已确认',
          notes: '按时就诊',
        },
        {
          id: 2,
          patientName: 'Bob',
          doctorName: 'Dr. Johnson',
          department: '外科',
          appointmentTime: '2025-04-11 10:30:00',
          status: '待确认',
          notes: '等待医生确认',
        },
        {
          id: 3,
          patientName: 'Charlie',
          doctorName: 'Dr. Lee',
          department: '儿科',
          appointmentTime: '2025-04-12 14:00:00',
          status: '已取消',
          notes: '患者取消预约',
        },
      ]
      this.appointment = appointments.find((a) => a.id == id) || {}
    },
    goBack() {
      this.$router.push({ name: 'AppointmentList' })
    },
  },
  mounted() {
    const appointmentId = this.$route.params.id
    this.fetchAppointment(appointmentId)
  },
}
</script>

<style scoped>
.appointment-detail {
  padding: 20px;
}
</style>
