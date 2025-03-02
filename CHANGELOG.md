# Changelog

All notable changes to this project will be documented in this file.

## [0.0.3] - 2025-03-01
- 🔧 Refactored Gin middleware to use a bucket-token-based algorithm.
- ⚡ Implemented test cases for validation.

## [0.0.2] - 2025-03-01
- ⚡ Implemented test cases for validation.

## [0.0.1] - 2025-03-01
### Added
- 🎉 Initial release of **Rate Limiter for Gin**.
- ⚡ Implemented **bucket-based rate limiting** using the **Token Bucket algorithm**.
- 🌍 Added support for **IP-based and Token-based rate limiting**.
- 🔄 Middleware support for **Gin web framework**.
- 🗄️ In-Memory and **Redis-backed storage support**.
- 🔧 CLI support with `gin-middleware` and `rate-limiter` commands.
- 📄 Added **example usage and documentation**.

---

## Format
This project follows [Semantic Versioning](https://semver.org/).

- **MAJOR**: Breaking changes.
- **MINOR**: Backward-compatible features.
- **PATCH**: Bug fixes and improvements.
