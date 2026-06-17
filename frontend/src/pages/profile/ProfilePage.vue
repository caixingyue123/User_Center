<template>
  <div class="profile-page">
    <h2 class="page-title">👤 个人中心</h2>

    <div class="profile-layout">
      <!-- Left: avatar card -->
      <div class="profile-card">
        <div class="avatar-xl">{{ avatarLetter }}</div>
        <div class="profile-name">{{ userStore.userInfo?.username }}</div>
        <div class="profile-role">{{ userStore.userInfo?.nickname || '用户' }}</div>
        <div class="profile-email">{{ userStore.userInfo?.email || '未设置邮箱' }}</div>
        <div class="profile-status">
          <span class="status-dot"></span> 正常
        </div>
      </div>

      <!-- Right: info / edit -->
      <div class="profile-detail">
        <div class="detail-header">
          <h3>基本信息</h3>
          <el-button
            :type="editing ? 'default' : 'primary'"
            size="small"
            @click="toggleEdit"
          >
            {{ editing ? '取消' : '✏️ 编辑' }}
          </el-button>
        </div>

        <!-- View mode -->
        <div v-if="!editing" class="info-grid">
          <div class="info-item">
            <span class="info-label">用户名</span>
            <span class="info-value">{{ userStore.userInfo?.username }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">昵称</span>
            <span class="info-value">{{ userStore.userInfo?.nickname || '-' }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">邮箱</span>
            <span class="info-value">{{ userStore.userInfo?.email || '-' }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">手机</span>
            <span class="info-value">{{ userStore.userInfo?.phone || '-' }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">头像地址</span>
            <span class="info-value">{{ userStore.userInfo?.avatar || '-' }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">注册时间</span>
            <span class="info-value">{{ formatDate(userStore.userInfo?.created_at) }}</span>
          </div>
        </div>

        <!-- Edit mode -->
        <el-form
          v-else
          ref="formRef"
          :model="form"
          label-width="80px"
          class="edit-form"
        >
          <el-form-item label="昵称" prop="nickname">
            <el-input v-model="form.nickname" placeholder="请输入昵称" maxlength="32" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email" placeholder="请输入邮箱" />
          </el-form-item>
          <el-form-item label="手机" prop="phone">
            <el-input v-model="form.phone" placeholder="请输入手机号" />
          </el-form-item>
          <el-form-item label="头像地址" prop="avatar">
            <el-input v-model="form.avatar" placeholder="请输入头像URL" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="saving" @click="handleSave">
              保 存
            </el-button>
            <el-button @click="toggleEdit">取 消</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, reactive, onMounted } from 'vue'
import { ElMessage, type FormInstance } from 'element-plus'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const formRef = ref<FormInstance>()
const editing = ref(false)
const saving = ref(false)

const form = reactive({
  nickname: '',
  email: '',
  phone: '',
  avatar: '',
})

const avatarLetter = computed(() => {
  const name = userStore.userInfo?.nickname || userStore.userInfo?.username || 'U'
  return name.charAt(0).toUpperCase()
})

function formatDate(dateStr?: string): string {
  if (!dateStr) return '-'
  const d = new Date(dateStr)
  const pad = (n: number) => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}`
}

function toggleEdit() {
  if (editing.value) {
    editing.value = false
  } else {
    form.nickname = userStore.userInfo?.nickname || ''
    form.email = userStore.userInfo?.email || ''
    form.phone = userStore.userInfo?.phone || ''
    form.avatar = userStore.userInfo?.avatar || ''
    editing.value = true
  }
}

async function handleSave() {
  saving.value = true
  try {
    await userStore.updateProfile({
      nickname: form.nickname || undefined,
      email: form.email || undefined,
      phone: form.phone || undefined,
      avatar: form.avatar || undefined,
    })
    ElMessage.success('保存成功')
    editing.value = false
  } catch {
    // Error handled by interceptor
  } finally {
    saving.value = false
  }
}

onMounted(async () => {
  if (!userStore.userInfo) {
    try {
      await userStore.fetchProfile()
    } catch {
      // Error handled by interceptor
    }
  }
})
</script>

<style scoped>
.profile-page {
  max-width: 900px;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: #1a1a2e;
  margin-bottom: 20px;
}

.profile-layout {
  display: flex;
  gap: 20px;
}

.profile-card {
  width: 220px;
  background: white;
  border-radius: 12px;
  border: 1px solid #e8e8e8;
  padding: 28px 20px;
  text-align: center;
  flex-shrink: 0;
  align-self: flex-start;
}

.avatar-xl {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  font-weight: 700;
  margin: 0 auto 14px;
}

.profile-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.profile-role {
  font-size: 13px;
  color: #999;
  margin-top: 4px;
}

.profile-email {
  font-size: 12px;
  color: #667eea;
  margin-top: 14px;
  word-break: break-all;
}

.profile-status {
  margin-top: 14px;
  font-size: 13px;
  color: #52c41a;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
}

.status-dot {
  width: 8px;
  height: 8px;
  background: #52c41a;
  border-radius: 50%;
}

.profile-detail {
  flex: 1;
  background: white;
  border-radius: 12px;
  border: 1px solid #e8e8e8;
  padding: 24px;
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.detail-header h3 {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 18px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.info-label {
  font-size: 12px;
  color: #999;
}

.info-value {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.edit-form {
  max-width: 400px;
}

@media (max-width: 768px) {
  .profile-layout {
    flex-direction: column;
  }

  .profile-card {
    width: 100%;
  }

  .info-grid {
    grid-template-columns: 1fr;
  }
}
</style>
