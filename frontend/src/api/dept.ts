import request from '@/utils/request'

export interface Dept {
  id: number
  code: string
  name: string
  parentId: number
  status: number
  sort: number
  updateAt?: string
  createAt?: string
}

export interface DeptForm {
  code: string
  name: string
  parentId: number
  status: number
  sort: number
}

// 获取部门列表
export function getDeptList() {
  return request.get<Dept[]>('/dept').then(res => res.data)
}

// 添加部门
export function addDept(data: DeptForm) {
  return request.post('/dept', data).then(res => res.data)
}

// 编辑部门
export function editDept(deptId: number, data: DeptForm) {
  return request.put(`/dept/${deptId}`, data).then(res => res.data)
}

// 删除部门
export function deleteDept(ids: string) {
  return request.delete(`/dept/${ids}`).then(res => res.data)
}

// 获取部门表单数据
export function getDeptForm(deptId: number) {
  return request.get<Dept>(`/dept/${deptId}/form`).then(res => res.data)
}

// 获取部门选项
export function getDeptOptions() {
  return request.get<Dept[]>('/dept/options').then(res => res.data)
} 