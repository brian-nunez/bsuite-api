#!/bin/sh

set -e

docker image prune -f
docker build . -f Dockerfile.dev -t bsuite-api
docker compose -f docker/docker-compose.yml up --no-recreate --attach bsuite-api --no-log-prefix
