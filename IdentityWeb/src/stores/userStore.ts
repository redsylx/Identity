import { create } from 'zustand'
import type { User, CreateUserRequest } from '../types'
import { api } from '../services/api'
import { withAsyncOperation } from '../utils/async'

interface UserState {
  users: User[]
  loading: boolean
  error: string | null

  fetchUsers: () => Promise<void>
  createUser: (userData: CreateUserRequest) => Promise<void>
  deleteUser: (id: number) => Promise<void>
  clearError: () => void
}

export const useUserStore = create<UserState>((set) => ({
  users: [],
  loading: false,
  error: null,

  fetchUsers: async () => {
    await withAsyncOperation(
      () => api.getUsers(),
      {
        setLoading: (loading) => set({ loading }),
        setError: (error) => set({ error }),
        setData: (users) => set({ users }),
      }
    )
  },

  createUser: async (userData: CreateUserRequest) => {
    await withAsyncOperation(
      () => api.createUser(userData),
      {
        setLoading: (loading) => set({ loading }),
        setError: (error) => set({ error }),
        setData: (newUser) =>
          set((state) => ({
            users: [...state.users, newUser],
          })),
        rethrow: true,
      }
    )
  },

  deleteUser: async (id: number) => {
    await withAsyncOperation(
      () => api.deleteUser(id),
      {
        setLoading: (loading) => set({ loading }),
        setError: (error) => set({ error }),
        setData: () =>
          set((state) => ({
            users: state.users.filter((user) => user.id !== id),
          })),
        rethrow: true,
      }
    )
  },

  clearError: () => set({ error: null }),
}))
