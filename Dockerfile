
FROM golang:alpine3.10

LABEL maintainer="Yasin Kızılkaya <vyasinw@gmail.com>"

WORKDIR $GOPATH/src/kont

# Exclude .env file
COPY . .

RUN export GO111MODULE=on
RUN go mod download

RUN go build -o kont .

CMD ["./kont"]