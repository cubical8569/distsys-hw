FROM golang:1.14

WORKDIR /server/
COPY . .

RUN go get -d -v ./...
RUN go test -v ./...