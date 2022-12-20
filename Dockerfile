# FROM arm32v7/golang:1.12.13 AS builder
FROM golang:bullseye AS builder
RUN mkdir /go/src/ampgoserver
WORKDIR /go/src/ampgoserver

COPY ampgoserver.go .
COPY ampgosetup.go .
COPY db.go .
COPY lib.go .
COPY firstlib.go .
COPY typeslib.go .
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
  mkdir ./Data && \
  mkdir ./Data/db && \
  mkdir ./static && \
  chmod -R +rwx ./static
  
RUN \
  mkdir ./logs && \
  touch ./logs/ampgo_setup_log.txt && \
  chmod -R +rwx ./logs

COPY assets/p1thumb.jpg ./static/

RUN \
  mkdir ./fsData && \
  mkdir ./fsData/music && \
  mkdir ./fsData/music/music && \
  mkdir ./fsData/music/thumb && \
  mkdir ./fsData/metadata && \
  chmod -R +rwx ./fsData 

STOPSIGNAL SIGINT
CMD ["./main"]

