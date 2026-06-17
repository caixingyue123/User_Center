<template>
  <div class="sidebar">
    <div class="sidebar-logo">
      <span class="logo-text">👤 User Center</span>
    </div>
    <div class="sidebar-menu">
      <router-link
        v-for="item in menuItems"
        :key="item.path"
        :to="item.path"
        class="menu-item"
        :class="{ active: isActive(item.path) }"
      >
        <el-icon class="menu-icon"><component :is="item.icon" /></el-icon>
        <span>{{ item.label }}</span>
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()

const menuItems = [
  { path: '/dashboard', label: 'Dashboard', icon: 'DataAnalysis' },
  { path: '/users', label: '用户管理', icon: 'User' },
  { path: '/profile', label: '个人中心', icon: 'UserFilled' },
]

function isActive(path: string): boolean {
  return route.path.startsWith(path)
}
</script>

<style scoped>
.sidebar {
  width: 220px;
  background: linear-gradient(180deg, #1a1a2e 0%, #16213e 100%);
  color: white;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
  overflow-y: auto;
}

.sidebar-logo {
  padding: 20px 20px 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  margin-bottom: 8px;
}

.logo-text {
  font-size: 18px;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea, #a78bfa);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.sidebar-menu {
  padding: 4px 0;
  flex: 1;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 20px;
  margin: 4px 12px;
  border-radius: 8px;
  color: #8890a0;
  text-decoration: none;
  font-size: 14px;
  transition: all 0.2s ease;
}

.menu-item:hover {
  color: #a78bfa;
  background: rgba(102, 126, 234, 0.1);
}

.menu-item.active {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.2), rgba(118, 75, 162, 0.2));
  color: #a78bfa;
  font-weight: 600;
}

.menu-icon {
  font-size: 18px;
}
</style>
