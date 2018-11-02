#!/usr/bin/env bash
export MGO_HOSTS=localhost:27017
export MGO_DATABASE=cards
export MGO_USERNAME=tigerbeatle
export MGO_PASSWORD=Mindstorm451

cd $GOPATH/src/github.com/tigerbeatle/cards
go clean -i
go build

./cards
