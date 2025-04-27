<template>
  <el-aside width="200px" class="app-sidebar">
    <el-menu :default-active="activeMenu" class="sidebar-menu" unique-opened>
      <el-menu-item index="dashboard">
        <!-- 指向具体的仪表盘路径，此处默认使用患者仪表盘 -->
        <router-link to="/dashboard/patient">仪表盘</router-link>
      </el-menu-item>
      <el-menu-item index="health-records">
        <router-link to="/health/records">健康档案</router-link>
      </el-menu-item>
      <el-menu-item index="appointments">
        <router-link to="/appointments">就诊预约</router-link>
      </el-menu-item>
      <el-menu-item index="prescriptions">
        <router-link to="/prescriptions">处方</router-link>
      </el-menu-item>
      <el-menu-item index="lab-reports">
        <router-link to="/lab/reports">检验报告</router-link>
      </el-menu-item>
      <el-menu-item index="insurance">
        <router-link to="/insurance">保险信息</router-link>
      </el-menu-item>
      <el-menu-item index="notifications">
        <router-link to="/notifications">通知中心</router-link>
      </el-menu-item>
    </el-menu>
  </el-aside>
</template>

<script>
export default {
  name: 'SidebarPage',
  computed: {
    activeMenu() {
      const path = this.$route.path
      if (path.includes('/dashboard')) {
        return 'dashboard'
      } else if (path.includes('/health')) {
        return 'health-records'
      } else if (path.includes('/appointments')) {
        return 'appointments'
      } else if (path.includes('/prescriptions')) {
        return 'prescriptions'
      } else if (path.includes('/lab')) {
        return 'lab-reports'
      } else if (path.includes('/insurance')) {
        return 'insurance'
      } else if (path.includes('/notifications')) {
        return 'notifications'
      } else {
        return ''
      }
    },
  },
}
</script>

<style scoped>
.app-sidebar {
  background-color: #f2f2f2;
  border-right: 1px solid #dcdfe6;
}
.sidebar-menu {
  border-right: none;
}
.sidebar-menu .el-menu-item a {
  color: #333;
  text-decoration: none;
}
</style>
