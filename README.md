# 🚀 Rate Limiter for Gin

[![Go Report Card](https://goreportcard.com/badge/github.com/rahulkookal/rate-limiter)](https://goreportcard.com/report/github.com/rahulkookal/rate-limiter)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/rahulkookal/rate-limiter.svg)](https://pkg.go.dev/github.com/rahulkookal/rate-limiter)

A high-performance, **bucket-based, token-based, and IP-based rate limiter** for the **Gin web framework**.  
Supports **in-memory** & **Redis-backed** storage (TODO).  

### ✨ **Features**
- ⚡ **Lightweight & Efficient**
- 🌍 **Token-based and IP-based rate limiting for Gin Middleware**
- 🛠️ **Independent Bucket-Based Rate Limits** (Does not associate with external params)
- 🔧 **Customizable via CLI & Config**
- 🟤️ **Pluggable storage (In-Memory, Redis - TODO)**

---

## 🚀 **Installation**
```sh
go get github.com/rahulkookal/rate-limiter
```

---

## 📝 **Example Usage**

### 1️⃣ **Standalone Rate Limiter**
To test the rate limiter as a standalone component:
```sh
go run main.go rate-limiter --rate 5 --interval 1s
```
OR  
Using `cobra-cli`:
```sh
go run . rate-limiter -r 4 -i 1s
```
**📌 Output**
```sh
✅ Request allowed 1
✅ Request allowed 2
✅ Request allowed 3
✅ Request allowed 4
❌ Request denied 5
✅ Request allowed 6
✅ Request allowed 7
✅ Request allowed 8
✅ Request allowed 9
❌ Request denied 10
...
```

---

### 2️⃣ **Gin Middleware Example**
To run a Gin server with rate limiting:
```sh
go run main.go gin-middleware --mode ip --rate 5 --interval 10s
```
OR  
Using `cobra-cli`:
```sh
go run . gin-middleware -m token -r 5 -i 10s
```

**📌 Example API Calls**
```sh
curl -X GET http://localhost:8080/ping
```
- Allowed ✅ → Returns `200 OK`
- Rate Limit Exceeded ❌ → Returns `429 Too Many Requests`

---

## 🏠 **Folder Structure**
```
rate-limiter/
|── examples/               # Example usage of rate limiter
|   ├── standalone.go       # CLI-based rate limiter
|   └── gin-server.go       # Gin middleware example
|── pkg/
|   ├── ratelimiter/        # Core rate limiter logic
|   └── gin-middleware/     # Middleware implementation for Gin
|── cmd/
|   ├── gin-server.go       # CLI command for running Gin middleware
|   └── rate-limiter.go     # CLI command for standalone rate limiter
|── main.go                 # Entry point for CLI
|── LICENSE                 # Open-source license (Unlicense)
|── README.md               # Documentation
```

---

## 🛠️ **Next Steps**
✅ Add **Redis-backed storage** for distributed rate limiting  
✅ Improve **logging & observability**  

---

### 📛 **License**
This project is licensed under **[MIT License.](https://github.com/rahulkookal/rate-limiter/blob/master/LICENSE)**.

---

