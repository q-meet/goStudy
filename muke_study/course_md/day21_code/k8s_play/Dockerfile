FROM golang:alpine

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .

RUN go mod init main
RUN go build -o app .

EXPOSE 11114

CMD ["/build/app"]
