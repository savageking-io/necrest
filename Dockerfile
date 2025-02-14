FROM golang:1.23-alpine AS builder

ARG CONFIG_DIR=/etc/noerrorcode
ARG CONFIG_FILE=rest.yaml

WORKDIR /necrest

COPY go.mod go.sum ./

RUN mkdir /etc/noerrorcode

COPY ${CONFIG_FILE} ${CONFIG_DIR}/${CONFIG_FILE}

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o necrest .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /necrest/necrest .

EXPOSE 12190

CMD ["./necrest", "serve"]
