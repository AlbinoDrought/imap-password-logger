FROM golang:alpine as builder

RUN apk update && apk add git
COPY . $GOPATH/src/github.com/AlbinoDrought/imap-password-logger
WORKDIR $GOPATH/src/github.com/AlbinoDrought/imap-password-logger

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go get -d -v && go build -a -installsuffix cgo -o /go/bin/imap-password-logger

FROM scratch

EXPOSE 1143
COPY --from=builder /go/bin/imap-password-logger /go/bin/imap-password-logger

ENTRYPOINT ["/go/bin/imap-password-logger"]
