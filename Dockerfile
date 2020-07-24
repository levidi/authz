FROM golang:1.14.4 as builder

ENV GO111MODULE=on PATH_CONFIG_MAP=./configMap

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM alpine:3.12.0
COPY --from=builder /app/authz .

EXPOSE 50051

ENTRYPOINT ["/authz"]
