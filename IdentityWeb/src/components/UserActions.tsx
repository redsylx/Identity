import { getErrorMessage } from '../utils/errors'
import type { CreateUserRequest } from '../types'

export async function handleCreateUser(
  userData: CreateUserRequest,
  createUser: (userData: CreateUserRequest) => Promise<void>
): Promise<{ success: boolean; error?: string }> {
  try {
    await createUser(userData)
    return { success: true }
  } catch (error) {
    return { success: false, error: getErrorMessage(error) }
  }
}

export async function handleDeleteUser(
  id: number,
  deleteUser: (id: number) => Promise<void>
): Promise<{ success: boolean; error?: string }> {
  if (window.confirm('Are you sure you want to delete this user?')) {
    try {
      await deleteUser(id)
      return { success: true }
    } catch (error) {
      return { success: false, error: getErrorMessage(error) }
    }
  }
  return { success: false, error: 'Deletion cancelled' }
}
