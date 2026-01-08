interface UserManagementHeaderProps {
  userCount: number
  onRefresh: () => void
  loading: boolean
}

export function UserManagementHeader({ userCount, onRefresh, loading }: UserManagementHeaderProps) {
  return (
    <div className="flex justify-between items-center mb-4">
      <h3 className="text-lg font-medium text-gray-900">
        All Users ({userCount})
      </h3>
      <button
        onClick={onRefresh}
        disabled={loading}
        className="text-blue-600 hover:text-blue-800 disabled:text-gray-400 font-medium text-sm"
      >
        {loading ? 'Refreshing...' : 'Refresh'}
      </button>
    </div>
  )
}
