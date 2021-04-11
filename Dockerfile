FROM golang:1.15

WORKDIR /go/src

COPY . .

RUN GOOS=linux go build main.go

EXPOSE 8000

ENTRYPOINT ("./main")

