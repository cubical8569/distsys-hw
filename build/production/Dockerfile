FROM    golang:1.14.4-alpine3.12 as builder
ARG     MOD_NAME
RUN     apk update && apk add bash ca-certificates git gcc g++ libc-dev binutils file
WORKDIR /usr/src/${MOD_NAME}/
COPY    ${MOD_NAME}/go.mod ${MOD_NAME}/go.sum ./
RUN     go mod download
COPY    ${MOD_NAME}/ .
RUN     go build -o /app ./

FROM    alpine:3.12 AS production
RUN     apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
WORKDIR /app
COPY    --from=builder /app ./
CMD     ["./app"]