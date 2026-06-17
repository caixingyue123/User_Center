<template>
  <div class="user-list-page">
    <div class="page-card">
      <div class="page-card-header">
        <h2 class="page-title">👥 用户管理</h2>
        <el-input
          v-model="keyword"
          placeholder="搜索用户名..."
          :prefix-icon="Search"
          style="width: 240px"
          clearable
          @input="handleSearch"
        />
      </div>

      <el-table
        :data="filteredUsers"
        v-loading="loading"
        stripe
        style="width: 100%"
        empty-text="暂无用户数据"
      >
        <el-table-column prop="id" label="ID" width="80" align="center" />
        <el-table-column prop="username" label="用户名" min-width="120" />
        <el-table-column prop="nickname" label="昵称" min-width="120">
          <template #default="{ row }">
            {{ row.nickname || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="email" label="邮箱" min-width="180">
          <template #default="{ row }">
            {{ row.email || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="phone" label="电话" min-width="140">
          <template #default="{ row }">
            {{ row.phone || '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="注册时间" min-width="160">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" align="center" fixed="right">
          <template #default="{ row }">
            <el-popconfirm
              title="确定要删除该用户吗？此操作不可撤销。"
              confirm-button-text="确定删除"
              cancel-button-text="取消"
              @confirm="handleDelete(row.id)"
            >
              <template #reference>
                <el-button type="danger" size="small" text>删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-wrapper">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="total"
          :page-sizes="[10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          background
          @size-change="fetchUsers"
          @current-change="fetchUsers"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import * as userApi from '@/api/user'
import type { User } from '@/types/user'

const users = ref<User[]>([])
const total = ref(0)
const loading = ref(false)
const keyword = ref('')

const pagination = ref({
  page: 1,
  pageSize: 10,
})

const filteredUsers = computed(() => {
  if (!keyword.value.trim()) return users.value
  const kw = keyword.value.toLowerCase()
  return users.value.filter(
    (u) =>
      u.username.toLowerCase().includes(kw) ||
      (u.nickname && u.nickname.toLowerCase().includes(kw)),
  )
})

function formatDate(dateStr: string): string {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  const pad = (n: number) => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}`
}

async function fetchUsers() {
  loading.value = true
  try {
    const res = await userApi.listUsers({
      page: pagination.value.page,
      page_size: pagination.value.pageSize,
    })
    users.value = res.list || []
    total.value = res.total
  } catch {
    // Error handled by interceptor
  } finally {
    loading.value = false
  }
}

function handleSearch() {
  // Client-side filtering; no API call needed
}

async function handleDelete(id: number) {
  try {
    await userApi.deleteUser(id)
    ElMessage.success('删除成功')
    fetchUsers()
  } catch {
    // Error handled by interceptor
  }
}

onMounted(() => {
  fetchUsers()
})
</script>

<style scoped>
.user-list-page {
  max-width: 1200px;
}

.page-card {
  background: white;
  border-radius: 12px;
  border: 1px solid #e8e8e8;
  overflow: hidden;
}

.page-card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;
}

.page-title {
  font-size: 16px;
  font-weight: 600;
  color: #1a1a2e;
  margin: 0;
}

.pagination-wrapper {
  display: flex;
  justify-content: flex-end;
  padding: 16px 20px;
  border-top: 1px solid #f0f0f0;
}
</style>
