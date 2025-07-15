import request from '@/utils/request'

export interface KeyValue {
  key: string
  value: string
}

export interface Menu {
  id: number
  parentId: number
  name: string
  routeName: string
  routePath: string
  component: string
  perm: string
  visible: number
  sort: number
  icon: string
  redirect: string
  keepAlive: number
  alwaysShow: number
  params: KeyValue[]
  updateAt?: string
  createAt?: string
}

export interface MenuForm {
  parentId: number
  name: string
  routeName: string
  routePath: string
  component: string
  perm: string
  visible: number
  sort: number
  icon: string
  redirect: string
  keepAlive: number
  alwaysShow: number
  params: KeyValue[]
}

// 获取菜单列表
export function getMenuList() {
  return request.get<Menu[]>('/api/v1/menus').then(res => res.data)
}

// 创建菜单
export function createMenu(data: MenuForm) {
  return request.post('/api/v1/menus', data).then(res => res.data)
}

// 更新菜单
export function updateMenu(id: number, data: MenuForm) {
  return request.put(`/api/v1/menus/${id}`, data).then(res => res.data)
}

// 删除菜单
export function deleteMenu(id: number) {
  return request.delete(`/api/v1/menus/${id}`).then(res => res.data)
}

// 获取菜单详情
export function getMenuDetail(id: number) {
  return request.get<Menu>(`/api/v1/menus/${id}`).then(res => res.data)
}

// 获取菜单表单数据
export function getMenuFormData(id: number) {
  return request.get<Menu>(`/api/v1/menus/${id}/form`).then(res => res.data)
}

// 更新菜单可见性
export function updateMenuVisibility(menuId: number, visible: number) {
  return request.patch(`/api/v1/menus/${menuId}`, { visible }).then(res => res.data)
}

// 获取菜单选项
export function getMenuOptions() {
  return request.get<Menu[]>('/api/v1/menus/options').then(res => res.data)
}

// 获取菜单路由
export function getMenuRoutes() {
  return request.get<Menu[]>('/api/v1/menus/routes').then(res => res.data)
} 