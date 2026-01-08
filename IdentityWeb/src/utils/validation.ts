import { VALIDATION } from '../constants/validation'

/**
 * Email validation regex
 * Supports standard email formats: user@domain.com
 */
const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

/**
 * Validation error class
 */
export class ValidationError extends Error {
  constructor(message: string) {
    super(message)
    this.name = 'ValidationError'
  }
}

/**
 * Validates an email address
 * @throws {ValidationError} if email is invalid
 */
export function validateEmail(email: string): void {
  const trimmed = email.trim()

  if (!trimmed) {
    throw new ValidationError('Email is required')
  }

  if (!EMAIL_REGEX.test(trimmed)) {
    throw new ValidationError('Please enter a valid email address')
  }
}

/**
 * Validates a name field
 * @throws {ValidationError} if name is invalid
 */
export function validateName(name: string): void {
  const trimmed = name.trim()

  if (!trimmed) {
    throw new ValidationError('Name is required')
  }

  if (trimmed.length < VALIDATION.NAME.MIN_LENGTH) {
    throw new ValidationError(`Name must be at least ${VALIDATION.NAME.MIN_LENGTH} characters long`)
  }

  if (trimmed.length > VALIDATION.NAME.MAX_LENGTH) {
    throw new ValidationError(`Name must not exceed ${VALIDATION.NAME.MAX_LENGTH} characters`)
  }
}

/**
 * Validates user creation data
 * @throws {ValidationError} if validation fails
 */
export function validateCreateUserRequest(data: { name: string; email: string }): void {
  validateName(data.name)
  validateEmail(data.email)
}
