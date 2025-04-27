<template>
  <div class="appointment-list">
    <el-card>
      <h2>就诊预约列表</h2>
      <el-table :data="appointments" style="width: 100%">
        <el-table-column prop="id" label="预约ID" width="80"></el-table-column>
        <el-table-column prop="patientName" label="患者姓名"></el-table-column>
        <el-table-column prop="doctorName" label="医生姓名"></el-table-column>
        <el-table-column prop="department" label="科室"></el-table-column>
        <el-table-column prop="appointmentTime" label="预约时间"></el-table-column>
        <el-table-column prop="status" label="状态"></el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button type="primary" size="small" @click="viewDetail(scope.row.id)">
              详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination" style="margin-top: 20px; text-align: right">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="appointments.length"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'AppointmentList',
  data() {
    return {
      appointments: [
        {
          id: 1,
          patientName: 'Alice',
          doctorName: 'Dr. Smith',
          department: '内科',
          appointmentTime: '2025-04-10 09:00:00',
          status: '已确认',
        },
        {
          id: 2,
          patientName: 'Bob',
          doctorName: 'Dr. Johnson',
          department: '外科',
          appointmentTime: '2025-04-11 10:30:00',
          status: '待确认',
        },
        {
          id: 3,
          patientName: 'Charlie',
          doctorName: 'Dr. Lee',
          department: '儿科',
          appointmentTime: '2025-04-12 14:00:00',
          status: '已取消',
        },
      ],
      currentPage: 1,
      pageSize: 10,
    }
  },
  methods: {
    viewDetail(appointmentId) {
      this.$router.push({ name: 'AppointmentDetail', params: { id: appointmentId } })
    },
    handlePageChange(page) {
      this.currentPage = page
      // 根据需要调用 API 获取对应页数据
    },
  },
}
</script>

<style scoped>
.appointment-list {
  padding: 20px;
}
</style>
