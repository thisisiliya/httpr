FROM golang:1.18

WORKDIR /app

COPY . .

RUN go get -v ./...
RUN go install -v ./...

CMD ["httpr"]
