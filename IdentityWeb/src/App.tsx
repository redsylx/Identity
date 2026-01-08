import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

interface User {
  id: number
  name: string
  email: string
}

function App() {
  const [randomNumber, setRandomNumber] = useState<number | null>(null)
  const [loading, setLoading] = useState(false)
  const [users, setUsers] = useState<User[]>([])
  const [loadingUsers, setLoadingUsers] = useState(false)
  const [newUserName, setNewUserName] = useState('')
  const [newUserEmail, setNewUserEmail] = useState('')
  const [creatingUser, setCreatingUser] = useState(false)

  const fetchUsers = async () => {
    setLoadingUsers(true)
    try {
      const response = await fetch('http://localhost:8080/api/users')
      if (!response.ok) {
        throw new Error('Failed to fetch users')
      }
      const data = await response.json()
      setUsers(data || [])
    } catch (error) {
      console.error('Failed to fetch users:', error)
      setUsers([])
    } finally {
      setLoadingUsers(false)
    }
  }

  // Fetch users on component mount
  useEffect(() => {
    fetchUsers()
  }, [])

  const fetchRandomNumber = async () => {
    setLoading(true)
    try {
      const response = await fetch('http://localhost:8080/random')
      const data = await response.json()
      setRandomNumber(data.number)
    } catch (error) {
      console.error('Failed to fetch random number:', error)
    } finally {
      setLoading(false)
    }
  }

  const createUser = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newUserName || !newUserEmail) return

    setCreatingUser(true)
    try {
      const response = await fetch('http://localhost:8080/api/users/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: newUserName,
          email: newUserEmail,
        }),
      })

      if (response.ok) {
        const createdUser = await response.json()
        setUsers([...users, createdUser])
        setNewUserName('')
        setNewUserEmail('')
      }
    } catch (error) {
      console.error('Failed to create user:', error)
    } finally {
      setCreatingUser(false)
    }
  }

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Identity Management</h1>

      {/* Random Number Section */}
      <div className="card">
        <button onClick={fetchRandomNumber} disabled={loading}>
          {loading ? 'Loading...' : 'Get Random Number'}
        </button>
        {randomNumber !== null && (
          <p style={{ marginTop: '1rem', fontSize: '1.5rem' }}>
            Random number: {randomNumber}
          </p>
        )}
      </div>

      {/* Users Section */}
      <div className="card" style={{ marginTop: '2rem' }}>
        <h2>Users</h2>
        <button onClick={fetchUsers} disabled={loadingUsers} style={{ marginBottom: '1rem' }}>
          {loadingUsers ? 'Loading...' : 'Refresh Users'}
        </button>

        {users && users.length > 0 && (
          <div style={{ textAlign: 'left', marginTop: '1rem' }}>
            {users.map((user) => (
              <div key={user.id} style={{ padding: '0.5rem', borderBottom: '1px solid #eee' }}>
                <strong>{user.name}</strong> ({user.email})
              </div>
            ))}
          </div>
        )}

        {/* Create User Form */}
        <form onSubmit={createUser} style={{ marginTop: '1.5rem', textAlign: 'left' }}>
          <h3>Create New User</h3>
          <div style={{ marginBottom: '0.5rem' }}>
            <input
              type="text"
              placeholder="Name"
              value={newUserName}
              onChange={(e) => setNewUserName(e.target.value)}
              style={{ padding: '0.5rem', marginRight: '0.5rem', width: '200px' }}
              required
            />
            <input
              type="email"
              placeholder="Email"
              value={newUserEmail}
              onChange={(e) => setNewUserEmail(e.target.value)}
              style={{ padding: '0.5rem', width: '200px' }}
              required
            />
          </div>
          <button type="submit" disabled={creatingUser || !newUserName || !newUserEmail}>
            {creatingUser ? 'Creating...' : 'Create User'}
          </button>
        </form>
      </div>

      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
