<template>
  <div class="top-header">
    <div class="breadcrumb">
      <span class="breadcrumb-home" @click="$router.push('/dashboard')">🏠 首页</span>
      <span class="breadcrumb-sep">/</span>
      <span class="breadcrumb-current">{{ pageTitle }}</span>
    </div>
    <div class="header-right">
      <el-dropdown trigger="click" @command="handleCommand">
        <div class="user-avatar-area">
          <div class="avatar-circle">{{ avatarLetter }}</div>
          <span class="username">{{ username }}</span>
          <el-icon class="arrow"><ArrowDown /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">
              <el-icon><User /></el-icon> 个人中心
            </el-dropdown-item>
            <el-dropdown-item command="logout" divided>
              <el-icon><SwitchButton /></el-icon> 退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessageBox } from 'element-plus'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const pageTitle = computed(() => {
  const titles: Record<string, string> = {
    Dashboard: 'Dashboard',
    Users: '用户管理',
    Profile: '个人中心',
  }
  return titles[route.name as string] || 'Dashboard'
})

const username = computed(() => userStore.userInfo?.username || '用户')
const avatarLetter = computed(() => {
  const name = userStore.userInfo?.nickname || userStore.userInfo?.username || 'U'
  return name.charAt(0).toUpperCase()
})

async function handleCommand(command: string) {
  if (command === 'profile') {
    router.push('/profile')
  } else if (command === 'logout') {
    try {
      await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
    } catch {
      return
    }
    userStore.logout()
    router.push('/login')
  }
}
</script>

<style scoped>
.top-header {
  height: 56px;
  background: white;
  border-bottom: 1px solid #e8e8e8;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  flex-shrink: 0;
}

.breadcrumb {
  font-size: 13px;
  color: #666;
  display: flex;
  align-items: center;
  gap: 6px;
}

.breadcrumb-home {
  cursor: pointer;
}

.breadcrumb-home:hover {
  color: #667eea;
}

.breadcrumb-sep {
  color: #ccc;
}

.breadcrumb-current {
  color: #333;
  font-weight: 500;
}

.header-right {
  display: flex;
  align-items: center;
}

.user-avatar-area {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 8px;
  transition: background 0.2s;
}

.user-avatar-area:hover {
  background: #f5f5f5;
}

.avatar-circle {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea, #764ba2);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  font-weight: 600;
}

.username {
  font-size: 13px;
  color: #333;
}

.arrow {
  font-size: 12px;
  color: #999;
}
</style>
