FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod ./
COPY src/ ./src/

RUN go mod download
RUN go build -o go-books ./src/main.go

FROM gcr.io/distroless/base

COPY --from=builder /app/go-books /go-books

EXPOSE 8080

CMD ["/go-books"]