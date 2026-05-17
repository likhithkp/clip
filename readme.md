# Clip - URL Shortener

A production-ready URL shortener built with Go, Fiber, MongoDB, and Redis. Features user authentication, custom short codes, and high-performance redirects with caching.

## Tech Stack

| Layer         | Technology   |
| ------------- | ------------ |
| Language      | Go 1.25      |
| Web Framework | Fiber v2     |
| Database      | MongoDB      |
| Cache         | Redis        |
| Auth          | JWT + bcrypt |
| DI            | Uber FX      |
| Logging       | Zap          |

## Features

- User signup & signin (bcrypt password hashing)
- JWT authentication with cookie support
- Create short URLs (auto-generated or custom codes)
- Duplicate code detection (409 Conflict)
- Redirect to long URLs (301 Moved Permanently)
- Redis caching for low-latency redirects (~70ms after warm-up)
- MongoDB persistence with unique indexes
- Clean architecture (handler -> service -> repository -> domain)
- Graceful shutdown with FX lifecycle hooks

## API Endpoints

### Auth

POST /api/v1/auth/signup - Create new user
POST /api/v1/auth/signin - Login, receive JWT cookie

### URLs

POST /api/v1/urls - Create short URL (auth required)
GET /api/v1/urls/:code - Redirect to original URL

## Project Structure

clip/
├── application/ # Application layer (handlers, DTOs, convertors)
│ ├── auth/
│ └── url/
├── data_access/ # Data layer
│ ├── mongo/ # MongoDB client + services
│ ├── redis/ # Redis client + services
│ └── repository/ # Repository layer
├── domain/ # Domain entities
├── utils/ # Shared utilities (config, jwt, middleware)
└── main.go # App entry (Uber FX)

## Getting Started

### Prerequisites

- Go 1.25+
- MongoDB (local or Atlas)
- Redis (local or cloud)

### Environment Variables

Create a .env file:
HTTP_PORT=":8080"
MONGODB_URI="mongodb://localhost:27017"
DB_NAME="clip"
JWT_SECRET="your-secret-key"
REDIS_ADDR="localhost:6379"
REDIS_PASSWORD=""

### Run Locally

go mod download
go run main.go

### Test with Curl

# Signup

curl -X POST http://localhost:8080/api/v1/auth/signup -H "Content-Type: application/json" -d '{"email":"test@example.com","firstName":"Test","lastName":"User","password":"12345678"}'

# Signin

curl -X POST http://localhost:8080/api/v1/auth/signin -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"12345678"}' -c cookies.txt

# Create short URL

curl -X POST http://localhost:8080/api/v1/urls -H "Content-Type: application/json" -b cookies.txt -d '{"longUrl":"https://example.com/very/long/url","title":"Example"}'

# Redirect

curl -i http://localhost:8080/api/v1/urls/{code}

## Performance

- Redirect latency (warm): ~70-90ms
- Redirect latency (cold start): ~500ms
- Cache hit: Redis first, MongoDB fallback

## License

MIT
