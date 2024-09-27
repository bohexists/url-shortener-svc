# Start with a Golang base image
FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /url-shortener-svc ./cmd/main.go

EXPOSE 8080

CMD ["/url-shortener-svc"]