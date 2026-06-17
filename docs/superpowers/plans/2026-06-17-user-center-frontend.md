# User Center Frontend Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a complete Vue 3 + TypeScript frontend for the User Center Go backend with login, registration, dashboard, user management, and profile pages.

**Architecture:** Vite dev server proxies `/api` to Go backend on `:8080`. Vue Router handles page navigation with auth guards. Pinia stores manage user state. Axios with interceptors handles JWT token injection and unified error handling. Element Plus provides UI components styled with a purple gradient theme.

**Tech Stack:** Vue 3 (Composition API `<script setup>`), TypeScript, Vite 5, Element Plus, Pinia, Vue Router 4, Axios

---

## File Structure

```
frontend/
├── index.html
├── package.json
├── vite.config.ts
├── tsconfig.json
├── tsconfig.app.json
├── tsconfig.node.json
├── env.d.ts
├── src/
│   ├── main.ts
│   ├── App.vue
│   ├── styles/
│   │   └── global.css
│   ├── types/
│   │   └── user.ts
│   ├── router/
│   │   └── index.ts
│   ├── stores/
│   │   ├── user.ts
│   │   └── app.ts
│   ├── api/
│   │   ├── request.ts
│   │   └── user.ts
│   ├── utils/
│   │   └── token.ts
│   ├── pages/
│   │   ├── login/LoginPage.vue
│   │   ├── register/RegisterPage.vue
│   │   ├── dashboard/DashboardPage.vue
│   │   ├── users/UserListPage.vue
│   │   └── profile/ProfilePage.vue
│   ├── layouts/
│   │   └── AdminLayout.vue
│   └── components/
│       ├── SidebarNav.vue
│       ├── TopHeader.vue
│       └── StatsCard.vue
```

**File boundary reasoning:**
- `types/user.ts` — single source of truth for User, LoginReq, RegisterReq types shared across all modules
- `api/request.ts` — axios factory with interceptors; decoupled from specific API calls so future modules can reuse
- `api/user.ts` — one file per API domain; all user endpoints live here
- `stores/user.ts` — auth state + userInfo; separated from `app.ts` (UI-only state) so auth logic is never mixed with sidebar toggle
- `utils/token.ts` — pure localStorage functions; no dependency on Vue/Pinia so it can be called from anywhere including router guards
- Each page in its own directory — room for page-specific sub-components without polluting `components/`

---

### Task 1: Scaffold Vite + Vue 3 + TypeScript project

**Files:**
- Create: `frontend/` directory and all scaffolding

- [ ] **Step 1: Create the Vite project**

```bash
cd /Users/caixingyue/User_Center && npm create vite@latest frontend -- --template vue-ts
```

Expected: `Scaffolding project in /Users/caixingyue/User_Center/frontend... Done.`

- [ ] **Step 2: Install base dependencies**

```bash
cd /Users/caixingyue/User_Center/frontend && npm install
```

Expected: packages installed without errors.

- [ ] **Step 3: Install project dependencies**

```bash
cd /Users/caixingyue/User_Center/frontend && npm install vue-router@4 pinia axios element-plus @element-plus/icons-vue
```

Expected: All packages added to `package.json` and `node_modules`.

- [ ] **Step 4: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/ && git commit -m "feat: scaffold Vite + Vue 3 + TypeScript frontend project

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 2: Configure Vite, TypeScript, and global styles

**Files:**
- Modify: `frontend/vite.config.ts`
- Modify: `frontend/tsconfig.app.json`
- Create: `frontend/env.d.ts`
- Create: `frontend/src/styles/global.css`

- [ ] **Step 1: Write vite.config.ts with API proxy**

```typescript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
    },
  },
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

- [ ] **Step 2: Write tsconfig.app.json with path aliases**

```json
{
  "compilerOptions": {
    "composite": true,
    "tsBuildInfoFile": "./node_modules/.tmp/tsconfig.app.tsbuildinfo",
    "target": "ES2020",
    "useDefineForExpose": true,
    "module": "ESNext",
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "isolatedModules": true,
    "moduleDetection": "force",
    "noEmit": true,
    "jsx": "preserve",
    "strict": true,
    "noUnusedLocals": false,
    "noUnusedParameters": false,
    "noFallthroughCasesInSwitch": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"]
    }
  },
  "include": ["src/**/*.ts", "src/**/*.tsx", "src/**/*.vue", "env.d.ts"]
}
```

- [ ] **Step 3: Write env.d.ts for Vite client types**

```typescript
/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}
```

- [ ] **Step 4: Write global.css with purple theme variables**

```css
:root {
  --primary-start: #667eea;
  --primary-end: #764ba2;
  --primary-gradient: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  --sidebar-bg: #1a1a2e;
  --sidebar-bg-end: #16213e;
  --sidebar-text: #8890a0;
  --sidebar-active: #a78bfa;
  --content-bg: #f0f2f5;
  --card-radius: 12px;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body, #app {
  height: 100%;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Microsoft YaHei', Arial, sans-serif;
}

body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* Element Plus overrides */
.el-button--primary {
  background: var(--primary-gradient) !important;
  border: none !important;
}

.el-menu-item.is-active {
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.2), rgba(118, 75, 162, 0.2)) !important;
  color: var(--sidebar-active) !important;
}

/* Login / Register fullscreen gradient */
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--primary-gradient);
}

.auth-card {
  background: white;
  border-radius: 16px;
  padding: 40px 36px;
  width: 400px;
  max-width: 90vw;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12);
}

.auth-title {
  font-size: 24px;
  font-weight: 700;
  color: #1a1a2e;
  text-align: center;
  margin-bottom: 4px;
}

.auth-subtitle {
  font-size: 13px;
  color: #999;
  text-align: center;
  margin-bottom: 28px;
}

.auth-link {
  text-align: center;
  margin-top: 18px;
  font-size: 13px;
  color: #999;
}

.auth-link a {
  color: #667eea;
  text-decoration: none;
  cursor: pointer;
}
```

- [ ] **Step 5: Clean up default Vite files**

```bash
cd /Users/caixingyue/User_Center/frontend && rm -f src/style.css src/components/HelloWorld.vue src/assets/vue.svg public/vite.svg
```

- [ ] **Step 6: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/ && git commit -m "feat: configure Vite proxy, TS paths, and global purple theme CSS

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 3: Type definitions and token utility

**Files:**
- Create: `frontend/src/types/user.ts`
- Create: `frontend/src/utils/token.ts`

- [ ] **Step 1: Write types/user.ts**

```typescript
export interface User {
  id: number
  username: string
  nickname: string
  email: string
  phone: string
  avatar: string
  status: number
  created_at: string
  updated_at: string
}

export interface LoginReq {
  username: string
  password: string
}

export interface RegisterReq {
  username: string
  password: string
  nickname?: string
  email?: string
  phone?: string
}

export interface UpdateProfileReq {
  nickname?: string
  email?: string
  phone?: string
  avatar?: string
}

export interface ListUsersParams {
  page: number
  page_size: number
}

export interface ListUsersRes {
  list: User[]
  total: number
  page: number
  page_size: number
}

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

export interface LoginRes {
  user: User
  token: string
}
```

- [ ] **Step 2: Write utils/token.ts**

```typescript
const TOKEN_KEY = 'user_center_token'

export function getToken(): string | null {
  return localStorage.getItem(TOKEN_KEY)
}

export function setToken(token: string): void {
  localStorage.setItem(TOKEN_KEY, token)
}

export function removeToken(): void {
  localStorage.removeItem(TOKEN_KEY)
}
```

- [ ] **Step 3: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/types/ frontend/src/utils/ && git commit -m "feat: add TypeScript type definitions and token utility

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 4: Axios request layer and user API module

**Files:**
- Create: `frontend/src/api/request.ts`
- Create: `frontend/src/api/user.ts`

- [ ] **Step 1: Write api/request.ts**

```typescript
import axios, { type AxiosInstance, type AxiosResponse } from 'axios'
import { ElMessage } from 'element-plus'
import { getToken, removeToken } from '@/utils/token'
import type { ApiResponse } from '@/types/user'

const request: AxiosInstance = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor: attach JWT token
request.interceptors.request.use(
  (config) => {
    const token = getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

// Response interceptor: unwrap {code, message, data}
request.interceptors.response.use(
  (response: AxiosResponse<ApiResponse>) => {
    const { code, message, data } = response.data

    if (code === 0) {
      return data as any
    }

    // Token invalid
    if (code === 30001) {
      removeToken()
      ElMessage.error('登录已过期，请重新登录')
      window.location.href = '/login'
      return Promise.reject(new Error(message))
    }

    ElMessage.error(message || '请求失败')
    return Promise.reject(new Error(message))
  },
  (error) => {
    if (error.response) {
      ElMessage.error(`服务器错误: ${error.response.status}`)
    } else if (error.code === 'ECONNABORTED') {
      ElMessage.error('请求超时，请稍后重试')
    } else {
      ElMessage.error('网络异常，请检查网络连接')
    }
    return Promise.reject(error)
  },
)

export default request
```

- [ ] **Step 2: Write api/user.ts**

```typescript
import request from './request'
import type {
  User,
  LoginReq,
  RegisterReq,
  UpdateProfileReq,
  ListUsersParams,
  ListUsersRes,
  LoginRes,
} from '@/types/user'

export function login(data: LoginReq): Promise<LoginRes> {
  return request.post('/login', data)
}

export function register(data: RegisterReq): Promise<User> {
  return request.post('/register', data)
}

export function listUsers(params: ListUsersParams): Promise<ListUsersRes> {
  return request.get('', { params })
}

export function getProfile(): Promise<User> {
  return request.get('/profile')
}

export function updateProfile(data: UpdateProfileReq): Promise<User> {
  return request.put('/profile', data)
}

export function deleteUser(id: number): Promise<null> {
  return request.delete(`/${id}`)
}
```

- [ ] **Step 3: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/api/ && git commit -m "feat: add axios request layer with JWT interceptor and user API module

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 5: Pinia stores

**Files:**
- Create: `frontend/src/stores/user.ts`
- Create: `frontend/src/stores/app.ts`

- [ ] **Step 1: Write stores/user.ts**

```typescript
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from '@/types/user'
import * as userApi from '@/api/user'
import { setToken, removeToken, getToken } from '@/utils/token'

export const useUserStore = defineStore('user', () => {
  const userInfo = ref<User | null>(null)
  const token = ref<string | null>(getToken())

  const isLoggedIn = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const res = await userApi.login({ username, password })
    token.value = res.token
    userInfo.value = res.user
    setToken(res.token)
  }

  async function register(
    username: string,
    password: string,
    nickname?: string,
    email?: string,
    phone?: string,
  ) {
    await userApi.register({ username, password, nickname, email, phone })
  }

  async function fetchProfile() {
    const user = await userApi.getProfile()
    userInfo.value = user
  }

  async function updateProfile(data: {
    nickname?: string
    email?: string
    phone?: string
    avatar?: string
  }) {
    const user = await userApi.updateProfile(data)
    userInfo.value = user
  }

  function logout() {
    token.value = null
    userInfo.value = null
    removeToken()
  }

  return {
    userInfo,
    token,
    isLoggedIn,
    login,
    register,
    fetchProfile,
    updateProfile,
    logout,
  }
})
```

- [ ] **Step 2: Write stores/app.ts**

```typescript
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const sidebarCollapsed = ref(false)
  const loading = ref(false)

  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  function setLoading(val: boolean) {
    loading.value = val
  }

  return {
    sidebarCollapsed,
    loading,
    toggleSidebar,
    setLoading,
  }
})
```

- [ ] **Step 3: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/stores/ && git commit -m "feat: add Pinia stores for user auth state and app UI state

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 6: Router with auth guards

**Files:**
- Create: `frontend/src/router/index.ts`

- [ ] **Step 1: Write router/index.ts**

```typescript
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { getToken } from '@/utils/token'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/pages/login/LoginPage.vue'),
    meta: { requiresAuth: false },
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/pages/register/RegisterPage.vue'),
    meta: { requiresAuth: false },
  },
  {
    path: '/',
    component: () => import('@/layouts/AdminLayout.vue'),
    meta: { requiresAuth: true },
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/pages/dashboard/DashboardPage.vue'),
        meta: { title: 'Dashboard' },
      },
      {
        path: 'users',
        name: 'Users',
        component: () => import('@/pages/users/UserListPage.vue'),
        meta: { title: '用户管理' },
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/pages/profile/ProfilePage.vue'),
        meta: { title: '个人中心' },
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, _from, next) => {
  const token = getToken()

  if (to.meta.requiresAuth !== false && !token) {
    next('/login')
  } else if ((to.path === '/login' || to.path === '/register') && token) {
    next('/dashboard')
  } else {
    next()
  }
})

export default router
```

- [ ] **Step 2: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/router/ && git commit -m "feat: add Vue Router with lazy-loaded routes and auth guards

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 7: App entry and root component

**Files:**
- Modify: `frontend/src/main.ts`
- Modify: `frontend/src/App.vue`
- Modify: `frontend/index.html`

- [ ] **Step 1: Write main.ts**

```typescript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'
import './styles/global.css'

const app = createApp(App)

// Register all Element Plus icons globally
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(createPinia())
app.use(router)
app.use(ElementPlus, { locale: undefined })
app.mount('#app')
```

- [ ] **Step 2: Write App.vue**

```vue
<template>
  <router-view />
</template>

<script setup lang="ts">
</script>
```

- [ ] **Step 3: Update index.html title**

The `index.html` file already exists from scaffolding. Just verify the `<title>` is set:

```html
<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>用户中心</title>
  </head>
  <body>
    <div id="app"></div>
    <script type="module" src="/src/main.ts"></script>
  </body>
</html>
```

- [ ] **Step 4: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/main.ts frontend/src/App.vue frontend/index.html && git commit -m "feat: wire up app entry with Pinia, Router, Element Plus, and icons

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 8: Layout components — SidebarNav, TopHeader, AdminLayout

**Files:**
- Create: `frontend/src/components/SidebarNav.vue`
- Create: `frontend/src/components/TopHeader.vue`
- Create: `frontend/src/layouts/AdminLayout.vue`

- [ ] **Step 1: Write components/SidebarNav.vue**

```vue
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
```

- [ ] **Step 2: Write components/TopHeader.vue**

```vue
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
```

- [ ] **Step 3: Write layouts/AdminLayout.vue**

```vue
<template>
  <div class="admin-layout">
    <SidebarNav />
    <div class="main-area">
      <TopHeader />
      <div class="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import SidebarNav from '@/components/SidebarNav.vue'
import TopHeader from '@/components/TopHeader.vue'
</script>

<style scoped>
.admin-layout {
  display: flex;
  height: 100vh;
  overflow: hidden;
}

.main-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.content {
  flex: 1;
  overflow-y: auto;
  background: #f0f2f5;
  padding: 24px;
}
</style>
```

- [ ] **Step 4: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/components/SidebarNav.vue frontend/src/components/TopHeader.vue frontend/src/layouts/ && git commit -m "feat: add admin layout with sidebar navigation and top header

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 9: Login page

**Files:**
- Create: `frontend/src/pages/login/LoginPage.vue`

- [ ] **Step 1: Write pages/login/LoginPage.vue**

```vue
<template>
  <div class="auth-page">
    <div class="auth-card">
      <h1 class="auth-title">👤 用户中心</h1>
      <p class="auth-subtitle">欢迎回来，请登录</p>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-position="top"
        @keyup.enter="handleLogin"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="form.username"
            placeholder="请输入用户名"
            :prefix-icon="User"
            size="large"
          />
        </el-form-item>

        <el-form-item label="密码" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            :prefix-icon="Lock"
            size="large"
            show-password
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="submit-btn"
            :loading="loading"
            @click="handleLogin"
          >
            登 录
          </el-button>
        </el-form-item>
      </el-form>

      <div class="auth-link">
        还没有账号？<router-link to="/register">立即注册 →</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { User, Lock } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const formRef = ref<FormInstance>()
const loading = ref(false)

const form = reactive({
  username: '',
  password: '',
})

const rules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 32, message: '用户名长度为 3-32 位', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 32, message: '密码长度为 6-32 位', trigger: 'blur' },
  ],
}

async function handleLogin() {
  if (!formRef.value) return
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    await userStore.login(form.username, form.password)
    ElMessage.success('登录成功')
    router.push('/dashboard')
  } catch (err: any) {
    // Error already shown by interceptor
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.submit-btn {
  width: 100%;
  height: 44px;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 4px;
}
</style>
```

- [ ] **Step 2: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/pages/login/ && git commit -m "feat: add login page with form validation and purple gradient design

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 10: Register page

**Files:**
- Create: `frontend/src/pages/register/RegisterPage.vue`

- [ ] **Step 1: Write pages/register/RegisterPage.vue**

```vue
<template>
  <div class="auth-page">
    <div class="auth-card">
      <h1 class="auth-title">✨ 创建账号</h1>
      <p class="auth-subtitle">加入用户中心</p>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-position="top"
        @keyup.enter="handleRegister"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="form.username"
            placeholder="3-32位用户名"
            :prefix-icon="User"
            size="large"
          />
        </el-form-item>

        <el-form-item label="昵称" prop="nickname">
          <el-input
            v-model="form.nickname"
            placeholder="可选"
            :prefix-icon="UserFilled"
            size="large"
          />
        </el-form-item>

        <el-form-item label="密码" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="6-32位密码"
            :prefix-icon="Lock"
            size="large"
            show-password
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="submit-btn"
            :loading="loading"
            @click="handleRegister"
          >
            注 册
          </el-button>
        </el-form-item>
      </el-form>

      <div class="auth-link">
        已有账号？<router-link to="/login">去登录 →</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { User, UserFilled, Lock } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const formRef = ref<FormInstance>()
const loading = ref(false)

const form = reactive({
  username: '',
  nickname: '',
  password: '',
})

const rules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 32, message: '用户名长度为 3-32 位', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 32, message: '密码长度为 6-32 位', trigger: 'blur' },
  ],
}

async function handleRegister() {
  if (!formRef.value) return
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  loading.value = true
  try {
    await userStore.register(
      form.username,
      form.password,
      form.nickname || undefined,
    )
    ElMessage.success('注册成功，请登录')
    router.push('/login')
  } catch (err: any) {
    // Error already shown by interceptor
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.submit-btn {
  width: 100%;
  height: 44px;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 4px;
}
</style>
```

- [ ] **Step 2: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/pages/register/ && git commit -m "feat: add register page with form validation

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 11: Dashboard page with StatsCard component

**Files:**
- Create: `frontend/src/components/StatsCard.vue`
- Create: `frontend/src/pages/dashboard/DashboardPage.vue`

- [ ] **Step 1: Write components/StatsCard.vue**

```vue
<template>
  <div class="stats-card" :class="{ 'stats-card--highlight': highlight }">
    <div class="stats-label">{{ label }}</div>
    <div class="stats-value">{{ value }}</div>
    <div class="stats-trend" v-if="trend">
      <span :class="trend > 0 ? 'up' : 'down'">{{ trend > 0 ? '↑' : '↓' }} {{ Math.abs(trend) }}%</span>
      <span class="trend-label">{{ trendLabel }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  label: string
  value: string | number
  highlight?: boolean
  trend?: number
  trendLabel?: string
}>()
</script>

<style scoped>
.stats-card {
  background: white;
  border: 1px solid #e8e8e8;
  border-radius: 12px;
  padding: 20px 24px;
  transition: transform 0.2s, box-shadow 0.2s;
}

.stats-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.1);
}

.stats-card--highlight {
  background: linear-gradient(135deg, #667eea, #764ba2);
  border: none;
  color: white;
}

.stats-card--highlight .stats-label {
  color: rgba(255, 255, 255, 0.8);
}

.stats-card--highlight .stats-value {
  color: white;
}

.stats-label {
  font-size: 13px;
  color: #999;
  margin-bottom: 6px;
}

.stats-value {
  font-size: 28px;
  font-weight: 700;
  color: #333;
}

.stats-trend {
  margin-top: 8px;
  font-size: 12px;
}

.up {
  color: #52c41a;
}

.down {
  color: #ff4d4f;
}

.trend-label {
  color: #999;
  margin-left: 4px;
}

.stats-card--highlight .trend-label {
  color: rgba(255, 255, 255, 0.7);
}
</style>
```

- [ ] **Step 2: Write pages/dashboard/DashboardPage.vue**

```vue
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
```

- [ ] **Step 3: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/components/StatsCard.vue frontend/src/pages/dashboard/ && git commit -m "feat: add dashboard page with stats cards and quick actions

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 12: User list page

**Files:**
- Create: `frontend/src/pages/users/UserListPage.vue`

- [ ] **Step 1: Write pages/users/UserListPage.vue**

```vue
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
```

- [ ] **Step 2: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/pages/users/ && git commit -m "feat: add user list page with table, search, pagination, and delete

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 13: Profile page

**Files:**
- Create: `frontend/src/pages/profile/ProfilePage.vue`

- [ ] **Step 1: Write pages/profile/ProfilePage.vue**

```vue
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
```

- [ ] **Step 2: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/src/pages/profile/ && git commit -m "feat: add profile page with view/edit toggle and avatar card

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

### Task 14: Final verification — build check

**Files:** None new.

- [ ] **Step 1: Install dependencies and verify build**

```bash
cd /Users/caixingyue/User_Center/frontend && npm install && npm run build
```

Expected: Build succeeds with no errors. Output in `dist/` directory.

If there are TypeScript or build errors, fix them, then re-run the build until it passes.

- [ ] **Step 2: Verify project file structure**

```bash
cd /Users/caixingyue/User_Center/frontend && find src -type f | sort
```

Expected output should match:

```
src/App.vue
src/api/request.ts
src/api/user.ts
src/components/SidebarNav.vue
src/components/StatsCard.vue
src/components/TopHeader.vue
src/layouts/AdminLayout.vue
src/main.ts
src/pages/dashboard/DashboardPage.vue
src/pages/login/LoginPage.vue
src/pages/profile/ProfilePage.vue
src/pages/register/RegisterPage.vue
src/pages/users/UserListPage.vue
src/router/index.ts
src/stores/app.ts
src/stores/user.ts
src/styles/global.css
src/types/user.ts
src/utils/token.ts
```

- [ ] **Step 3: Verify no stale default files remain**

```bash
cd /Users/caixingyue/User_Center/frontend && ls src/components/HelloWorld.vue 2>&1
```

Expected: `No such file or directory`

- [ ] **Step 4: Commit**

```bash
cd /Users/caixingyue/User_Center && git add frontend/ && git status
```

Review the status. If all looks correct:

```bash
cd /Users/caixingyue/User_Center && git commit -m "feat: complete frontend implementation — all 5 pages, auth flow, API layer

Co-Authored-By: Claude <noreply@anthropic.com>"
```

---

## How to Run

```bash
# Terminal 1: Start Go backend
cd /Users/caixingyue/User_Center && go run cmd/server/main.go
# Backend starts on :8080

# Terminal 2: Start Vite dev server
cd /Users/caixingyue/User_Center/frontend && npm run dev
# Frontend starts on :5173, proxies /api to :8080
```

Open http://localhost:5173 in the browser.
