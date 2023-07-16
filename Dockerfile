FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN make build

EXPOSE 8888

CMD ["./app"]
