# User Center 前端设计文档

> 状态：已确认
> 日期：2026-06-17
> 后端项目：Go + Gin + GORM + MySQL + JWT

---

## 1. 概述

为现有的 User Center Go 后端项目搭建完整的前端界面。目标是界面优美、交互流畅、代码结构清晰。

---

## 2. 技术选型

| 类别 | 选型 | 说明 |
|------|------|------|
| 框架 | Vue 3 (Composition API) | 最新版本，`<script setup>` 语法 |
| 语言 | TypeScript | 类型安全 |
| 构建工具 | Vite 5 | HMR 极快，Vue 官方推荐 |
| 组件库 | Element Plus | 企业后台经典，组件全面 |
| 状态管理 | Pinia | Vue 3 官方推荐 |
| 路由 | Vue Router 4 | 嵌套路由 + 导航守卫 |
| HTTP | Axios | 拦截器统一处理 |
| 视觉风格 | 现代渐变紫 | 紫色渐变主色调 + 深色侧边栏 + 浅色内容区 |

---

## 3. 页面路由

| 路由路径 | 页面 | 组件 | 需要认证 | 说明 |
|----------|------|------|----------|------|
| `/login` | 登录页 | LoginPage.vue | 否 | 全屏居中卡片 |
| `/register` | 注册页 | RegisterPage.vue | 否 | 全屏居中卡片 |
| `/dashboard` | 仪表盘 | DashboardPage.vue | 是 | 统计卡片 + 快捷入口 |
| `/users` | 用户列表 | UserListPage.vue | 是 | 表格 + 分页 + 搜索 + 删除 |
| `/profile` | 个人中心 | ProfilePage.vue | 是 | 信息展示 + 编辑 |

路由守卫：未登录 → 重定向到 `/login`；已登录访问 login/register → 重定向到 `/dashboard`。

---

## 4. 目录结构

```
frontend/
├── index.html
├── package.json
├── vite.config.ts              # proxy → :8080
├── tsconfig.json
├── src/
│   ├── main.ts                 # 入口，注册 router/pinia/element-plus
│   ├── App.vue
│   ├── router/
│   │   └── index.ts            # 路由配置 + 导航守卫
│   ├── stores/
│   │   ├── user.ts             # 用户状态：token, userInfo, login/logout actions
│   │   └── app.ts              # 应用状态：sidebarCollapsed, loading
│   ├── api/
│   │   ├── request.ts          # axios 实例，拦截器，统一错误处理
│   │   └── user.ts             # 用户相关 API：login, register, list, profile, update, delete
│   ├── pages/
│   │   ├── login/
│   │   │   └── LoginPage.vue
│   │   ├── register/
│   │   │   └── RegisterPage.vue
│   │   ├── dashboard/
│   │   │   └── DashboardPage.vue
│   │   ├── users/
│   │   │   └── UserListPage.vue
│   │   └── profile/
│   │       └── ProfilePage.vue
│   ├── layouts/
│   │   └── AdminLayout.vue     # 侧边栏 + 顶栏 + 内容区
│   ├── components/
│   │   ├── SidebarNav.vue      # 深色侧边栏，logo + 菜单项
│   │   ├── TopHeader.vue       # 面包屑 + 用户头像下拉（退出登录）
│   │   ├── StatsCard.vue       # 可复用的统计卡片
│   │   └── ConfirmDialog.vue   # 确认弹窗（删除等操作）
│   └── utils/
│       └── token.ts            # localStorage token 读写
```

---

## 5. 组件树

```
App.vue
├── <router-view>
│
├── [未登录路由]  无外层布局
│   ├── LoginPage.vue
│   └── RegisterPage.vue
│
└── [已登录路由]  套 AdminLayout.vue
    ├── AdminLayout.vue
    │   ├── SidebarNav.vue          # logo + 菜单
    │   ├── TopHeader.vue           # 面包屑 + 用户下拉
    │   └── <router-view>           # 内容区
    │       ├── DashboardPage.vue
    │       │   └── StatsCard.vue
    │       ├── UserListPage.vue
    │       └── ProfilePage.vue
```

---

## 6. 数据流

### 6.1 登录流程

1. LoginPage 输入用户名密码，调用 `userStore.login()`
2. `userStore` 调用 `api/user.ts` → `POST /api/v1/login`
3. 后端返回 `{user, token}`
4. `userStore` 保存 userInfo，`utils/token.ts` 写 token 到 localStorage
5. Router 跳转到 `/dashboard`

### 6.2 API 请求统一链路

```
组件调用 → api/user.ts
  → request.ts (axios 实例)
    → 请求拦截器：自动附加 Authorization: Bearer <token>
    → Vite proxy：/api/* → http://localhost:8080
    → 响应拦截器：统一检查 code !== 0
      → code === 30001 (TokenInvalid) → 清除 token → 跳转 /login
      → 其他错误 → 返回 Promise.reject
```

### 6.3 Pinia Store 职责

- **userStore**：`userInfo`, `token`, `isLoggedIn`, `login()`, `logout()`, `fetchProfile()`
- **appStore**：`sidebarCollapsed`, `loading`

---

## 7. 页面详细设计

### 7.1 登录页

- 紫色渐变全屏背景
- 居中白色卡片（圆角 16px，阴影）
- 标题 "👤 用户中心" + 副标题 "欢迎回来，请登录"
- 输入框：用户名、密码（el-input, el-input[type=password]）
- 渐变紫色登录按钮
- 底部 "还没有账号？立即注册 →" 链接

### 7.2 注册页

- 与登录页同一风格
- 标题 "✨ 创建账号"
- 输入框：用户名、昵称（可选）、密码
- 渐变紫色注册按钮
- 底部 "已有账号？去登录 →" 链接
- 注册成功后自动跳转登录页

### 7.3 Dashboard 页

- 标题 "📈 数据概览"
- 3 个统计卡片（StatsCard），一行三列：
  - 总用户数：紫色渐变背景，显示数字 + 变化趋势
  - 活跃用户：白色边框
  - 今日新增：白色边框
- 快捷操作区：按钮跳转用户列表 / 编辑个人资料

### 7.4 用户列表页

- el-table 展示用户：ID、用户名、昵称、邮箱、电话、注册时间、操作
- 顶部搜索框（按用户名模糊搜索，如果后端支持搜索参数则拼接，暂按现有 API 适配）
- el-pagination 分页器
- 删除按钮 + el-popconfirm 确认弹窗
- 删除成功后刷新列表

### 7.5 个人中心页

- 左右两栏布局：
  - 左侧：头像（首字母大写）、用户名、昵称、邮箱
  - 右侧：基本信息展示
- 编辑模式切换：右侧从 el-descriptions 切换到 el-form
- 编辑字段：昵称、邮箱、手机、头像地址
- 保存/取消按钮

---

## 8. 错误处理

| 场景 | 处理方式 |
|------|----------|
| 网络错误 | axios 响应拦截器统一 `ElMessage.error('网络异常')` |
| Token 过期 | 响应拦截器 code=30001 → 清除登录态 → 跳转登录页 |
| 表单校验失败 | Element Plus 表单校验，红色提示信息 |
| 业务错误 | `ElMessage.error(message)` 显示后端返回的错误信息 |
| 删除确认 | el-popconfirm 二次确认，防止误删 |

---

## 9. API 对接映射

| 前端调用 | HTTP 方法 | 后端路由 | 请求参数 | 响应 |
|----------|-----------|----------|----------|------|
| `login()` | POST | `/api/v1/login` | `{username, password}` | `{user, token}` |
| `register()` | POST | `/api/v1/register` | `{username, password, nickname, email, phone}` | `user` |
| `listUsers()` | GET | `/api/v1` | `?page=1&page_size=10` | `{list, total, page, page_size}` |
| `getProfile()` | GET | `/api/v1/profile` | — (JWT 自动) | `user` |
| `updateProfile()` | PUT | `/api/v1/profile` | `{nickname, email, phone, avatar}` | `user` |
| `deleteUser()` | DELETE | `/api/v1/:id` | — | `null` |

---

## 10. 不在范围内

- E2E 测试、单元测试（后续迭代添加）
- 国际化 i18n
- 暗黑模式切换（仅做渐变紫一种主题）
- 头像上传功能（使用 URL 输入代替）
- 用户列表高级搜索（仅基本搜索）

---

## 11. 后端需确认的改动

> 注册接口：当前后端 `PUT /profile` 更新时需传入完整字段。前端只传修改的字段（nickname, email, phone, avatar），请确认后端允许部分字段更新。

> 用户列表接口：当前 `GET /api/v1` 的分页参数为 `page` + `page_size`，前端即使用这两个参数名。

---

## 12. 交付物

1. 完整前端项目代码（Vite + Vue 3 + TypeScript）
2. 5 个功能页面
3. 登录认证流程完整可用
4. 与现有 Go 后端对接正常工作
