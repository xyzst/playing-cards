FROM golang:buster

WORKDIR /go/cards
COPY . .

RUN go test

CMD ["go", "run", "./main.go", "./deck.go"]