FROM    golang:1.14.4-alpine3.12 as builder
RUN     apk update && apk add bash ca-certificates git gcc g++ libc-dev binutils file
WORKDIR /usr/src/app/
COPY    go.mod go.sum ./
RUN     go mod download
COPY    . .
RUN     go build -o /server ./cmd/server/

FROM    alpine:3.12 AS production
RUN     apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
WORKDIR /app
COPY    --from=builder /server ./
CMD     ["./server"]