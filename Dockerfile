FROM golang:1.19.10-bookworm

WORKDIR /go/src/github.com/ruffaustin25/ElectronicsSorting
COPY . .

RUN go get github.com/ruffaustin25/ElectronicsSorting
# RUN go install -v ./...

RUN go build -o main main.go

EXPOSE 2796

CMD ["./main"]