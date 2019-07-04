FROM golang:1.11

WORKDIR $GOPATH/src/github.com/elgiavilla/mc_ucer

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8000

CMD ["mc_user"]