#!/usr/bin/env bash
set -euo pipefail

# Start local dependencies and run server
docker-compose up -d

go run ./cmd/server
