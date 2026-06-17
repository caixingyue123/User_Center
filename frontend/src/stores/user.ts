import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from '@/types/user'
import * as userApi from '@/api/user'
import { setToken, removeToken, getToken } from '@/utils/token'

export const useUserStore = defineStore('user', () => {
  const userInfo = ref<User | null>(null)
  const token = ref<string | null>(getToken())

  const isLoggedIn = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const res = await userApi.login({ username, password })
    token.value = res.token
    userInfo.value = res.user
    setToken(res.token)
  }

  async function register(
    username: string,
    password: string,
    nickname?: string,
    email?: string,
    phone?: string,
  ) {
    await userApi.register({ username, password, nickname, email, phone })
  }

  async function fetchProfile() {
    const user = await userApi.getProfile()
    userInfo.value = user
  }

  async function updateProfile(data: {
    nickname?: string
    email?: string
    phone?: string
    avatar?: string
  }) {
    const user = await userApi.updateProfile(data)
    userInfo.value = user
  }

  function logout() {
    token.value = null
    userInfo.value = null
    removeToken()
  }

  return {
    userInfo,
    token,
    isLoggedIn,
    login,
    register,
    fetchProfile,
    updateProfile,
    logout,
  }
})
