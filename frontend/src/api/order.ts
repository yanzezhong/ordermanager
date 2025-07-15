import request from '@/utils/request'

export interface Products {
  productId: string
  productName: string
  price: number
  count: number
}

export interface Order {
  id: string
  products: Products[]
  shopId: string
  shopName: string
  address: string
  state: number
  payment: number
  purchaserId: string
  driverId: string
  picture: string
  updateAt?: string
  createAt?: string
}

export interface OrderForm {
  products: Products[]
  shopId: string
  shopName: string
  address: string
  state: number
  payment: number
  purchaserId: string
  driverId: string
  picture: string
}

// 添加订单
export function addOrder(data: OrderForm) {
  return request.post('/order', data).then(res => res.data)
}

// 获取订单列表
export function getOrderList() {
  return request.get<Order[]>('/order').then(res => res.data)
}

// 更新订单
export function updateOrder(data: OrderForm) {
  return request.put('/order', data).then(res => res.data)
} 