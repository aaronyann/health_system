<template>
  <el-container style="min-height: 100vh">
    <!-- 顶部导航栏 -->
    <Header />
    <el-container>
      <!-- 左侧侧边栏 -->
      <Sidebar />
      <!-- 主内容区域 -->
      <el-main>
        <router-view />
      </el-main>
    </el-container>
    <!-- 底部信息 -->
    <Footer />
  </el-container>
</template>

<script>
import Header from '@/components/common/Header.vue'
import Sidebar from '@/components/common/Sidebar.vue'
import Footer from '@/components/common/Footer.vue'

export default {
  name: 'MainLayout',
  components: {
    // eslint-disable-next-line vue/no-reserved-component-names
    Header,
    Sidebar,
    // eslint-disable-next-line vue/no-reserved-component-names
    Footer,
  },
}
</script>

<style scoped>
/* 可根据需要在此添加布局相关样式 */
</style>
