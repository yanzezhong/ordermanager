<template>
  <div class="dept-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>部门管理</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            添加部门
          </el-button>
        </div>
      </template>

      <!-- 部门列表 -->
      <el-table :data="deptList" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="code" label="部门编码" width="120" />
        <el-table-column prop="name" label="部门名称" />
        <el-table-column prop="parentId" label="父部门ID" width="100" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="排序" width="80" />
        <el-table-column prop="createAt" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑部门对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      @close="resetForm"
    >
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="部门编码" prop="code">
          <el-input v-model="form.code" placeholder="请输入部门编码" />
        </el-form-item>
        <el-form-item label="部门名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入部门名称" />
        </el-form-item>
        <el-form-item label="父部门" prop="parentId">
          <el-select v-model="form.parentId" placeholder="请选择父部门" style="width: 100%">
            <el-option label="无" :value="0" />
            <el-option
              v-for="dept in deptOptions"
              :key="dept.id"
              :label="dept.name"
              :value="dept.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" :max="999" />
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
import { getDeptList, addDept, editDept, deleteDept, getDeptOptions, type Dept, type DeptForm } from '@/api/dept'

const loading = ref(false)
const deptList = ref<Dept[]>([])
const deptOptions = ref<Dept[]>([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)
const currentId = ref<number>()

const formRef = ref<FormInstance>()
const form = reactive<DeptForm>({
  code: '',
  name: '',
  parentId: 0,
  status: 1,
  sort: 0
})

const rules: FormRules = {
  code: [
    { required: true, message: '请输入部门编码', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入部门名称', trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: '请选择父部门', trigger: 'change' }
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'change' }
  ],
  sort: [
    { required: true, message: '请输入排序', trigger: 'blur' }
  ]
}

// 获取部门列表
const fetchDeptList = async () => {
  loading.value = true
  try {
    const data = await getDeptList()
    deptList.value = data
  } catch (error) {
    console.error('获取部门列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取部门选项
const fetchDeptOptions = async () => {
  try {
    const data = await getDeptOptions()
    deptOptions.value = data
  } catch (error) {
    console.error('获取部门选项失败:', error)
  }
}

// 添加部门
const handleAdd = () => {
  dialogTitle.value = '添加部门'
  isEdit.value = false
  dialogVisible.value = true
}

// 编辑部门
const handleEdit = (row: Dept) => {
  dialogTitle.value = '编辑部门'
  isEdit.value = true
  currentId.value = row.id
  Object.assign(form, row)
  dialogVisible.value = true
}

// 删除部门
const handleDelete = async (row: Dept) => {
  try {
    await ElMessageBox.confirm('确定要删除这个部门吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    await deleteDept(row.id.toString())
    ElMessage.success('删除成功')
    fetchDeptList()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除部门失败:', error)
    }
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    if (isEdit.value && currentId.value) {
      await editDept(currentId.value, form)
      ElMessage.success('编辑成功')
    } else {
      await addDept(form)
      ElMessage.success('添加成功')
    }
    
    dialogVisible.value = false
    fetchDeptList()
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
    code: '',
    name: '',
    parentId: 0,
    status: 1,
    sort: 0
  })
  isEdit.value = false
  currentId.value = undefined
}

onMounted(() => {
  fetchDeptList()
  fetchDeptOptions()
})
</script>

<style scoped>
.dept-container {
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