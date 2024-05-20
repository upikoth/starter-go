# Stage 1. Build.

FROM golang:1.22.2-alpine as build

RUN apk add --no-cache make \
	&& rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

WORKDIR /starter-go

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN make build

# Stage 2.

FROM alpine:3.18

COPY --from=build /starter-go/app ./
COPY --from=build /starter-go/docs ./docs

EXPOSE 8888

CMD ["./app"]
