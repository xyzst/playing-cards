# Standard Playing Cards

    A simple Go application which emulates the behavior of a standard deck of 
    playing cards (Joker not included)
    
## Pre-requisites

- golang, 1.14+
- Docker (optional)

## Instructions

#### W/ Docker
1. `$ git clone git@github.com:xyzst/playing-cards.git Cards`
2. `$ cd Cards`
3. `$ docker build -t xyzst/playing-cards .`
4. `$ docker run -it xyzst/playing-cards`

#### W/O Docker
1. `$ git clone git@github.com:xyzst/playing-cards.git Cards`
2. `$ cd Cards`
3. `$ go test && go run main.go deck.go`