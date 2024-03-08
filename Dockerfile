FROM golang:1.21 as builder

WORKDIR /app
COPY go.* ./
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/api ./cmd/main.go

FROM alpine:latest

COPY --from=builder /app/bin/api .

EXPOSE 8080

CMD ["./api"]
