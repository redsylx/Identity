import { create } from 'zustand'
import { api } from '../services/api'
import { withAsyncOperation } from '../utils/async'

interface AppState {
  randomNumber: number | null
  randomNumberLoading: boolean
  error: string | null

  fetchRandomNumber: () => Promise<void>
  clearError: () => void
}

export const useAppStore = create<AppState>((set) => ({
  randomNumber: null,
  randomNumberLoading: false,
  error: null,

  fetchRandomNumber: async () => {
    await withAsyncOperation(
      () => api.getRandomNumber(),
      {
        setLoading: (loading) => set({ randomNumberLoading: loading }),
        setError: (error) => set({ error }),
        setData: (number) => set({ randomNumber: number }),
      }
    )
  },

  clearError: () => set({ error: null }),
}))
