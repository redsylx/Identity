/**
 * API error response interface
 * Provides type safety for error handling across the application
 */
export interface ApiErrorResponse {
  error?: string
  message?: string
  details?: unknown
}

/**
 * Determines if a response is an API error response
 */
export function isApiErrorResponse(data: unknown): data is ApiErrorResponse {
  return (
    typeof data === 'object' &&
    data !== null &&
    ('error' in data || 'message' in data)
  )
}
