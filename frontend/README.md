# 订单管理系统前端

基于 Vue 3 + TypeScript + Element Plus 的现代化管理系统前端。

## 技术栈

- **Vue 3** - 渐进式 JavaScript 框架
- **TypeScript** - 类型安全的 JavaScript 超集
- **Element Plus** - 基于 Vue 3 的组件库
- **Vite** - 下一代前端构建工具
- **Vue Router** - Vue.js 官方路由管理器
- **Pinia** - Vue 的状态管理库
- **Axios** - HTTP 客户端

## 功能模块

### 1. 部门管理
- 部门列表展示
- 添加新部门
- 编辑部门信息
- 删除部门
- 部门层级管理

### 2. 菜单管理
- 菜单列表展示
- 添加新菜单
- 编辑菜单信息
- 删除菜单
- 菜单可见性控制
- 菜单权限管理

### 3. 订单管理
- 订单列表展示
- 添加新订单
- 编辑订单信息
- 查看订单详情
- 订单状态管理
- 商品列表管理

### 4. 商品管理
- 商品列表展示
- 添加新商品
- 编辑商品信息
- 查看商品详情
- 商品状态管理
- 价格信息管理

## 项目结构

```
frontend/
├── src/
│   ├── api/              # API 接口
│   │   ├── dept.ts       # 部门相关接口
│   │   ├── menu.ts       # 菜单相关接口
│   │   ├── order.ts      # 订单相关接口
│   │   └── product.ts    # 商品相关接口
│   ├── components/       # 公共组件
│   ├── router/           # 路由配置
│   ├── utils/            # 工具函数
│   │   └── request.ts    # HTTP 请求封装
│   ├── views/            # 页面组件
│   │   ├── dept/         # 部门管理页面
│   │   ├── menu/         # 菜单管理页面
│   │   ├── order/        # 订单管理页面
│   │   └── product/      # 商品管理页面
│   ├── App.vue           # 根组件
│   └── main.ts           # 入口文件
├── public/               # 静态资源
├── package.json          # 项目依赖
├── vite.config.ts        # Vite 配置
└── README.md            # 项目说明
```

## 开发环境

### 环境要求

- Node.js >= 16.0.0
- npm >= 8.0.0

### 安装依赖

```bash
npm install
```

### 启动开发服务器

```bash
npm run dev
```

访问 http://localhost:3000

### 构建生产版本

```bash
npm run build
```

### 代码检查

```bash
npm run lint
```

### 代码格式化

```bash
npm run format
```

## API 接口

项目使用 RESTful API 设计，主要接口包括：

### 部门管理
- `GET /v1/dept` - 获取部门列表
- `POST /v1/dept` - 添加部门
- `PUT /v1/dept/:id` - 编辑部门
- `DELETE /v1/dept/:id` - 删除部门

### 菜单管理
- `GET /api/v1/menus` - 获取菜单列表
- `POST /api/v1/menus` - 添加菜单
- `PUT /api/v1/menus/:id` - 编辑菜单
- `DELETE /api/v1/menus/:id` - 删除菜单

### 订单管理
- `GET /v1/order` - 获取订单列表
- `POST /v1/order` - 添加订单
- `PUT /v1/order` - 编辑订单

### 商品管理
- `GET /v1/product` - 获取商品列表
- `POST /v1/product` - 添加商品
- `PUT /v1/product` - 编辑商品

## 特性

- 🎨 **现代化 UI** - 基于 Element Plus 的美观界面
- 📱 **响应式设计** - 适配各种屏幕尺寸
- 🔒 **类型安全** - 完整的 TypeScript 支持
- ⚡ **高性能** - 基于 Vite 的快速构建
- 🛠️ **开发友好** - 热重载、代码检查等开发工具
- 📦 **模块化** - 清晰的代码组织结构

## 浏览器支持

- Chrome >= 87
- Firefox >= 78
- Safari >= 14
- Edge >= 88

## 许可证

MIT License
