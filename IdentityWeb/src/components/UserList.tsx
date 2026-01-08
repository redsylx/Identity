import type { User } from '../types'

interface UserListProps {
  users: User[]
  onDelete?: (id: number) => void
  loading?: boolean
}

export function UserList({ users, onDelete, loading }: UserListProps) {
  if (loading) {
    return (
      <div className="flex justify-center items-center py-8">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
      </div>
    )
  }

  if (users.length === 0) {
    return (
      <div className="text-center py-8 text-gray-500">
        No users found. Create one to get started!
      </div>
    )
  }

  return (
    <div className="bg-white shadow rounded-lg overflow-hidden">
      <ul className="divide-y divide-gray-200">
        {users.map((user) => (
          <li
            key={user.id}
            className="px-6 py-4 flex items-center justify-between hover:bg-gray-50"
          >
            <div>
              <p className="text-sm font-medium text-gray-900">{user.name}</p>
              <p className="text-sm text-gray-500">{user.email}</p>
            </div>
            {onDelete && (
              <button
                onClick={() => onDelete(user.id)}
                className="text-red-600 hover:text-red-800 text-sm font-medium px-3 py-1 rounded border border-red-300 hover:border-red-500 transition-colors"
              >
                Delete
              </button>
            )}
          </li>
        ))}
      </ul>
    </div>
  )
}
