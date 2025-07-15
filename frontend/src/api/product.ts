import request from '@/utils/request'

export interface Price {
  terminal: number
  wholesale: number
  cost: number
  srp: number
  warning: number
}

export interface Product {
  id: string
  name: string
  price: Price
  specification: number
  isActive: boolean
  image: string
  tag: string
  brandId: string
  nickName: string
  barCode: string
  updateAt?: string
  createAt?: string
}

export interface ProductForm {
  name: string
  price: Price
  specification: number
  isActive: boolean
  image: string
  tag: string
  brandId: string
  nickName: string
  barCode: string
}

// 添加商品
export function addProduct(data: ProductForm) {
  return request.post('/product', data).then(res => res.data)
}

// 更新商品
export function updateProduct(data: ProductForm) {
  return request.put('/product', data).then(res => res.data)
}

// 获取商品列表
export function getProductList() {
  return request.get<Product[]>('/product').then(res => res.data)
} 