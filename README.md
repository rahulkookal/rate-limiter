# 🚀 Rate Limiter for Gin

[![Go Report Card](https://goreportcard.com/badge/github.com/rahulkookal/rate-limiter)](https://goreportcard.com/report/github.com/rahulkookal/rate-limiter)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/rahulkookal/rate-limiter.svg)](https://pkg.go.dev/github.com/rahulkookal/rate-limiter)

A high-performance, **bucket-based, token-based and ip-based rate limiter** for the **Gin web framework**. Supports in-memory & Redis-backed storage.

## ✨ Features
- ⚡ **Lightweight & Efficient**
- 🌍 **Supports IP-based & Token-based Rate Limiting – Works seamlessly with Gin middleware.**
- 🛠️ **Standalone Bucket-Based Rate Limits – Each limiter works independently without external dependencies.**
- 📝 **Easy Integration**
- 📦 **Supports Redis for Distributed Throttling**(TODO)

---

## 📦 Installation

```sh
go get github.com/rahulkookal/rate-limiter


## 📦 Testing

```sh
go test ./...
