FROM golang:alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN go build -o k8s-manager
EXPOSE ${PORT:-8080}

COPY ./k8s-manager /app/k8s-manager

WORKDIR /app

CMD ["sh" "-c" "./k8s-manager"]