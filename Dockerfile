# FROM arm32v7/golang:1.12.13 AS builder
FROM golang:latest AS builder
RUN mkdir /go/src/ampgoserver
WORKDIR /go/src/ampgoserver

COPY ampgoserver.go .

COPY go.mod .
COPY go.sum .
RUN export GOPATH=/go/src/ampgoserver
RUN go get -v /go/src/ampgoserver
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/ampgoserver

# FROM arm32v6/alpine:latest
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/ampgoserver/main .
RUN \
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./static && \
  chmod -R +rwx ./static

COPY assets/p1thumb.jpg ./static/

RUN \
  mkdir ./fsData && \
  mkdir ./fsData/thumb && \
  mkdir ./fsData/crap && \
  chmod -R +rwx ./fsData && \
  mkdir ./logs && \
  chmod -R +rwx ./logs

STOPSIGNAL SIGINT
CMD ["./main"]

