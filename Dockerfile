FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:3.18

WORKDIR /app

RUN mkdir -p /app/templates /app/public /app/db

COPY --from=builder /app/app .
COPY --from=builder /app/templates/* /app/templates/
COPY --from=builder /app/public/* /app/public/
COPY --from=builder /app/db/schema.sql /app/db/schema.sql

RUN chmod +x ./app

EXPOSE 8080

CMD ["./app"]
