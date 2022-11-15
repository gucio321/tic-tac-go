FROM golang:latest

WORKDIR /app

ADD . /app

RUN go get -d ./...

CMD go run github.com/gucio321/tic-tac-go/cmd/terminal-game
