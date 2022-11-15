FROM docker.io/library/golang:bullseye

RUN apt-get --allow-releaseinfo-change update

RUN apt-get install -y libgtk-3-dev libasound2-dev libxxf86vm-dev

WORKDIR /app

ADD . /app

RUN go get -d ./...

CMD go run github.com/gucio321/tic-tac-go/cmd/giu-game
