<template>
  <div class="order-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>订单管理</span>
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            添加订单
          </el-button>
        </div>
      </template>

      <!-- 订单列表 -->
      <el-table :data="orderList" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="订单ID" width="120" />
        <el-table-column prop="shopName" label="商店名称" width="150" />
        <el-table-column prop="address" label="收货地址" />
        <el-table-column prop="state" label="配送状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStateType(row.state)">
              {{ getStateText(row.state) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="payment" label="支付状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getPaymentType(row.payment)">
              {{ getPaymentText(row.payment) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="purchaserId" label="下单人" width="120" />
        <el-table-column prop="driverId" label="配送司机" width="120" />
        <el-table-column prop="createAt" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="info" @click="handleView(row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 添加/编辑订单对话框 -->
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
        <el-form-item label="商店ID" prop="shopId">
          <el-input v-model="form.shopId" placeholder="请输入商店ID" />
        </el-form-item>
        <el-form-item label="商店名称" prop="shopName">
          <el-input v-model="form.shopName" placeholder="请输入商店名称" />
        </el-form-item>
        <el-form-item label="收货地址" prop="address">
          <el-input v-model="form.address" type="textarea" placeholder="请输入收货地址" />
        </el-form-item>
        <el-form-item label="配送状态" prop="state">
          <el-select v-model="form.state" placeholder="请选择配送状态" style="width: 100%">
            <el-option label="待配送" :value="0" />
            <el-option label="配送中" :value="1" />
            <el-option label="已送达" :value="2" />
            <el-option label="已取消" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="支付状态" prop="payment">
          <el-select v-model="form.payment" placeholder="请选择支付状态" style="width: 100%">
            <el-option label="未支付" :value="0" />
            <el-option label="已支付" :value="1" />
            <el-option label="已退款" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="下单人" prop="purchaserId">
          <el-input v-model="form.purchaserId" placeholder="请输入下单人ID" />
        </el-form-item>
        <el-form-item label="配送司机" prop="driverId">
          <el-input v-model="form.driverId" placeholder="请输入配送司机ID" />
        </el-form-item>
        <el-form-item label="回单照片" prop="picture">
          <el-input v-model="form.picture" placeholder="请输入照片URL" />
        </el-form-item>
        
        <!-- 商品列表 -->
        <el-form-item label="商品列表" prop="products">
          <div class="products-container">
            <div v-for="(product, index) in form.products" :key="index" class="product-item">
              <el-row :gutter="10">
                <el-col :span="6">
                  <el-input v-model="product.productId" placeholder="商品ID" />
                </el-col>
                <el-col :span="6">
                  <el-input v-model="product.productName" placeholder="商品名称" />
                </el-col>
                <el-col :span="4">
                  <el-input-number v-model="product.price" :precision="2" placeholder="价格" />
                </el-col>
                <el-col :span="4">
                  <el-input-number v-model="product.count" :min="1" placeholder="数量" />
                </el-col>
                <el-col :span="4">
                  <el-button type="danger" size="small" @click="removeProduct(index)">删除</el-button>
                </el-col>
              </el-row>
            </div>
            <el-button type="primary" size="small" @click="addProduct">添加商品</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 查看订单详情对话框 -->
    <el-dialog
      v-model="viewDialogVisible"
      title="订单详情"
      width="600px"
    >
      <div v-if="currentOrder">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="订单ID">{{ currentOrder.id }}</el-descriptions-item>
          <el-descriptions-item label="商店名称">{{ currentOrder.shopName }}</el-descriptions-item>
          <el-descriptions-item label="收货地址">{{ currentOrder.address }}</el-descriptions-item>
          <el-descriptions-item label="配送状态">
            <el-tag :type="getStateType(currentOrder.state)">
              {{ getStateText(currentOrder.state) }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="支付状态">
            <el-tag :type="getPaymentType(currentOrder.payment)">
              {{ getPaymentText(currentOrder.payment) }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="下单人">{{ currentOrder.purchaserId }}</el-descriptions-item>
          <el-descriptions-item label="配送司机">{{ currentOrder.driverId }}</el-descriptions-item>
          <el-descriptions-item label="创建时间">{{ currentOrder.createAt }}</el-descriptions-item>
        </el-descriptions>
        
        <div class="products-detail">
          <h4>商品列表</h4>
          <el-table :data="currentOrder.products" border>
            <el-table-column prop="productId" label="商品ID" width="120" />
            <el-table-column prop="productName" label="商品名称" />
            <el-table-column prop="price" label="价格" width="100" />
            <el-table-column prop="count" label="数量" width="80" />
            <el-table-column label="小计" width="100">
              <template #default="{ row }">
                {{ (row.price * row.count).toFixed(2) }}
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { getOrderList, addOrder, updateOrder, type Order, type OrderForm, type Products } from '@/api/order'

const loading = ref(false)
const orderList = ref<Order[]>([])
const dialogVisible = ref(false)
const viewDialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)
const currentOrder = ref<Order | null>(null)

const formRef = ref<FormInstance>()
const form = reactive<OrderForm>({
  products: [],
  shopId: '',
  shopName: '',
  address: '',
  state: 0,
  payment: 0,
  purchaserId: '',
  driverId: '',
  picture: ''
})

const rules: FormRules = {
  shopId: [
    { required: true, message: '请输入商店ID', trigger: 'blur' }
  ],
  shopName: [
    { required: true, message: '请输入商店名称', trigger: 'blur' }
  ],
  address: [
    { required: true, message: '请输入收货地址', trigger: 'blur' }
  ],
  state: [
    { required: true, message: '请选择配送状态', trigger: 'change' }
  ],
  payment: [
    { required: true, message: '请选择支付状态', trigger: 'change' }
  ],
  purchaserId: [
    { required: true, message: '请输入下单人ID', trigger: 'blur' }
  ]
}

// 获取订单列表
const fetchOrderList = async () => {
  loading.value = true
  try {
    const data = await getOrderList()
    orderList.value = data
  } catch (error) {
    console.error('获取订单列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 添加订单
const handleAdd = () => {
  dialogTitle.value = '添加订单'
  isEdit.value = false
  dialogVisible.value = true
}

// 编辑订单
const handleEdit = (row: Order) => {
  dialogTitle.value = '编辑订单'
  isEdit.value = true
  Object.assign(form, row)
  dialogVisible.value = true
}

// 查看订单
const handleView = (row: Order) => {
  currentOrder.value = row
  viewDialogVisible.value = true
}

// 添加商品
const addProduct = () => {
  form.products.push({
    productId: '',
    productName: '',
    price: 0,
    count: 1
  })
}

// 删除商品
const removeProduct = (index: number) => {
  form.products.splice(index, 1)
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    if (isEdit.value) {
      await updateOrder(form)
      ElMessage.success('编辑成功')
    } else {
      await addOrder(form)
      ElMessage.success('添加成功')
    }
    
    dialogVisible.value = false
    fetchOrderList()
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
    products: [],
    shopId: '',
    shopName: '',
    address: '',
    state: 0,
    payment: 0,
    purchaserId: '',
    driverId: '',
    picture: ''
  })
  isEdit.value = false
}

// 获取状态类型
const getStateType = (state: number) => {
  const types = ['info', 'warning', 'success', 'danger']
  return types[state] || 'info'
}

// 获取状态文本
const getStateText = (state: number) => {
  const texts = ['待配送', '配送中', '已送达', '已取消']
  return texts[state] || '未知'
}

// 获取支付类型
const getPaymentType = (payment: number) => {
  const types = ['danger', 'success', 'info']
  return types[payment] || 'info'
}

// 获取支付文本
const getPaymentText = (payment: number) => {
  const texts = ['未支付', '已支付', '已退款']
  return texts[payment] || '未知'
}

onMounted(() => {
  fetchOrderList()
})
</script>

<style scoped>
.order-container {
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

.products-container {
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  padding: 10px;
}

.product-item {
  margin-bottom: 10px;
  padding: 10px;
  border: 1px solid #f0f0f0;
  border-radius: 4px;
  background-color: #fafafa;
}

.products-detail {
  margin-top: 20px;
}

.products-detail h4 {
  margin-bottom: 10px;
  color: #606266;
}
</style> 