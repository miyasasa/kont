
FROM golang:alpine3.10

LABEL maintainer="Yasin Kızılkaya <vyasinw@gmail.com>"

WORKDIR $GOPATH/src/miya
COPY . .

RUN export GO111MODULE=on
RUN go mod download

RUN go build -o miya .

CMD ["./miya"]