<template>
  <div class="menu-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>菜单管理</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            添加菜单
          </el-button>
        </div>
      </template>

      <!-- 菜单列表 -->
      <el-table :data="menuList" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="菜单名称" />
        <el-table-column prop="routeName" label="路由名称" width="150" />
        <el-table-column prop="routePath" label="路由路径" width="150" />
        <el-table-column prop="component" label="组件" width="150" />
        <el-table-column prop="parentId" label="父菜单ID" width="100" />
        <el-table-column prop="visible" label="可见性" width="80">
          <template #default="{ row }">
            <el-tag :type="row.visible === 1 ? 'success' : 'info'">
              {{ row.visible === 1 ? '可见' : '隐藏' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="icon" label="图标" width="100" />
        <el-table-column prop="createAt" label="创建时间" width="180" />
        <el-table-column label="操作" width="250" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button 
              size="small" 
              :type="row.visible === 1 ? 'warning' : 'success'"
              @click="handleToggleVisibility(row)"
            >
              {{ row.visible === 1 ? '隐藏' : '显示' }}
            </el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑菜单对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="resetForm"
    >
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="菜单名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入菜单名称" />
        </el-form-item>
        <el-form-item label="路由名称" prop="routeName">
          <el-input v-model="form.routeName" placeholder="请输入路由名称" />
        </el-form-item>
        <el-form-item label="路由路径" prop="routePath">
          <el-input v-model="form.routePath" placeholder="请输入路由路径" />
        </el-form-item>
        <el-form-item label="组件" prop="component">
          <el-input v-model="form.component" placeholder="请输入组件路径" />
        </el-form-item>
        <el-form-item label="权限标识" prop="perm">
          <el-input v-model="form.perm" placeholder="请输入权限标识" />
        </el-form-item>
        <el-form-item label="父菜单" prop="parentId">
          <el-select v-model="form.parentId" placeholder="请选择父菜单" style="width: 100%">
            <el-option label="无" :value="0" />
            <el-option
              v-for="menu in menuOptions"
              :key="menu.id"
              :label="menu.name"
              :value="menu.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-input v-model="form.icon" placeholder="请输入图标名称" />
        </el-form-item>
        <el-form-item label="重定向" prop="redirect">
          <el-input v-model="form.redirect" placeholder="请输入重定向路径" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" :max="999" />
        </el-form-item>
        <el-form-item label="可见性" prop="visible">
          <el-radio-group v-model="form.visible">
            <el-radio :label="1">可见</el-radio>
            <el-radio :label="0">隐藏</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="缓存" prop="keepAlive">
          <el-radio-group v-model="form.keepAlive">
            <el-radio :label="1">是</el-radio>
            <el-radio :label="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="总是显示" prop="alwaysShow">
          <el-radio-group v-model="form.alwaysShow">
            <el-radio :label="1">是</el-radio>
            <el-radio :label="0">否</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { 
  getMenuList, 
  createMenu, 
  updateMenu, 
  deleteMenu, 
  getMenuOptions, 
  updateMenuVisibility,
  type Menu, 
  type MenuForm 
} from '@/api/menu'

const loading = ref(false)
const menuList = ref<Menu[]>([])
const menuOptions = ref<Menu[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)
const currentId = ref<number>()

const formRef = ref<FormInstance>()
const form = reactive<MenuForm>({
  parentId: 0,
  name: '',
  routeName: '',
  routePath: '',
  component: '',
  perm: '',
  visible: 1,
  sort: 0,
  icon: '',
  redirect: '',
  keepAlive: 0,
  alwaysShow: 0,
  params: []
})

const rules: FormRules = {
  name: [
    { required: true, message: '请输入菜单名称', trigger: 'blur' }
  ],
  routeName: [
    { required: true, message: '请输入路由名称', trigger: 'blur' }
  ],
  routePath: [
    { required: true, message: '请输入路由路径', trigger: 'blur' }
  ],
  component: [
    { required: true, message: '请输入组件路径', trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: '请选择父菜单', trigger: 'change' }
  ],
  visible: [
    { required: true, message: '请选择可见性', trigger: 'change' }
  ],
  sort: [
    { required: true, message: '请输入排序', trigger: 'blur' }
  ]
}

// 获取菜单列表
const fetchMenuList = async () => {
  loading.value = true
  try {
    const data = await getMenuList()
    menuList.value = data
  } catch (error) {
    console.error('获取菜单列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取菜单选项
const fetchMenuOptions = async () => {
  try {
    const data = await getMenuOptions()
    menuOptions.value = data
  } catch (error) {
    console.error('获取菜单选项失败:', error)
  }
}

// 添加菜单
const handleAdd = () => {
  dialogTitle.value = '添加菜单'
  isEdit.value = false
  dialogVisible.value = true
}

// 编辑菜单
const handleEdit = (row: Menu) => {
  dialogTitle.value = '编辑菜单'
  isEdit.value = true
  currentId.value = row.id
  Object.assign(form, row)
  dialogVisible.value = true
}

// 切换可见性
const handleToggleVisibility = async (row: Menu) => {
  try {
    const newVisible = row.visible === 1 ? 0 : 1
    await updateMenuVisibility(row.id, newVisible)
    ElMessage.success('更新成功')
    fetchMenuList()
  } catch (error) {
    console.error('更新可见性失败:', error)
  }
}

// 删除菜单
const handleDelete = async (row: Menu) => {
  try {
    await ElMessageBox.confirm('确定要删除这个菜单吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await deleteMenu(row.id)
    ElMessage.success('删除成功')
    fetchMenuList()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除菜单失败:', error)
    }
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    if (isEdit.value && currentId.value) {
      await updateMenu(currentId.value, form)
      ElMessage.success('编辑成功')
    } else {
      await createMenu(form)
      ElMessage.success('添加成功')
    }
    
    dialogVisible.value = false
    fetchMenuList()
  } catch (error) {
    console.error('提交失败:', error)
  }
}

// 重置表单
const resetForm = () => {
  if (formRef.value) {
    formRef.value.resetFields()
  }
  Object.assign(form, {
    parentId: 0,
    name: '',
    routeName: '',
    routePath: '',
    component: '',
    perm: '',
    visible: 1,
    sort: 0,
    icon: '',
    redirect: '',
    keepAlive: 0,
    alwaysShow: 0,
    params: []
  })
  isEdit.value = false
  currentId.value = undefined
}

onMounted(() => {
  fetchMenuList()
  fetchMenuOptions()
})
</script>

<style scoped>
.menu-container {
  padding: 20px;
  height: 100%;
  width: 100%;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style> 