<template>
  <div class="prescription-list">
    <el-card>
      <h2>处方列表</h2>
      <el-table :data="prescriptions" style="width: 100%">
        <el-table-column prop="id" label="处方ID" width="80"></el-table-column>
        <el-table-column prop="patientName" label="患者姓名"></el-table-column>
        <el-table-column prop="doctorName" label="医生姓名"></el-table-column>
        <el-table-column prop="date" label="日期"></el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="primary" size="small" @click="viewDetail(scope.row.id)"
              >详情</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination" style="margin-top: 20px; text-align: right">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="prescriptions.length"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'PrescriptionList',
  data() {
    return {
      prescriptions: [
        {
          id: 1,
          patientName: 'Alice',
          doctorName: 'Dr. Smith',
          date: '2025-04-10',
        },
        {
          id: 2,
          patientName: 'Bob',
          doctorName: 'Dr. Johnson',
          date: '2025-04-11',
        },
        {
          id: 3,
          patientName: 'Charlie',
          doctorName: 'Dr. Lee',
          date: '2025-04-12',
        },
      ],
      currentPage: 1,
      pageSize: 10,
    }
  },
  methods: {
    viewDetail(id) {
      this.$router.push({ name: 'PrescriptionDetail', params: { id } })
    },
    handlePageChange(page) {
      this.currentPage = page
      // 根据需要调用 API 获取对应页数据
    },
  },
}
</script>

<style scoped>
.prescription-list {
  padding: 20px;
}
</style>
