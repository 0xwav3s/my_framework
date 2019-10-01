FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod . 
COPY go.sum .
# cache dependencies
RUN go mod download
COPY . .
# cache go build
RUN go build -o /go/bin/main
# Add 
FROM alpine:3.5 
WORKDIR /app
COPY --from=builder /go/bin/main .
COPY ./config/config.yaml ./config/ 
ENV GO_MONGO_DATABASE__HOST mongo-docker
#
ENTRYPOINT ["./main"]
