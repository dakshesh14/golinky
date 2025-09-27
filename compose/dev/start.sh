#!/bin/bash
set -e

# Wait for postgres to be ready
echo "Waiting for PostgreSQL..."
until PGPASSWORD=golinky-password psql -h postgres -U golinky-user -d golinky -c '\q'; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done
echo "PostgreSQL is up - continuing"


echo "Running migrations..."
make apply-migrations



echo "Starting application..."
cd /app && go run cmd/server/main.go