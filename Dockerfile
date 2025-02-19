FROM golang:1.23-alpine AS builder

ARG CONFIG_DIR=/etc/noerrorcode

WORKDIR /necrest

COPY go.mod go.sum ./

COPY config/rest.yaml ./
COPY config/kafka.yaml ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o necrest .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /necrest/necrest .

RUN mkdir /etc/noerrorcode

COPY config/rest.yaml /etc/noerrorcode/rest.yaml
COPY config/kafka.yaml /etc/noerrorcode/kafka.yaml

EXPOSE 12190

CMD ["./necrest", "serve", "--log=trace"]
