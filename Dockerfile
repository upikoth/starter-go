FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go mod download

RUN make build

EXPOSE 8080

CMD ["./app"]
