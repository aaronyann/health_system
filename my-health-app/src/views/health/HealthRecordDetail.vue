<template>
  <div class="health-record-detail">
    <el-card>
      <h2>健康档案详情</h2>
      <el-descriptions title="档案信息" :column="1">
        <el-descriptions-item label="ID">{{ record.id }}</el-descriptions-item>
        <el-descriptions-item label="患者姓名">{{ record.patientName }}</el-descriptions-item>
        <el-descriptions-item label="数据哈希">{{ record.dataHash }}</el-descriptions-item>
        <el-descriptions-item label="记录时间">{{ record.timestamp }}</el-descriptions-item>
      </el-descriptions>
      <div style="margin-top: 20px;">
        <el-button @click="goBack">返回列表</el-button>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'HealthRecordDetail',
  data() {
    return {
      record: {}
    }
  },
  methods: {
    fetchRecord(id) {
      // 模拟获取记录详情，实际开发中请替换为 API 调用
      const records = [
        { id: 1, patientName: 'Alice', dataHash: 'QmHash1', timestamp: '2025-04-01 10:00:00' },
        { id: 2, patientName: 'Bob', dataHash: 'QmHash2', timestamp: '2025-04-02 11:00:00' },
        { id: 3, patientName: 'Charlie', dataHash: 'QmHash3', timestamp: '2025-04-03 12:00:00' },
      ];
      this.record = records.find(r => r.id == id) || {};
    },
    goBack() {
      this.$router.push({ name: 'HealthRecords' });
    }
  },
  mounted() {
    const recordId = this.$route.params.id;
    this.fetchRecord(recordId);
  }
}
</script>

<style scoped>
.health-record-detail {
  padding: 20px;
}
</style>
