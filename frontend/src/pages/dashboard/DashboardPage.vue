<template>
  <div class="dashboard">
    <h2 class="page-title">📈 数据概览</h2>

    <div class="stats-grid">
      <StatsCard
        label="总用户数"
        :value="totalUsers"
        :highlight="true"
        :trend="12"
        trend-label="较上月"
      />
      <StatsCard
        label="活跃用户"
        :value="892"
        :trend="5"
        trend-label="较上月"
      />
      <StatsCard
        label="今日新增"
        :value="56"
        :trend="-3"
        trend-label="较昨日"
      />
    </div>

    <div class="quick-actions">
      <h3 class="section-title">⚡ 快捷操作</h3>
      <div class="actions-row">
        <el-button @click="$router.push('/users')">
          <el-icon><User /></el-icon> 查看用户列表
        </el-button>
        <el-button @click="$router.push('/profile')">
          <el-icon><UserFilled /></el-icon> 编辑个人资料
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import StatsCard from '@/components/StatsCard.vue'
import * as userApi from '@/api/user'

const totalUsers = ref<number>(0)

onMounted(async () => {
  try {
    const res = await userApi.listUsers({ page: 1, page_size: 1 })
    totalUsers.value = res.total
  } catch {
    totalUsers.value = 0
  }
})
</script>

<style scoped>
.dashboard {
  max-width: 960px;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a2e;
  margin-bottom: 20px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-bottom: 24px;
}

.section-title {
  font-size: 15px;
  font-weight: 600;
  color: #333;
  margin-bottom: 12px;
}

.quick-actions {
  background: white;
  border: 1px solid #e8e8e8;
  border-radius: 12px;
  padding: 20px 24px;
}

.actions-row {
  display: flex;
  gap: 12px;
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
