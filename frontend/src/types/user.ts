export interface User {
  id: number
  username: string
  nickname: string
  email: string
  phone: string
  avatar: string
  status: number
  created_at: string
  updated_at: string
}

export interface LoginReq {
  username: string
  password: string
}

export interface RegisterReq {
  username: string
  password: string
  nickname?: string
  email?: string
  phone?: string
}

export interface UpdateProfileReq {
  nickname?: string
  email?: string
  phone?: string
  avatar?: string
}

export interface ListUsersParams {
  page: number
  page_size: number
}

export interface ListUsersRes {
  list: User[]
  total: number
  page: number
  page_size: number
}

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

export interface LoginRes {
  user: User
  token: string
}
