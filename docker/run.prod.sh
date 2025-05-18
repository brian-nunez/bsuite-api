#!/bin/bash

set -e

docker build . -f Dockerfile -t bsuite-api
docker run -d \
  --name bsuite-api \
  -p 8080:8080 \
  -e "ENV=production" \
  bsuite-api

