<script setup lang="ts">
import { useRoute } from 'vue-router'
import { OfficeBuilding, Menu, Document, Goods } from '@element-plus/icons-vue'

const route = useRoute()

const getCurrentPageTitle = () => {
  const routeMap: Record<string, string> = {
    '/dept': '部门管理',
    '/menu': '菜单管理',
    '/order': '订单管理',
    '/product': '商品管理'
  }
  return routeMap[route.path] || '首页'
}
</script>

<template>
  <el-container class="app-container">
    <!-- 侧边栏 -->
    <el-aside width="200px" class="sidebar">
      <div class="logo">
        <h2>订单管理系统</h2>
      </div>
      <el-menu
        :default-active="$route.path"
        class="sidebar-menu"
        router
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409EFF"
      >
        <el-menu-item index="/dept">
          <el-icon><OfficeBuilding /></el-icon>
          <span>部门管理</span>
        </el-menu-item>
        <el-menu-item index="/menu">
          <el-icon><Menu /></el-icon>
          <span>菜单管理</span>
        </el-menu-item>
        <el-menu-item index="/order">
          <el-icon><Document /></el-icon>
          <span>订单管理</span>
        </el-menu-item>
        <el-menu-item index="/product">
          <el-icon><Goods /></el-icon>
          <span>商品管理</span>
        </el-menu-item>
      </el-menu>
    </el-aside>

    <!-- 主内容区域 -->
    <el-container class="main-container">
      <!-- 顶部导航 -->
      <el-header class="header">
        <div class="header-left">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{ getCurrentPageTitle() }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <div class="header-right">
          <el-dropdown>
            <span class="user-info">
              <el-avatar size="small" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
              <span>管理员</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>个人信息</el-dropdown-item>
                <el-dropdown-item>修改密码</el-dropdown-item>
                <el-dropdown-item divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 主要内容 -->
      <el-main class="main-content">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<style>
html, body, #app {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  overflow: hidden;
}
#app {
  width: 100vw !important;
  min-width: 0;
  max-width: 100vw;
  height: 100vh;
  min-height: 0;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>

<style scoped>
.app-container {
  height: 100vh;
  width: 100vw;
  display: flex;
  min-width: 0;
  min-height: 0;
  overflow: hidden;
}
.sidebar {
  height: 100vh;
  min-height: 0;
  min-width: 0;
  overflow-y: auto;
}
.main-container {
  display: flex;
  flex-direction: column;
  flex: 1 1 0%;
  min-width: 0;
  min-height: 0;
  height: 100vh;
  width: 100%;
  overflow: hidden;
}
.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #2b2f3a;
  color: #fff;
}
.logo h2 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
}
.sidebar-menu {
  border-right: none;
}
.header {
  background-color: #fff;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  height: 60px;
  flex-shrink: 0;
  min-width: 0;
  min-height: 0;
}
.header-left {
  display: flex;
  align-items: center;
}
.header-right {
  display: flex;
  align-items: center;
}
.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
  padding: 0 10px;
}
.user-info span {
  margin-left: 8px;
  color: #606266;
}
.main-content {
  flex: 1 1 0%;
  min-width: 0;
  min-height: 0;
  overflow-y: auto;
  height: 0;
  width: 100%;
  background: #f0f2f5;
  padding: 20px;
}
</style>
