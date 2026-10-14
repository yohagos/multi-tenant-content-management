FROM golang:1.26-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go-api ./cmd/api

FROM alpine:3.24

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache wget

WORKDIR /app

COPY --from=builder /go-api /app/go-api

COPY .env /app/.env

EXPOSE 8080

CMD [ "/app/go-api" ]