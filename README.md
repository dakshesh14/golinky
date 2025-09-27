# GoLinky

A simple URL shortener built with Go, PostgreSQL, and Redis. Designed for fast URL redirection and analytics tracking. The goal of the project was to learn Go, concurrency patterns, and building a feature rich microservice.

https://github.com/user-attachments/assets/10db5922-729d-41ed-9947-c97b72a4a42b

## Features

- Shorten long URLs into compact codes
- Redirect short URLs to their original URLs
- Caching with Redis for high performance
- Analytics tracking for URL clicks (IP address, user agent, timestamp)
- Asynchronous processing with Redis queues
- Modular, feature-driven Go architecture
- Fully dockerized development environment

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.24+ (for local development)
- Goose v3.24+ (for database migrations)

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/golinky.git
cd golinky
```

### 2. Set up environment variables

Replace with your actual configuration values.

```
cp .env.example .env
```

### 3. Start the server

Depending upon your preference, you can either run the server locally or use Docker.

#### Option A: Run Locally with Dockerized DB and Cache

```bash
docker compose -f dev-db.compose.yml up -d
make apply-migration
go run cmd/server/main.go
```

#### Option B: Run Entirely with Docker

```bash
docker compose -f dev.compose.yml up --build
```

I personally prefer _Option A_ for development as it allows for faster iteration. You might also have to update values in `.env` depending upon the option you choose.

## API Endpoints

- `POST /api/links`: Shorten a long URL
- `GET /{code}`: Redirect to the original URL
- `GET /api/analytics/links/{code}/analysis`: Get analytics for a shortened URL

## URL Shortening Algorithm

This project has two method to generate short codes:

- Base62 encoding of integer IDs
- MD5 / SHA hashing with truncation

I might add more algorithms in the future and make them configurable using environment variables.

## Analytics Tracking

- Clicks are tracked asynchronously using Redis queues.
- Stored in PostgreSQL with details like IP address, user agent, and timestamp.
- Consumer pattern ensures non-blocking URL redirection.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. Check the [LICENSE](LICENSE) file for details.
