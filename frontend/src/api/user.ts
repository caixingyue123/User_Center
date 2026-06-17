import request from './request'
import type {
  User,
  LoginReq,
  RegisterReq,
  UpdateProfileReq,
  ListUsersParams,
  ListUsersRes,
  LoginRes,
} from '@/types/user'

export function login(data: LoginReq): Promise<LoginRes> {
  return request.post('/login', data)
}

export function register(data: RegisterReq): Promise<User> {
  return request.post('/register', data)
}

export function listUsers(params: ListUsersParams): Promise<ListUsersRes> {
  return request.get('', { params })
}

export function getProfile(): Promise<User> {
  return request.get('/profile')
}

export function updateProfile(data: UpdateProfileReq): Promise<User> {
  return request.put('/profile', data)
}

export function deleteUser(id: number): Promise<null> {
  return request.delete(`/${id}`)
}
