import type { User, CreateUserRequest } from '../types'
import type { ApiErrorResponse } from '../types/api'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export const api = {
  async getRandomNumber(): Promise<number> {
    const response = await fetch(`${API_BASE_URL}/random`)
    if (!response.ok) {
      throw new Error('Failed to fetch random number')
    }
    const data = await response.json()
    return data.number
  },

  async getUsers(): Promise<User[]> {
    const response = await fetch(`${API_BASE_URL}/api/users`)
    if (!response.ok) {
      throw new Error('Failed to fetch users')
    }
    const data = await response.json()
    return data || []
  },

  async createUser(userData: CreateUserRequest): Promise<User> {
    const response = await fetch(`${API_BASE_URL}/api/users/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    })

    if (!response.ok) {
      const errorData = await response.json().catch<ApiErrorResponse | null>(() => null)
      const errorMessage = errorData?.error || errorData?.message || 'Failed to create user'
      throw new Error(errorMessage)
    }

    return response.json()
  },

  async deleteUser(id: number): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/api/users/${id}`, {
      method: 'DELETE',
    })

    if (!response.ok) {
      throw new Error('Failed to delete user')
    }
  },
}
