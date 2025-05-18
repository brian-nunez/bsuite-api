FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o server ./cmd/

FROM alpine:latest

RUN apk add --no-cache ffmpeg

COPY --from=builder /app/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
