FROM golang:1.20.0

RUN mkdir /build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./exec ./cmd/main.go

EXPOSE $PORT

CMD ["/build/exec"]