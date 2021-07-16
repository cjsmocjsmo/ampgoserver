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
  mkdir ./assets && \
  mkdir ./assets/css && \
  mkdir ./assets/templates && \
  chmod -R +rwx ./static && \
  chmod -R +rwx ./assets

COPY init-mongo.js ./data
COPY assets/p1thumb.jpg ./assets/
COPY assets/css/loginstyles.css ./assets/css/
COPY assets/templates/home.html ./assets/templates/

RUN \
  mkdir ./fsData && \
  mkdir ./fsData/thumb && \
  chmod -R +rwx ./fsData && \
  mkdir ./logs && \
  chmod -R +rwx ./logs && \
  echo "Creating log file" > ./logs/ampgo_log.txt && \
  chmod -R +rwx ./logs/ampgo_log.txt
STOPSIGNAL SIGINT
CMD ["./main"]

