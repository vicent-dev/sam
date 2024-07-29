FROM golang:1.22

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get sam
RUN go build ./cmd/server/main.go

CMD ["/app/main"]