<template>
  <div class="admin-dashboard">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card>
          <div class="dashboard-item">
            <div class="dashboard-title">用户总数</div>
            <div class="dashboard-value">{{ totalUsers }}</div>
            <el-button type="primary" size="small" @click="goToUserManagement">管理用户</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div class="dashboard-item">
            <div class="dashboard-title">健康档案</div>
            <div class="dashboard-value">{{ totalHealthRecords }}</div>
            <el-button type="primary" size="small" @click="goToHealthRecords">查看档案</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div class="dashboard-item">
            <div class="dashboard-title">交易日志</div>
            <div class="dashboard-value">{{ totalTransactions }}</div>
            <el-button type="primary" size="small" @click="goToTransactionLogs">查看日志</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="8">
        <el-card>
          <div class="dashboard-item">
            <div class="dashboard-title">待审核记录</div>
            <div class="dashboard-value">{{ pendingAudits }}</div>
            <el-button type="primary" size="small" @click="goToAuditLogs">审核记录</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div class="dashboard-item">
            <div class="dashboard-title">医院信息</div>
            <div class="dashboard-value">{{ totalHospitals }}</div>
            <el-button type="primary" size="small" @click="goToHospitalManagement"
              >管理医院</el-button
            >
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div class="dashboard-item">
            <div class="dashboard-title">通知中心</div>
            <div class="dashboard-value">{{ notificationsCount }}</div>
            <el-button type="primary" size="small" @click="goToNotifications">查看通知</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'AdminDashboard',
  data() {
    return {
      totalUsers: 150,
      totalHealthRecords: 1200,
      totalTransactions: 350,
      pendingAudits: 5,
      totalHospitals: 12,
      notificationsCount: 8,
    }
  },
  methods: {
    goToUserManagement() {
      this.$router.push({ name: 'UserManagement' })
    },
    goToHealthRecords() {
      this.$router.push({ name: 'HealthRecords' })
    },
    goToTransactionLogs() {
      this.$router.push({ name: 'TransactionLogs' })
    },
    goToAuditLogs() {
      this.$router.push({ name: 'AuditLogs' })
    },
    goToHospitalManagement() {
      this.$router.push({ name: 'HospitalManagement' })
    },
    goToNotifications() {
      this.$router.push({ name: 'NotificationList' })
    },
  },
}
</script>

<style scoped>
.admin-dashboard {
  padding: 20px;
}
.dashboard-item {
  text-align: center;
}
.dashboard-title {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: 10px;
}
.dashboard-value {
  font-size: 2rem;
  color: #409eff;
  margin-bottom: 10px;
}
</style>
