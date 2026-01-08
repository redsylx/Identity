import { useCallback } from 'react'
import { useUserStore } from '../stores/userStore'
import type { CreateUserRequest } from '../types'

/**
 * Custom hook for user action business logic
 * Encapsulates user creation and deletion logic with proper error handling
 */
export function useUserActions() {
  const { createUser, deleteUser } = useUserStore()

  const handleCreate = useCallback(
    async (userData: CreateUserRequest): Promise<{ success: boolean; error?: string }> => {
      try {
        await createUser(userData)
        return { success: true }
      } catch (error) {
        return {
          success: false,
          error: error instanceof Error ? error.message : 'Failed to create user',
        }
      }
    },
    [createUser]
  )

  const handleDelete = useCallback(
    async (id: number): Promise<{ success: boolean; error?: string }> => {
      if (!window.confirm('Are you sure you want to delete this user?')) {
        return { success: false, error: 'Deletion cancelled' }
      }

      try {
        await deleteUser(id)
        return { success: true }
      } catch (error) {
        return {
          success: false,
          error: error instanceof Error ? error.message : 'Failed to delete user',
        }
      }
    },
    [deleteUser]
  )

  return { handleCreate, handleDelete }
}
