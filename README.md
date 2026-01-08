# Identity

A complete authentication and authorization system handling modern identity management use cases.

## Overview

This project is a side project focused on building a simple yet comprehensive identity management system. It provides a complete solution for user authentication and authorization with support for modern security features.

## Features

### Core Authentication
- Sign Up (User Registration)
- Sign In (User Login)
- JWT-based Authorization
- Refresh Token Management
- Forget Password (Password Reset)
- Account Deletion

### Advanced Security
- **Two-Factor Authentication (2FA)** - TOTP-based secondary verification
- **OAuth Integration** - Third-party authentication providers
- **Passkey Support** - Passwordless authentication using WebAuthn

## Project Structure

```
Identity/
├── IdentityService/    # Backend service (Go)
│   ├── main.go
│   └── go.mod
├── IdentityWeb/        # Frontend application (React + TypeScript)
│   ├── src/
│   ├── package.json
│   └── vite.config.ts
└── README.md
```

## Tech Stack

### Backend (IdentityService)
- **Language**: Go 1.25.5
- **Module**: identity-service

### Frontend (IdentityWeb)
- **Framework**: React 19.2.0
- **Language**: TypeScript 5.9.3
- **Build Tool**: Vite 7.2.4
- **Package Manager**: pnpm

## Getting Started

### Prerequisites
- Go 1.25.5 or later
- Node.js (with pnpm)
- Git

### Backend Setup

```bash
cd IdentityService
go mod download
go run main.go
```

The backend server will start on `http://0.0.0.0:8080`

### Frontend Setup

```bash
cd IdentityWeb
pnpm install
pnpm dev
```

The frontend development server will start (default port is typically 5173)

## Development

### Backend Commands
```bash
go run main.go          # Run the server
go build                # Build the binary
```

### Frontend Commands
```bash
pnpm dev                # Start development server
pnpm build              # Build for production
pnpm preview            # Preview production build
pnpm lint               # Run ESLint
```

## Roadmap

This is an active side project. The following features are planned:

- [ ] Complete JWT implementation
- [ ] Refresh token rotation
- [ ] Password reset flow with email verification
- [ ] TOTP-based 2FA implementation
- [ ] OAuth providers (Google, GitHub, etc.)
- [ ] WebAuthn/Passkey integration
- [ ] User profile management
- [ ] Admin panel
- [ ] Rate limiting and security headers

## Contributing

This is a personal side project, but suggestions and improvements are welcome.

## License

TBD
