<template>
  <div class="dashboard-container">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card class="dashboard-card">
          <h3>健康档案</h3>
          <p>
            您共有 <strong>{{ healthRecordsCount }}</strong> 条健康档案记录
          </p>
          <el-button type="primary" @click="goToHealthRecords">查看详情</el-button>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="dashboard-card">
          <h3>就诊预约</h3>
          <p>
            您有 <strong>{{ appointmentsCount }}</strong> 个预约
          </p>
          <el-button type="primary" @click="goToAppointments">查看预约</el-button>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card class="dashboard-card">
          <h3>通知</h3>
          <p>
            您有 <strong>{{ notificationsCount }}</strong> 条新通知
          </p>
          <el-button type="primary" @click="goToNotifications">查看通知</el-button>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'PatientDashboard',
  data() {
    return {
      healthRecordsCount: 0,
      appointmentsCount: 0,
      notificationsCount: 0,
    }
  },
  methods: {
    goToHealthRecords() {
      this.$router.push({ name: 'HealthRecords' })
    },
    goToAppointments() {
      this.$router.push({ name: 'AppointmentList' })
    },
    goToNotifications() {
      this.$router.push({ name: 'NotificationList' })
    },
  },
  mounted() {
    // 模拟数据获取，实际项目中应调用 API 获取真实数据
    this.healthRecordsCount = 5
    this.appointmentsCount = 2
    this.notificationsCount = 3
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
