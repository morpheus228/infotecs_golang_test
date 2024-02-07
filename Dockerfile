FROM golang:latest

COPY ./ ./

RUN go mod download
RUN go build -o wallets-app ./cmd/main.go

CMD ["./wallets-app"]