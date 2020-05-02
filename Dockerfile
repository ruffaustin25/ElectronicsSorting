FROM golang:1.14.2-buster

WORKDIR /go/src/github.com/ruffaustin25/HouseManagement
COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

RUN go build -o main main.go

EXPOSE 1234

CMD ["./main"]