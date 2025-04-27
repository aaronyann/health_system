<template>
  <div class="notification-list">
    <el-card>
      <h2>通知中心</h2>
      <el-table :data="notifications" style="width: 100%">
        <el-table-column prop="id" label="ID" width="60"></el-table-column>
        <el-table-column prop="message" label="消息内容"></el-table-column>
        <el-table-column prop="readStatus" label="是否已读" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.readStatus ? 'success' : 'warning'">
              {{ scope.row.readStatus ? '已读' : '未读' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="发送时间" width="180"></el-table-column>
      </el-table>
      <div class="pagination" style="margin-top: 20px; text-align: right">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="notifications.length"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'NotificationList',
  data() {
    return {
      notifications: [
        {
          id: 1,
          message: '系统维护通知：请于今晚10点后更新您的软件',
          readStatus: false,
          createdAt: '2025-04-01 18:30:00',
        },
        { id: 2, message: '新健康档案已上链', readStatus: true, createdAt: '2025-04-02 09:15:00' },
        { id: 3, message: '预约确认成功', readStatus: false, createdAt: '2025-04-03 14:20:00' },
      ],
      currentPage: 1,
      pageSize: 10,
    }
  },
  methods: {
    handlePageChange(page) {
      this.currentPage = page
      // 实际开发中请根据页码调用 API 获取数据
    },
  },
}
</script>

<style scoped>
.notification-list {
  padding: 20px;
}
</style>
