<script setup lang="ts">
import { useRouter } from 'vue-router'
import { 
  OfficeBuilding, 
  Menu, 
  Document, 
  Goods,
  Plus,
  Edit,
  Delete,
  View
} from '@element-plus/icons-vue'

const router = useRouter()

// 统计数据
const stats = [
  {
    title: '部门总数',
    number: '12',
    icon: 'OfficeBuilding',
    type: 'primary'
  },
  {
    title: '菜单总数',
    number: '25',
    icon: 'Menu',
    type: 'success'
  },
  {
    title: '订单总数',
    number: '156',
    icon: 'Document',
    type: 'warning'
  },
  {
    title: '商品总数',
    number: '89',
    icon: 'Goods',
    type: 'info'
  }
]

// 快速导航
const quickNavs = [
  {
    title: '部门管理',
    desc: '管理系统部门结构',
    icon: 'OfficeBuilding',
    path: '/dept'
  },
  {
    title: '菜单管理',
    desc: '管理系统菜单配置',
    icon: 'Menu',
    path: '/menu'
  },
  {
    title: '订单管理',
    desc: '查看和处理订单',
    icon: 'Document',
    path: '/order'
  },
  {
    title: '商品管理',
    desc: '管理商品信息',
    icon: 'Goods',
    path: '/product'
  }
]

// 最近活动
const recentActivities = [
  {
    id: 1,
    title: '新增商品：iPhone 15 Pro',
    time: '2分钟前',
    icon: 'Plus'
  },
  {
    id: 2,
    title: '编辑订单：ORD-2024-001',
    time: '5分钟前',
    icon: 'Edit'
  },
  {
    id: 3,
    title: '删除部门：技术部',
    time: '10分钟前',
    icon: 'Delete'
  },
  {
    id: 4,
    title: '查看菜单：系统管理',
    time: '15分钟前',
    icon: 'View'
  }
]

// 导航点击处理
const handleNavClick = (path: string) => {
  router.push(path)
}
</script>

<template>
  <div class="home-container">
    <el-row :gutter="20">
      <!-- 统计卡片 -->
      <el-col :span="6" v-for="stat in stats" :key="stat.title">
        <el-card class="stat-card" :class="stat.type">
          <div class="stat-content">
            <div class="stat-icon">
              <el-icon :size="40">
                <component :is="stat.icon" />
              </el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-number">{{ stat.number }}</div>
              <div class="stat-title">{{ stat.title }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px;">
      <!-- 快速导航 -->
      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>快速导航</span>
            </div>
          </template>
          <div class="quick-nav">
            <el-row :gutter="20">
              <el-col :span="12" v-for="nav in quickNavs" :key="nav.title">
                <el-card 
                  class="nav-card" 
                  shadow="hover" 
                  @click="handleNavClick(nav.path)"
                >
                  <div class="nav-content">
                    <el-icon :size="30" class="nav-icon">
                      <component :is="nav.icon" />
                    </el-icon>
                    <div class="nav-title">{{ nav.title }}</div>
                    <div class="nav-desc">{{ nav.desc }}</div>
                  </div>
                </el-card>
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>

      <!-- 最近活动 -->
      <el-col :span="12">
        <el-card>
          <template #header>
            <div class="card-header">
              <span>最近活动</span>
            </div>
          </template>
          <div class="recent-activities">
            <div v-for="activity in recentActivities" :key="activity.id" class="activity-item">
              <div class="activity-icon">
                <el-icon :size="20">
                  <component :is="activity.icon" />
                </el-icon>
              </div>
              <div class="activity-content">
                <div class="activity-title">{{ activity.title }}</div>
                <div class="activity-time">{{ activity.time }}</div>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped>
.home-container {
  padding: 20px;
  height: 100%;
  width: 100%;
}

.stat-card {
  height: 120px;
  cursor: pointer;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
}

.stat-card.primary {
  border-left: 4px solid #409eff;
}

.stat-card.success {
  border-left: 4px solid #67c23a;
}

.stat-card.warning {
  border-left: 4px solid #e6a23c;
}

.stat-card.info {
  border-left: 4px solid #909399;
}

.stat-content {
  display: flex;
  align-items: center;
  height: 100%;
}

.stat-icon {
  margin-right: 20px;
  color: #409eff;
}

.stat-info {
  flex: 1;
}

.stat-number {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
}

.stat-title {
  font-size: 14px;
  color: #909399;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.quick-nav {
  padding: 10px 0;
}

.nav-card {
  height: 120px;
  cursor: pointer;
  transition: all 0.3s;
}

.nav-card:hover {
  transform: translateY(-3px);
}

.nav-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  text-align: center;
}

.nav-icon {
  color: #409eff;
  margin-bottom: 10px;
}

.nav-title {
  font-size: 16px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
}

.nav-desc {
  font-size: 12px;
  color: #909399;
}

.recent-activities {
  max-height: 300px;
  overflow-y: auto;
}

.activity-item {
  display: flex;
  align-items: center;
  padding: 15px 0;
  border-bottom: 1px solid #f0f0f0;
}

.activity-item:last-child {
  border-bottom: none;
}

.activity-icon {
  margin-right: 15px;
  color: #409eff;
}

.activity-content {
  flex: 1;
}

.activity-title {
  font-size: 14px;
  color: #303133;
  margin-bottom: 5px;
}

.activity-time {
  font-size: 12px;
  color: #909399;
}
</style>
