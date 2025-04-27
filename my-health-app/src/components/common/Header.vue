<template>
  <el-header class="app-header">
    <div class="logo">
      <router-link to="/">HealthApp</router-link>
    </div>
    <el-menu mode="horizontal" :default-active="activeIndex" class="nav-menu">
      <el-menu-item index="home">
        <router-link to="/">首页</router-link>
      </el-menu-item>
      <el-menu-item index="profile">
        <router-link to="/profile">个人中心</router-link>
      </el-menu-item>
      <el-menu-item index="logout" @click="handleLogout"> 退出 </el-menu-item>
    </el-menu>
  </el-header>
</template>

<script>
export default {
  name: 'HeaderL',
  data() {
    return {
      activeIndex: '',
    }
  },
  methods: {
    handleLogout() {
      // 模拟退出登录操作（清除登录状态后跳转登录页面）
      localStorage.removeItem('isAuthenticated')
      this.$router.push({ name: 'Login' })
    },
  },
  watch: {
    $route(to) {
      // 根据当前路由设置菜单选中状态（示例简单判断）
      if (to.path === '/') {
        this.activeIndex = 'home'
      } else if (to.path.includes('/profile')) {
        this.activeIndex = 'profile'
      } else {
        this.activeIndex = ''
      }
    },
  },
  mounted() {
    // 初始化选中项
    if (this.$route.path === '/') {
      this.activeIndex = 'home'
    } else if (this.$route.path.includes('/profile')) {
      this.activeIndex = 'profile'
    }
  },
}
</script>

<style scoped>
.app-header {
  display: flex;
  align-items: center;
  background-color: #409eff;
  color: #fff;
  padding: 0 20px;
}
.logo {
  font-size: 1.5rem;
  font-weight: bold;
  margin-right: 30px;
}
.nav-menu {
  flex: 1;
}
.nav-menu .el-menu-item a {
  color: #fff;
  text-decoration: none;
}
</style>
