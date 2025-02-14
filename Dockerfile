FROM golang:1.23-alpine AS builder

WORKDIR /necrest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o necrest .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /necrest/necrest .

EXPOSE 12190

CMD ["./necrest", "serve"]
