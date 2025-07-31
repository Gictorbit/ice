FROM golang:1.23 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todoservice ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/todoservice /todoservice

EXPOSE 8080

# Command to run
ENTRYPOINT ["/todoservice"]
