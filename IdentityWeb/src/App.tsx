import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [randomNumber, setRandomNumber] = useState<number | null>(null)
  const [loading, setLoading] = useState(false)

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
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={fetchRandomNumber} disabled={loading}>
          {loading ? 'Loading...' : 'Get Random Number'}
        </button>
        {randomNumber !== null && (
          <p style={{ marginTop: '1rem', fontSize: '1.5rem' }}>
            Random number: {randomNumber}
          </p>
        )}
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
