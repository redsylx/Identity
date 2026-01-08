import { getErrorMessage } from './errors'

/**
 * Async operation state
 */
export interface AsyncOperationState<T> {
  data: T | null
  loading: boolean
  error: string | null
}

/**
 * Async operation configuration
 */
export interface AsyncOperationConfig<T> {
  setLoading: (loading: boolean) => void
  setError: (error: string | null) => void
  setData?: (data: T) => void
  rethrow?: boolean
}

/**
 * Wrapper for async operations with consistent loading and error handling
 * Reduces code duplication across stores
 *
 * @example
 * ```ts
 * await withAsyncOperation(
 *   async () => api.getUsers(),
 *   {
 *     setLoading: (loading) => set({ loading }),
 *     setError: (error) => set({ error }),
 *     setData: (users) => set({ users })
 *   }
 * )
 * ```
 */
export async function withAsyncOperation<T>(
  operation: () => Promise<T>,
  config: AsyncOperationConfig<T>
): Promise<void> {
  const { setLoading, setError, setData, rethrow = false } = config

  setLoading(true)
  setError(null)

  try {
    const result = await operation()
    if (setData) {
      setData(result)
    }
  } catch (error) {
    setError(getErrorMessage(error))
    if (rethrow) {
      throw error
    }
  } finally {
    setLoading(false)
  }
}
