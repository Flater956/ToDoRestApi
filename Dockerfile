FROM golang:1.16.4-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o todo-restapi ./cmd/app/main.go

CMD ["./todo-restapi"]