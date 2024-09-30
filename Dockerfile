FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener-svc ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /url-shortener-svc .
COPY --from=builder /app/config ./config

ENV MONGO_URI=mongodb://mongo:27017

CMD ["./url-shortener-svc"]