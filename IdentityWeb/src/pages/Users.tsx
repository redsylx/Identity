import { useEffect, useRef } from 'react'
import { useUserStore } from '../stores/userStore'
import { UserList } from '../components/UserList'
import { UserForm } from '../components/UserForm'
import { UserManagementHeader } from '../components/UserManagementHeader'
import { useUserActions } from '../hooks/useUserActions'

export function Users() {
  const { users, loading, fetchUsers } = useUserStore()
  const { handleCreate, handleDelete } = useUserActions()
  const isMounted = useRef(true)

  useEffect(() => {
    fetchUsers()

    return () => {
      isMounted.current = false
    }
  }, [fetchUsers])

  const handleCreateUser = async (userData: { name: string; email: string }) => {
    return handleCreate(userData)
  }

  const handleDeleteUser = async (id: number) => {
    if (!isMounted.current) return

    const result = await handleDelete(id)
    if (!result.success && result.error !== 'Deletion cancelled') {
      // Error is handled by the store
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-3xl font-bold text-gray-900">User Management</h2>
        <p className="mt-2 text-gray-600">View, create, and manage users in the system.</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div className="lg:col-span-1">
          <div className="bg-white shadow rounded-lg p-6">
            <UserForm onSubmit={handleCreateUser} loading={loading} />
          </div>
        </div>

        <div className="lg:col-span-1">
          <div className="bg-white shadow rounded-lg p-6">
            <UserManagementHeader userCount={users.length} onRefresh={fetchUsers} loading={loading} />
            <UserList users={users} onDelete={handleDeleteUser} loading={loading} />
          </div>
        </div>
      </div>
    </div>
  )
}
