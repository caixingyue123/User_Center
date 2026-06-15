# Go 语言 35 天项目实战学习计划

> 适合对象：已经学过 Go 基础语法，但缺乏项目实践能力的学习者。  
> 学习目标：从小 Demo 开始，逐步掌握 Go 后端工程实践，最终具备设计和实现中大型系统的能力。

---

## 一、整体学习路线

本计划不单纯按照语法点学习，而是按照真实项目能力递进：

```text
小 Demo → CLI 工具 → HTTP 服务 → 数据库项目 → WebSocket / SSE → Redis / 消息队列 → 微服务拆分 → 中大型系统设计
```

最终你需要完成三个核心项目：

1. **Todo CLI 工具**：练习 Go 基础工程能力。
2. **认证版用户中心**：练习 Web 后端核心能力。
3. **企业知识问答与 AI 运维 Multi-Agent 平台**：练习中大型系统设计与落地能力。

---

# 第一阶段：Go 工程基础与小型 Demo

> 阶段目标：从“会语法”变成“能写结构清晰的小程序”。

---

## Day 1：搭建标准 Go 项目结构

### 今日目标

熟悉真实 Go 项目的基本结构，不再只写单文件 `main.go`。

### 项目任务

创建项目结构：

```bash
go-practice/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── service/
│   ├── model/
│   └── utils/
├── pkg/
├── configs/
├── scripts/
├── go.mod
└── README.md
```

实现一个简单的用户管理 Demo：

```go
type User struct {
    ID   int64
    Name string
    Age  int
}
```

需要实现以下方法：

```go
CreateUser()
GetUserByID()
ListUsers()
DeleteUser()
```

暂时使用内存 `map[int64]User` 作为存储。

### 涉及知识点

- Go 项目目录结构
- `go mod`
- `internal` 和 `pkg` 的区别
- struct 建模
- map 作为临时存储
- 包的拆分
- 函数命名和职责划分

### 今日交付物

一个可以运行的用户管理小程序。

---

## Day 2：错误处理与日志封装

### 今日目标

掌握 Go 中工程化错误处理方式。

### 项目任务

在 Day 1 的用户管理 Demo 中加入错误处理。

例如：

```go
var ErrUserNotFound = errors.New("user not found")
```

实现：

```go
if errors.Is(err, ErrUserNotFound) {
    // handle
}
```

封装一个简单日志工具：

```go
Info()
Warn()
Error()
```

可以先使用标准库 `log`。

### 涉及知识点

- `error`
- `errors.New`
- `fmt.Errorf("%w", err)`
- `errors.Is`
- `errors.As`
- 日志分级
- 不要滥用 `panic`

### 今日交付物

用户不存在、重复创建、参数错误等场景都有明确错误返回。

---

## Day 3：JSON、配置文件与命令行参数

### 今日目标

让程序具备读取配置和 JSON 序列化能力。

### 项目任务

增加配置文件：

```json
{
  "app_name": "go-practice",
  "port": 8080,
  "debug": true
}
```

启动时读取配置，同时实现用户数据导出为 JSON 文件：

```bash
users.json
```

### 涉及知识点

- `encoding/json`
- `os.ReadFile`
- `os.WriteFile`
- 结构体 tag

```go
type Config struct {
    AppName string `json:"app_name"`
    Port    int    `json:"port"`
    Debug   bool   `json:"debug"`
}
```

- 命令行参数 `flag`

### 今日交付物

程序可以读取配置，并把用户列表导出为 JSON。

---

## Day 4：文件读写小项目：Todo CLI

### 今日目标

做一个真正可用的小工具。

### 项目任务

开发一个命令行 Todo 工具：

```bash
todo add "学习 Go"
todo list
todo done 1
todo delete 1
```

数据保存在本地 JSON 文件中：

```json
[
  {
    "id": 1,
    "title": "学习 Go",
    "done": false,
    "created_at": "2026-05-25 10:00:00"
  }
]
```

### 涉及知识点

- CLI 设计
- 文件持久化
- 时间处理 `time`
- 切片增删改查
- JSON 读写
- 简单分层：handler / service / repository

### 今日交付物

一个完整的 Todo 命令行工具。

---

## Day 5：并发基础实战：批量任务处理器

### 今日目标

把 goroutine 和 channel 用起来。

### 项目任务

实现一个批量 URL 检测工具。

输入多个 URL，程序并发检测状态码和耗时：

```bash
go run main.go https://baidu.com https://google.com
```

输出示例：

```text
URL: https://baidu.com
Status: 200
Cost: 120ms
```

### 涉及知识点

- goroutine
- channel
- WaitGroup
- HTTP client
- 超时控制
- 并发数量限制

### 重点练习

不要无限开 goroutine，要实现 worker pool：

```go
jobs := make(chan string)
results := make(chan Result)
```

### 今日交付物

一个支持并发控制的 URL 健康检测工具。

---

## Day 6：Context 超时与取消控制

### 今日目标

掌握后端服务中非常重要的 `context.Context`。

### 项目任务

改造 Day 5 的 URL 检测工具，支持：

- 单个请求超时
- 全局任务取消
- Ctrl+C 优雅退出

### 涉及知识点

- `context.WithTimeout`
- `context.WithCancel`
- `signal.Notify`
- HTTP 请求绑定 context
- 超时和取消的区别

### 今日交付物

一个支持超时和优雅退出的并发任务工具。

---

## Day 7：第一阶段总结项目：本地任务调度器

### 今日目标

综合使用前 6 天知识。

### 项目任务

做一个本地任务调度器，支持：

1. 添加任务
2. 查看任务
3. 定时执行任务
4. 记录执行结果
5. 支持任务超时取消
6. 支持 JSON 文件持久化

示例任务：

```json
{
  "id": 1,
  "name": "check baidu",
  "type": "http_check",
  "target": "https://baidu.com",
  "interval_seconds": 10
}
```

### 涉及知识点

- struct 建模
- JSON 持久化
- goroutine
- ticker
- context
- 错误处理
- 日志

### 今日交付物

一个可以长期运行的本地任务调度器。

---

# 第二阶段：进入 Web 后端开发

> 阶段目标：能独立写 RESTful API，并连接数据库。

---

## Day 8：net/http 写第一个 Web 服务

### 今日目标

不用框架，先理解 HTTP 服务本质。

### 项目任务

基于标准库写用户管理 API：

```text
POST   /users
GET    /users
GET    /users/{id}
DELETE /users/{id}
```

暂时仍然使用内存 map 存储。

### 涉及知识点

- `net/http`
- Handler
- Request / Response
- JSON 请求解析
- JSON 响应
- HTTP 状态码
- RESTful API 设计

### 今日交付物

- 一个标准库版用户管理 HTTP 服务。
- http.HandlerFunc 的职责是什么？
- http.ResponseWriter 和 *http.Request 分别代表什么？
- 为什么创建成功使用 201，删除成功使用 204？
- 为什么业务校验放在 Service，而不是全部写进 Handler？
- 为什么错误判断使用 errors.Is，而不是比较错误字符串？
- 为什么 Web 服务中的共享 map 需要加锁？
- Gin 的路由、JSON 绑定和 JSON 响应，本质上替标准库代码简化了哪些步骤？

---

## Day 9：使用 Gin 或 GoFrame 写 API

### 今日目标

开始接触实际项目常用框架。

### 技术选择建议

- 想快速入门 Web：选 Gin。
- 想贴近企业后端工程：选 GoFrame。

如果是为了实习快速上手，Gin 更轻；如果想做完整工程，GoFrame 更体系化。

### 项目任务

把 Day 8 的用户 API 改成框架版本。

统一响应格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 涉及知识点

- 路由分组
- 中间件
- 参数绑定
- 参数校验
- 统一响应结构

### 今日交付物

一个框架版 RESTful API 项目。

---

## Day 10：MySQL 基础接入

### 今日目标

让数据真正落库。

### 项目任务

用 MySQL 替换内存 map。

设计用户表：

```sql
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL,
    age INT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
```

实现：

```text
创建用户
查询用户
分页查询
删除用户
```

### 涉及知识点

- MySQL 基础
- `database/sql`
- GORM 或 GoFrame ORM
- SQL 注入问题
- 数据库连接池
- DAO / Repository 层

### 今日交付物

用户管理 API 支持 MySQL 持久化。

---

## Day 11：参数校验与统一错误码

### 今日目标

让接口更像真实业务接口。

### 项目任务

新增统一错误码：

```go
const (
    CodeSuccess       = 0
    CodeInvalidParam  = 40001
    CodeUserNotFound  = 40401
    CodeInternalError = 50001
)
```

请求参数校验：

```json
{
  "name": "张三",
  "age": 18
}
```

要求：

- name 不能为空
- age 必须大于 0
- 查询不存在用户返回明确错误

### 涉及知识点

- 参数校验
- 错误码设计
- 业务错误和系统错误区分
- 日志记录
- Response 统一封装

### 今日交付物

接口具备统一错误处理能力。

---

## Day 12：中间件：日志、耗时、Recover

### 今日目标

掌握 Web 服务中间件机制。

### 项目任务

写三个中间件：

1. 请求日志中间件
2. 接口耗时统计中间件
3. panic recover 中间件

输出类似：

```text
method=POST path=/users status=200 cost=12ms
```

### 涉及知识点

- Middleware
- 请求链路
- panic recover
- 请求耗时
- traceId 初步理解

### 今日交付物

服务具备基础可观测能力。

---

## Day 13：分页、搜索和排序

### 今日目标

写出更接近真实后台管理系统的接口。

### 项目任务

用户列表接口支持：

```text
GET /users?page=1&page_size=10&keyword=张&sort=created_at_desc
```

返回：

```json
{
  "list": [],
  "total": 100,
  "page": 1,
  "page_size": 10
}
```

### 涉及知识点

- 分页查询
- 模糊搜索
- 排序
- limit offset
- SQL 条件拼接
- 防止非法排序字段

### 今日交付物

用户列表接口具备分页搜索排序能力。

---

## Day 14：第二阶段总结项目：用户中心小系统

### 今日目标

完成第一个可展示的小型后端项目。

### 项目内容

实现一个用户中心：

1. 用户注册
2. 用户登录
3. 用户信息查询
4. 用户列表分页
5. 用户删除
6. 参数校验
7. MySQL 持久化
8. 统一响应
9. 错误码
10. 日志中间件

### 涉及知识点

- HTTP API
- MySQL
- 分层架构
- 参数校验
- 错误处理
- 中间件
- RESTful 设计

### 今日交付物

一个可以写进简历的基础后端项目雏形。

---

# 第三阶段：认证、缓存、并发与工程化

> 阶段目标：开始掌握真实后端系统常用组件。

---

## Day 15：JWT 登录认证

### 今日目标

实现登录态。

### 项目任务

用户登录成功后返回 JWT：

```json
{
  "token": "xxxxx"
}
```

需要登录的接口：

```text
GET /profile
PUT /profile
```

请求头：

```text
Authorization: Bearer token
```

### 涉及知识点

- JWT 原理
- Token 生成和解析
- 登录态
- 认证中间件
- 用户上下文
- Token 过期时间

### 今日交付物

用户中心支持 JWT 登录认证。

---

## Day 16：密码加密与安全意识

### 今日目标

不要明文存密码。

### 项目任务

注册时使用 bcrypt 加密密码，登录时验证密码。

### 涉及知识点

- bcrypt
- 密码哈希
- 盐值
- 明文密码不能入库
- 登录失败错误提示设计
- 简单防爆破思路

### 今日交付物

用户注册登录具备基础安全性。

---

## Day 17：Redis 缓存用户信息

### 今日目标

学习缓存的基本使用。

### 项目任务

查询用户详情时：

1. 先查 Redis
2. Redis 没有再查 MySQL
3. 查到后写入 Redis
4. 更新用户信息时删除缓存

缓存 key：

```text
user:info:{user_id}
```

### 涉及知识点

- Redis 基础
- go-redis
- 缓存旁路模式 Cache Aside
- 缓存穿透
- 缓存击穿
- 缓存雪崩
- TTL 设置

### 今日交付物

用户详情接口支持 Redis 缓存。

---

## Day 18：限流中间件

### 今日目标

实现接口保护。

### 项目任务

写一个简单限流中间件：

- 每个 IP 每分钟最多请求 60 次
- 超过返回 429

可以先用内存 map 实现，再升级 Redis。

### 涉及知识点

- 限流
- 滑动窗口
- 固定窗口
- 令牌桶
- IP 获取
- Redis 计数器
- 过期时间

### 今日交付物

服务具备基础限流能力。

---

## Day 19：后台异步任务

### 今日目标

理解同步和异步的区别。

### 项目任务

用户注册后，不直接在接口里做所有事情，而是投递一个异步任务。

例如：

```text
用户注册成功 → 异步发送欢迎通知
```

先用 channel 模拟任务队列。

### 涉及知识点

- channel 队列
- worker pool
- 异步任务
- 任务失败重试
- 优雅关闭
- 生产者消费者模型

### 今日交付物

用户注册后支持异步任务处理。

---

## Day 20：消息队列思想，用 Redis Stream 模拟 MQ

### 今日目标

理解消息队列在系统中的作用。

### 项目任务

把 Day 19 的 channel 队列替换为 Redis Stream。

流程：

```text
注册接口 → 写入 Redis Stream → worker 消费 → 处理任务
```

### 涉及知识点

- 消息队列
- Redis Stream
- Consumer Group
- 消费确认 ACK
- 消息重试
- 幂等处理
- 异步解耦

### 今日交付物

系统具备简易 MQ 异步处理能力。

---

## Day 21：第三阶段总结项目：认证版用户中心

### 今日目标

完成一个比较完整的用户中心系统。

### 项目功能

1. 注册
2. 登录
3. JWT 鉴权
4. 密码加密
5. 用户信息缓存
6. 接口限流
7. 异步任务
8. Redis Stream 消息队列
9. MySQL 持久化
10. 中间件日志

### 今日交付物

一个完整的 Go 后端基础项目，可以作为后续中大型系统的基础模块。

---

# 第四阶段：实时通信与流式输出

> 阶段目标：掌握 WebSocket 和 SSE，这对 AI 应用、MV 项目、聊天系统都很有用。

---

## Day 22：SSE 流式接口

### 今日目标

掌握服务端流式返回。

### 项目任务

实现一个模拟 AI 输出接口：

```text
GET /chat/stream
```

服务端每隔 300ms 返回一段内容：

```text
data: 你好

data: 我是

data: AI 助手
```

### 涉及知识点

- SSE
- `text/event-stream`
- `http.Flusher`
- 长连接
- 流式响应
- 客户端断开检测
- context cancel

### 今日交付物

一个可以模拟 ChatGPT 打字效果的 SSE 接口。

---

## Day 23：SSE 对话上下文管理

### 今日目标

把 SSE 变成真实业务雏形。

### 项目任务

设计对话模型：

```sql
conversations
messages
```

实现：

```text
POST /conversations
POST /conversations/{id}/messages/stream
GET  /conversations/{id}/messages
```

流式输出过程中，每个 chunk 保存或最后整体保存。

### 涉及知识点

- conversationId
- messageId
- 上下文管理
- 流式输出落库
- 客户端断开处理
- 幂等设计

### 今日交付物

一个简易 AI 对话后端。

---

## Day 24：WebSocket 入门：聊天室

### 今日目标

掌握 WebSocket 基础。

### 项目任务

实现一个聊天室：

1. 用户连接
2. 加入房间
3. 发送消息
4. 广播消息
5. 用户断开

### 涉及知识点

- WebSocket 协议
- gorilla/websocket
- 连接管理
- 读写协程
- 心跳
- 广播
- 房间模型

### 今日交付物

一个本地可运行的 WebSocket 聊天室。

---

## Day 25：WebSocket 网关设计

### 今日目标

让聊天室结构更工程化。

### 项目任务

抽象以下模块：

```text
Client
Hub
Room
Message
ConnectionManager
```

实现：

```text
用户上线
用户下线
房间广播
私聊消息
心跳检测
```

### 涉及知识点

- 长连接管理
- map 并发安全
- sync.RWMutex
- channel 广播
- 心跳机制
- 连接清理
- 消息协议设计

### 今日交付物

一个结构清晰的 WebSocket 网关雏形。

---

## Day 26：第四阶段总结项目：实时对话系统

### 今日目标

把 SSE 和 WebSocket 结合起来理解。

### 项目内容

实现一个实时对话系统：

- 普通聊天用 WebSocket
- AI 流式回复用 SSE
- 消息落库 MySQL
- 用户登录 JWT
- Redis 缓存用户状态
- 心跳检测
- 历史消息查询

### 今日交付物

你应该能清楚解释：

```text
SSE 适合服务端单向流式输出
WebSocket 适合客户端和服务端双向实时通信
```

---

# 第五阶段：中大型系统项目：企业知识问答与 AI 运维平台

> 阶段目标：完成一个能写进简历、能面试讲清楚的中大型系统。

---

## Day 27：系统需求分析与架构设计

### 今日目标

不要急着写代码，先设计系统。

### 项目背景

企业内部有大量文档、接口说明、运维手册、错误日志。用户可以上传文档，系统把文档向量化后存入知识库。用户提问时，系统通过 RAG 检索相关内容，并通过大模型生成答案。对于运维类问题，系统可以调用 Prometheus 工具查询监控指标。

### 核心模块

```text
用户中心
文档中心
知识库中心
问答中心
Agent 调度中心
工具调用中心
监控中心
```

### 架构图文字版

```text
前端
  ↓
API Gateway
  ↓
用户服务 ─ MySQL / Redis
  ↓
文档服务 ─ MinIO / 本地文件
  ↓
知识库服务 ─ Milvus
  ↓
问答服务 ─ LLM / SSE
  ↓
Agent 服务 ─ Tool / Prometheus / MCP
```

### 涉及知识点

- 中大型系统模块拆分
- 单体架构和微服务架构区别
- API Gateway 思想
- RAG 架构
- Agent 架构
- 数据流设计
- 表结构设计

### 今日交付物

写一份系统设计文档，包括：

1. 项目背景
2. 核心功能
3. 技术架构
4. 数据库设计
5. 接口设计
6. 核心流程图

---

## Day 28：项目初始化与基础脚手架

### 今日目标

搭建中大型项目骨架。

### 推荐目录

```bash
ai-ops-platform/
├── cmd/
│   ├── api/
│   ├── worker/
│   └── agent/
├── internal/
│   ├── user/
│   ├── auth/
│   ├── document/
│   ├── knowledge/
│   ├── chat/
│   ├── agent/
│   ├── tool/
│   └── common/
├── pkg/
│   ├── response/
│   ├── logger/
│   ├── jwt/
│   ├── redisx/
│   └── mysqlx/
├── configs/
├── deployments/
│   └── docker-compose.yml
├── docs/
└── README.md
```

### 项目任务

搭建：

- 配置加载
- 日志模块
- MySQL 连接
- Redis 连接
- 统一响应
- 错误码
- 路由注册
- 健康检查接口

接口：

```text
GET /health
```

### 涉及知识点

- 工程初始化
- 配置管理
- 基础设施封装
- 路由分组
- 模块化组织
- Docker Compose 初步

### 今日交付物

中大型项目的基础骨架。

---

## Day 29：用户中心模块

### 今日目标

把之前用户中心能力迁移进大项目。

### 项目任务

实现：

```text
POST /api/v1/auth/register
POST /api/v1/auth/login
GET  /api/v1/user/profile
PUT  /api/v1/user/profile
```

### 数据表

```sql
users
user_tokens
```

### 涉及知识点

- 用户模块拆分
- JWT
- bcrypt
- 登录鉴权中间件
- Redis 存登录态，可选
- 用户上下文传递

### 今日交付物

平台具备完整登录能力。

---

## Day 30：文档上传与解析模块

### 今日目标

实现知识库入口。

### 项目任务

实现文档上传接口：

```text
POST   /api/v1/documents/upload
GET    /api/v1/documents
GET    /api/v1/documents/{id}
DELETE /api/v1/documents/{id}
```

支持上传：

- `.txt`
- `.md`
- `.pdf` 可以先预留，后面再做

文档表：

```sql
documents
```

字段：

```text
id
user_id
filename
file_path
file_type
status
created_at
updated_at
```

### 涉及知识点

- 文件上传
- multipart/form-data
- 文件存储
- 文档元数据
- 文件大小限制
- 上传状态设计

### 今日交付物

用户可以上传文档并查看文档列表。

---

## Day 31：文本切分与知识库入库

### 今日目标

实现 RAG 的数据准备阶段。

### 项目任务

文档上传后，异步执行：

```text
文档读取 → 文本清洗 → 文本切分 → embedding → Milvus 入库
```

先不用真的调用大模型 embedding，可以先 mock：

```go
func Embed(text string) []float32 {
    return make([]float32, 768)
}
```

设计 chunk 表：

```sql
document_chunks
```

字段：

```text
id
document_id
chunk_index
content
embedding_id
created_at
```

状态流转：

```text
uploaded
processing
indexed
failed
```

### 涉及知识点

- 文本切分
- chunk 设计
- 异步任务
- embedding 思想
- Milvus 基础概念
- 向量数据库
- 文档状态流转

### 今日交付物

文档可以被切分并写入知识库。

---

## Day 32：知识库检索模块

### 今日目标

实现 RAG 的 Retrieval 阶段。

### 项目任务

实现接口：

```text
POST /api/v1/knowledge/search
```

请求：

```json
{
  "query": "如何排查接口超时？",
  "top_k": 5
}
```

返回：

```json
{
  "chunks": [
    {
      "document_id": 1,
      "content": "...",
      "score": 0.89
    }
  ]
}
```

如果 Milvus 暂时不会接，可以先用 MySQL 模糊搜索模拟，再升级为向量检索。

### 涉及知识点

- RAG 检索流程
- topK
- 相似度分数
- query embedding
- 向量检索
- fallback 检索
- metadata 过滤

### 今日交付物

系统可以根据问题检索相关文档片段。

---

## Day 33：SSE 问答模块

### 今日目标

实现一个完整的知识问答链路。

### 项目任务

接口：

```text
POST /api/v1/chat/stream
```

流程：

```text
用户提问
  ↓
检索知识库 topK chunks
  ↓
构造 Prompt
  ↓
调用 LLM，模拟也可以
  ↓
SSE 流式返回答案
  ↓
保存问题和回答
```

Prompt 模板：

```text
你是企业知识库问答助手。
请基于以下资料回答用户问题。

资料：
{{retrieved_chunks}}

用户问题：
{{question}}

要求：
1. 如果资料中没有答案，请说明不知道。
2. 不要编造。
3. 回答要简洁清晰。
```

### 涉及知识点

- RAG 问答链路
- Prompt 拼接
- SSE 流式输出
- 对话记录
- 上下文管理
- token 长度控制
- 引用来源返回

### 今日交付物

一个完整的企业知识库问答功能。

---

## Day 34：Agent 工具调用模块

### 今日目标

让系统不仅能问文档，还能调用工具。

### 项目任务

设计工具接口：

```go
type Tool interface {
    Name() string
    Description() string
    Call(ctx context.Context, input string) (string, error)
}
```

实现几个工具：

1. 查询当前时间工具
2. 查询服务健康状态工具
3. 查询 Prometheus 指标工具，先 mock
4. 查询最近错误日志工具，先 mock

Agent 流程：

```text
用户问题
  ↓
判断是否需要工具
  ↓
调用 Tool
  ↓
整合工具结果
  ↓
生成回答
```

示例问题：

```text
帮我看看 user-service 最近是否有异常
```

### 涉及知识点

- Agent 基础
- Tool 抽象
- ReAct 思想
- 工具注册表
- 工具调用参数
- 结果整合
- 超时控制

### 今日交付物

系统支持简单 Agent 工具调用。

---

## Day 35：最终整合：中大型系统设计复盘

### 今日目标

把项目整理成可以面试讲、简历写、继续扩展的系统。

### 项目最终功能

你的系统应该具备：

1. 用户注册登录
2. JWT 鉴权
3. 文档上传
4. 文档切分
5. 知识库索引
6. 知识库检索
7. RAG 问答
8. SSE 流式输出
9. 对话历史保存
10. Redis 缓存
11. 异步任务
12. Agent 工具调用
13. 日志中间件
14. 限流中间件
15. Docker Compose 部署

### 最终项目结构

```text
用户层：
前端 / Postman / Apifox

接入层：
HTTP API / JWT / 限流 / 日志

业务层：
用户中心 / 文档中心 / 知识库中心 / 问答中心 / Agent 中心

数据层：
MySQL / Redis / Milvus / 文件存储

外部能力：
LLM / Embedding / Prometheus / MCP Tool
```

### 今日交付物

整理以下内容：

1. README
2. 接口文档
3. 数据库表结构
4. 系统架构图
5. 核心流程图
6. 项目亮点
7. 简历描述
8. 面试讲解稿

---

# 每周复盘重点

## 第 1 周复盘

你应该掌握：

- Go 项目结构
- JSON 文件读写
- 错误处理
- goroutine
- channel
- context
- 简单任务调度器

核心能力：

> 能写结构清晰的小型 Go 程序。

---

## 第 2 周复盘

你应该掌握：

- HTTP API
- Web 框架
- MySQL
- RESTful
- 分页查询
- 参数校验
- 错误码
- 中间件

核心能力：

> 能独立写一个 CRUD 后端服务。

---

## 第 3 周复盘

你应该掌握：

- JWT
- bcrypt
- Redis 缓存
- 限流
- 异步任务
- Redis Stream
- worker pool

核心能力：

> 能把项目做得更像真实企业后端服务。

---

## 第 4 周复盘

你应该掌握：

- SSE
- WebSocket
- 长连接管理
- 心跳机制
- 流式输出
- 对话上下文
- 消息落库

核心能力：

> 能做 AI 对话、聊天室、实时系统这类项目。

---

## 第 5 周复盘

你应该掌握：

- 系统设计
- 模块拆分
- RAG
- 文档处理
- 向量检索
- Agent 工具调用
- Docker Compose
- 中大型项目组织方式

核心能力：

> 能设计并实现一个中大型 Go 后端系统。

---

# 最终建议重点完成的 3 个项目

## 项目一：Todo CLI

适合练基础工程能力。

技术点：

```text
Go 基础工程结构
JSON 文件读写
命令行参数
错误处理
```

---

## 项目二：认证版用户中心

适合练 Web 后端核心能力。

技术点：

```text
HTTP API
MySQL
Redis
JWT
bcrypt
中间件
限流
异步任务
```

---

## 项目三：企业知识问答与 AI 运维平台

适合写进简历和面试。

技术点：

```text
Go
SSE
RAG
MySQL
Redis
Milvus
Agent
Prometheus
MCP
Docker
```

简历描述可以写成：

```text
企业知识问答与 AI 运维 Multi-Agent 平台
基于 Go 构建企业级知识问答与智能运维平台，支持文档上传、文本切分、向量化检索、RAG 问答、SSE 流式输出和 Agent 工具调用。系统采用 MySQL 存储业务数据，Redis 实现缓存与异步任务队列，Milvus 存储文档向量，结合 Prometheus 工具查询实现基础 AI Ops 能力。项目中设计了统一错误码、JWT 鉴权、中间件日志、接口限流、文档状态流转和异步索引任务，提高了系统的可维护性和扩展性。
```

---

# 每天学习的固定流程

建议每天按照这个节奏：

```text
1. 先看当天知识点，控制在 30~60 分钟。
2. 直接写代码，不要一直看教程。
3. 写完接口后用 Postman / Apifox 测试。
4. 给项目补 README。
5. 当天晚上复盘：今天用了哪些技术，遇到什么问题，怎么解决。
```

每天复盘可以回答 5 个问题：

```text
1. 今天实现了什么功能？
2. 今天用到了哪些 Go 知识点？
3. 哪个地方卡住了？
4. 这个功能在真实项目里有什么用？
5. 如果面试官问这个模块，我怎么讲？
```

---

# 学完之后应达到的水平

如果认真把这 35 天做完，你的能力应该达到：

```text
1. 能独立搭建 Go 后端项目。
2. 能写 RESTful API。
3. 能使用 MySQL 和 Redis。
4. 能处理 JWT 登录认证。
5. 能写中间件。
6. 能做异步任务。
7. 能理解 SSE 和 WebSocket。
8. 能设计简单消息队列模型。
9. 能做一个 RAG 问答系统。
10. 能讲清楚中大型系统架构。
```

到这个阶段，你再去看公司里的 Go 项目，就不会只停留在“看懂语法”，而是能理解：

```text
这个模块为什么这么拆？
这个中间件解决什么问题？
这个缓存 key 为什么这么设计？
这个异步任务为什么不直接同步执行？
这个系统为什么要拆成用户、文档、知识库、问答、Agent 几个模块？
```

这才是从 Go 基础走向 Go 工程能力的关键。
