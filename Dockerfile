FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN go build -o k8s-manager

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/k8s-manager .
EXPOSE ${PORT:-8080}
CMD ["./k8s-manager"]