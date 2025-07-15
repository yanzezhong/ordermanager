<template>
  <div class="product-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>商品管理</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            添加商品
          </el-button>
        </div>
      </template>

      <!-- 商品列表 -->
      <el-table :data="productList" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="商品ID" width="120" />
        <el-table-column prop="name" label="商品名称" />
        <el-table-column prop="nickName" label="别名" width="120" />
        <el-table-column prop="barCode" label="条码" width="120" />
        <el-table-column prop="specification" label="规格" width="80" />
        <el-table-column prop="isActive" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.isActive ? 'success' : 'danger'">
              {{ row.isActive ? '生效' : '失效' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="tag" label="标签" width="100" />
        <el-table-column prop="brandId" label="品牌ID" width="100" />
        <el-table-column prop="createAt" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="info" @click="handleView(row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑商品对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="800px"
      @close="resetForm"
    >
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="商品名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入商品名称" />
        </el-form-item>
        <el-form-item label="别名" prop="nickName">
          <el-input v-model="form.nickName" placeholder="请输入别名" />
        </el-form-item>
        <el-form-item label="条码" prop="barCode">
          <el-input v-model="form.barCode" placeholder="请输入条码" />
        </el-form-item>
        <el-form-item label="规格" prop="specification">
          <el-input-number v-model="form.specification" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="isActive">
          <el-switch v-model="form.isActive" />
        </el-form-item>
        <el-form-item label="标签" prop="tag">
          <el-input v-model="form.tag" placeholder="请输入标签" />
        </el-form-item>
        <el-form-item label="品牌ID" prop="brandId">
          <el-input v-model="form.brandId" placeholder="请输入品牌ID" />
        </el-form-item>
        <el-form-item label="图片" prop="image">
          <el-input v-model="form.image" placeholder="请输入图片URL" />
        </el-form-item>
        
        <!-- 价格信息 -->
        <el-divider content-position="left">价格信息</el-divider>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="终端零售价" prop="price.terminal">
              <el-input-number v-model="form.price.terminal" :precision="2" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="批发零售价" prop="price.wholesale">
              <el-input-number v-model="form.price.wholesale" :precision="2" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="进价" prop="price.cost">
              <el-input-number v-model="form.price.cost" :precision="2" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="建议零售价" prop="price.srp">
              <el-input-number v-model="form.price.srp" :precision="2" :min="0" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="报警价格" prop="price.warning">
          <el-input-number v-model="form.price.warning" :precision="2" :min="0" style="width: 100%" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 查看商品详情对话框 -->
    <el-dialog
      v-model="viewDialogVisible"
      title="商品详情"
      width="600px"
    >
      <div v-if="currentProduct">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="商品ID">{{ currentProduct.id }}</el-descriptions-item>
          <el-descriptions-item label="商品名称">{{ currentProduct.name }}</el-descriptions-item>
          <el-descriptions-item label="别名">{{ currentProduct.nickName }}</el-descriptions-item>
          <el-descriptions-item label="条码">{{ currentProduct.barCode }}</el-descriptions-item>
          <el-descriptions-item label="规格">{{ currentProduct.specification }}</el-descriptions-item>
          <el-descriptions-item label="状态">
            <el-tag :type="currentProduct.isActive ? 'success' : 'danger'">
              {{ currentProduct.isActive ? '生效' : '失效' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="标签">{{ currentProduct.tag }}</el-descriptions-item>
          <el-descriptions-item label="品牌ID">{{ currentProduct.brandId }}</el-descriptions-item>
          <el-descriptions-item label="创建时间">{{ currentProduct.createAt }}</el-descriptions-item>
        </el-descriptions>
        
        <div class="price-detail">
          <h4>价格信息</h4>
          <el-descriptions :column="2" border>
            <el-descriptions-item label="终端零售价">¥{{ currentProduct.price.terminal }}</el-descriptions-item>
            <el-descriptions-item label="批发零售价">¥{{ currentProduct.price.wholesale }}</el-descriptions-item>
            <el-descriptions-item label="进价">¥{{ currentProduct.price.cost }}</el-descriptions-item>
            <el-descriptions-item label="建议零售价">¥{{ currentProduct.price.srp }}</el-descriptions-item>
            <el-descriptions-item label="报警价格">¥{{ currentProduct.price.warning }}</el-descriptions-item>
          </el-descriptions>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { getProductList, addProduct, updateProduct, type Product, type ProductForm } from '@/api/product'

const loading = ref(false)
const productList = ref<Product[]>([])
const dialogVisible = ref(false)
const viewDialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)
const currentProduct = ref<Product | null>(null)

const formRef = ref<FormInstance>()
const form = reactive<ProductForm>({
  name: '',
  price: {
    terminal: 0,
    wholesale: 0,
    cost: 0,
    srp: 0,
    warning: 0
  },
  specification: 0,
  isActive: true,
  image: '',
  tag: '',
  brandId: '',
  nickName: '',
  barCode: ''
})

const rules: FormRules = {
  name: [
    { required: true, message: '请输入商品名称', trigger: 'blur' }
  ],
  barCode: [
    { required: true, message: '请输入条码', trigger: 'blur' }
  ],
  brandId: [
    { required: true, message: '请输入品牌ID', trigger: 'blur' }
  ],
  'price.terminal': [
    { required: true, message: '请输入终端零售价', trigger: 'blur' }
  ],
  'price.wholesale': [
    { required: true, message: '请输入批发零售价', trigger: 'blur' }
  ],
  'price.cost': [
    { required: true, message: '请输入进价', trigger: 'blur' }
  ]
}

// 获取商品列表
const fetchProductList = async () => {
  loading.value = true
  try {
    const data = await getProductList()
    productList.value = data
  } catch (error) {
    console.error('获取商品列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 添加商品
const handleAdd = () => {
  dialogTitle.value = '添加商品'
  isEdit.value = false
  dialogVisible.value = true
}

// 编辑商品
const handleEdit = (row: Product) => {
  dialogTitle.value = '编辑商品'
  isEdit.value = true
  Object.assign(form, row)
  dialogVisible.value = true
}

// 查看商品
const handleView = (row: Product) => {
  currentProduct.value = row
  viewDialogVisible.value = true
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    if (isEdit.value) {
      await updateProduct(form)
      ElMessage.success('编辑成功')
    } else {
      await addProduct(form)
      ElMessage.success('添加成功')
    }
    
    dialogVisible.value = false
    fetchProductList()
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
    name: '',
    price: {
      terminal: 0,
      wholesale: 0,
      cost: 0,
      srp: 0,
      warning: 0
    },
    specification: 0,
    isActive: true,
    image: '',
    tag: '',
    brandId: '',
    nickName: '',
    barCode: ''
  })
  isEdit.value = false
}

onMounted(() => {
  fetchProductList()
})
</script>

<style scoped>
.product-container {
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

.price-detail {
  margin-top: 20px;
}

.price-detail h4 {
  margin-bottom: 10px;
  color: #606266;
}
</style> 