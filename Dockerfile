FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get -v ./...
RUN go install -v ./...

CMD ["httpr"]
