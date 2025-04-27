<template>
  <el-container>
    <el-header>
      <el-menu mode="horizontal" default-active="1">
        <el-menu-item index="1">
          <router-link to="/">首页</router-link>
        </el-menu-item>
        <el-menu-item index="2">
          <router-link to="/login">登录</router-link>
        </el-menu-item>
        <el-menu-item index="3">
          <router-link to="/register">注册</router-link>
        </el-menu-item>
      </el-menu>
    </el-header>
    <el-main class="home-main">
      <h1>欢迎来到健康档案系统</h1>
      <p>基于区块链的健康档案管理示例项目</p>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'IndexHome',
}
</script>

<style scoped>
.home-main {
  text-align: center;
  margin-top: 50px;
}
</style>
