import { useState } from 'react'
import type { CreateUserRequest } from '../types'
import { validateName, validateEmail, ValidationError } from '../utils/validation'

interface UserFormProps {
  onSubmit: (userData: CreateUserRequest) => Promise<{ success: boolean; error?: string }>
  loading?: boolean
}

interface FormErrors {
  name?: string
  email?: string
}

export function UserForm({ onSubmit, loading }: UserFormProps) {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [errors, setErrors] = useState<FormErrors>({})
  const [submitError, setSubmitError] = useState<string | null>(null)

  const validateForm = (): boolean => {
    const newErrors: FormErrors = {}

    try {
      validateName(name)
    } catch (error) {
      if (error instanceof ValidationError) {
        newErrors.name = error.message
      }
    }

    try {
      validateEmail(email)
    } catch (error) {
      if (error instanceof ValidationError) {
        newErrors.email = error.message
      }
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitError(null)

    if (!validateForm()) {
      return
    }

    const trimmedName = name.trim()
    const trimmedEmail = email.trim()

    const result = await onSubmit({ name: trimmedName, email: trimmedEmail })
    if (result.success) {
      setName('')
      setEmail('')
      setErrors({})
      setSubmitError(null)
    } else {
      setSubmitError(result.error || 'Failed to create user')
    }
  }

  const handleNameChange = (value: string) => {
    setName(value)
    // Clear error as user types
    if (errors.name) {
      setErrors((prev) => ({ ...prev, name: undefined }))
    }
    if (submitError) {
      setSubmitError(null)
    }
  }

  const handleEmailChange = (value: string) => {
    setEmail(value)
    // Clear error as user types
    if (errors.email) {
      setErrors((prev) => ({ ...prev, email: undefined }))
    }
    if (submitError) {
      setSubmitError(null)
    }
  }

  const isFormValid = name.trim().length > 0 && email.trim().length > 0 && Object.keys(errors).length === 0

  return (
    <form onSubmit={handleSubmit} className="space-y-4" noValidate>
      <h3 className="text-lg font-medium text-gray-900">Create New User</h3>

      {submitError && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-3 py-2 rounded text-sm">
          {submitError}
        </div>
      )}

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
        <div>
          <label htmlFor="name" className="block text-sm font-medium text-gray-700">
            Name <span className="text-red-500">*</span>
          </label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => handleNameChange(e.target.value)}
            className={`mt-1 block w-full rounded-md shadow-sm px-3 py-2 border focus:outline-none focus:ring-2 sm:text-sm ${
              errors.name
                ? 'border-red-300 focus:border-red-500 focus:ring-red-500'
                : 'border-gray-300 focus:border-blue-500 focus:ring-blue-500'
            }`}
            placeholder="John Doe"
            aria-invalid={!!errors.name}
            aria-describedby={errors.name ? 'name-error' : undefined}
          />
          {errors.name && (
            <p id="name-error" className="mt-1 text-sm text-red-600">
              {errors.name}
            </p>
          )}
        </div>

        <div>
          <label htmlFor="email" className="block text-sm font-medium text-gray-700">
            Email <span className="text-red-500">*</span>
          </label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => handleEmailChange(e.target.value)}
            className={`mt-1 block w-full rounded-md shadow-sm px-3 py-2 border focus:outline-none focus:ring-2 sm:text-sm ${
              errors.email
                ? 'border-red-300 focus:border-red-500 focus:ring-red-500'
                : 'border-gray-300 focus:border-blue-500 focus:ring-blue-500'
            }`}
            placeholder="john@example.com"
            aria-invalid={!!errors.email}
            aria-describedby={errors.email ? 'email-error' : undefined}
          />
          {errors.email && (
            <p id="email-error" className="mt-1 text-sm text-red-600">
              {errors.email}
            </p>
          )}
        </div>
      </div>

      <button
        type="submit"
        disabled={loading || !isFormValid}
        className="w-full sm:w-auto sm:px-6 bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors font-medium"
      >
        {loading ? 'Creating...' : 'Create User'}
      </button>
    </form>
  )
}
