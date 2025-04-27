<template>
  <div class="dashboard-container">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card class="dashboard-card">
          <h3>待处理预约</h3>
          <p>
            您有 <strong>{{ pendingAppointmentsCount }}</strong> 个待处理预约
          </p>
          <el-button type="primary" @click="goToAppointments">查看预约</el-button>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="dashboard-card">
          <h3>健康档案审核</h3>
          <p>
            有 <strong>{{ pendingRecordsCount }}</strong> 条健康档案待审核
          </p>
          <el-button type="primary" @click="goToHealthRecords">审核档案</el-button>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="dashboard-card">
          <h3>消息通知</h3>
          <p>
            您有 <strong>{{ notificationsCount }}</strong> 条新消息
          </p>
          <el-button type="primary" @click="goToNotifications">查看消息</el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'DoctorDashboard',
  data() {
    return {
      pendingAppointmentsCount: 0,
      pendingRecordsCount: 0,
      notificationsCount: 0,
    }
  },
  methods: {
    goToAppointments() {
      this.$router.push({ name: 'AppointmentList' })
    },
    goToHealthRecords() {
      this.$router.push({ name: 'HealthRecords' })
    },
    goToNotifications() {
      this.$router.push({ name: 'NotificationList' })
    },
  },
  mounted() {
    // 模拟数据获取，实际项目中请调用 API 获取真实数据
    this.pendingAppointmentsCount = 4
    this.pendingRecordsCount = 3
    this.notificationsCount = 5
  },
}
</script>

<style scoped>
.dashboard-container {
  padding: 20px;
}
.dashboard-card {
  text-align: center;
}
.dashboard-card h3 {
  margin-bottom: 10px;
}
</style>
