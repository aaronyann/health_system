<template>
  <div class="health-records">
    <el-card>
      <h2>健康档案列表</h2>
      <el-table :data="records" style="width: 100%">
        <el-table-column prop="id" label="ID" width="60"></el-table-column>
        <el-table-column prop="patientName" label="患者姓名"></el-table-column>
        <el-table-column prop="dataHash" label="数据哈希"></el-table-column>
        <el-table-column prop="timestamp" label="记录时间"></el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="primary" size="small" @click="viewDetail(scope.row.id)"
              >详情</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <!-- 可根据需要添加分页 -->
      <div class="pagination" style="margin-top: 20px; text-align: right">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="records.length"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'HealthRecords',
  data() {
    return {
      records: [
        { id: 1, patientName: 'Alice', dataHash: 'QmHash1', timestamp: '2025-04-01 10:00:00' },
        { id: 2, patientName: 'Bob', dataHash: 'QmHash2', timestamp: '2025-04-02 11:00:00' },
        { id: 3, patientName: 'Charlie', dataHash: 'QmHash3', timestamp: '2025-04-03 12:00:00' },
      ],
      currentPage: 1,
      pageSize: 10,
    }
  },
  methods: {
    viewDetail(recordId) {
      this.$router.push({ name: 'HealthRecordDetail', params: { id: recordId } })
    },
    handlePageChange(page) {
      this.currentPage = page
      // 根据需要调用 API 获取对应页数据
    },
  },
}
</script>

<style scoped>
.health-records {
  padding: 20px;
}
</style>
