FROM golang:1.14.2-buster

WORKDIR /go/src/github.com/ruffaustin25/ElectronicsSorting
COPY . .

RUN go get github.com/ruffaustin25/ElectronicsSorting
# RUN go install -v ./...

RUN go build -o main main.go

EXPOSE 2796

CMD ["./main"]