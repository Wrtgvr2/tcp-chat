FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o server-bin ./server/server.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/server-bin /app/server

CMD ["/app/server"]