# FROM arm32v7/golang:1.12.13 AS builder
FROM golang:bullseye AS builder
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
RUN apk --no-cache add ca-certificates
# FROM debian:bullseye

# RUN \
# 	apt-get update && \
# 	apt-get -y dist-upgrade && \
# 	apt-get -y autoclean && \
# 	apt-get -y autoremove && \
#   apt-get install -y python3 python3-dev python3-pil python3-mutagen

WORKDIR /root/

COPY --from=builder /go/src/ampgoserver/main .
# COPY create_json.py .

RUN \
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./static && \
  chmod -R +rwx ./static
  
RUN \
  mkdir ./logs && \
  touch ./logs/ampgo_setup_log.txt && \
  touch ./logs/ampgo_server_log.txt && \
  chmod -R +rwx ./logs

COPY assets/p1thumb.jpg ./static/

RUN \
  mkdir ./fsData && \
<<<<<<< HEAD
  mkdir ./fsData/music && \
  mkdir ./fsData/thumb

RUN \
  mkdir ./metadata && \
  chmod -R +rwx ./metadata

=======
  mkdir ./fsData/thumb && \
  mkdir ./fsData/crap && \
  chmod -R +rwx ./fsData 
>>>>>>> f20da2ddcb7a98082434f6523561d1aa4ff66f68

STOPSIGNAL SIGINT
CMD ["./main"]

