FROM golang:1.24-bookworm

WORKDIR /app

# Install PostgreSQL client for health checks
RUN apt-get update && apt-get install -y postgresql-client make && rm -rf /var/lib/apt/lists/*

# Install goose for database migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copy start script and make it executable
COPY compose/dev/start.sh /start.sh
RUN chmod +x /start.sh

EXPOSE 8080

CMD ["/start.sh"]
