FROM golang:1.22-alpine3.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin
RUN swag init -g cmd/main.go

RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]


