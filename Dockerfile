FROM golang:alpine3.10
LABEL maintainer="Yasin Kızılkaya <vyasinw@gmail.com>"
WORKDIR $GOPATH/src/kont

COPY . .

RUN export GO111MODULE=on \
    && go mod download \
    && go build -o kont .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /code/
RUN mkdir /var/lib/kont
COPY --from=0 /go/src/kont .
CMD ["./kont"]