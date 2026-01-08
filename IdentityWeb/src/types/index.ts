export interface User {
  id: number
  name: string
  email: string
}

export interface CreateUserRequest {
  name: string
  email: string
}

export interface ApiError {
  message: string
}

// Re-export API types for convenience
export type { ApiErrorResponse, isApiErrorResponse } from './api'
