# 🛡️ sharedlib

Shared authentication and logging library for microservices. Provides JWT validation, structured logging middleware, and common utilities.

---

## 📦 Features

- ✅ RS256 JWT validation with JWKS
- 🔐 Middleware for token verification
- 🧾 Logging middleware with distributed request ID support
- ♻️ Pluggable via `net/http` and Gin
- 🚀 Built for reuse across services

---

## 📥 Installation

Make sure your Go version supports modules (1.18+):

```bash
go get github.com/vityasyyy/sharedlib@latest
