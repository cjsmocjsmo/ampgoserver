# FROM arm32v7/golang:1.12.13 AS builder
FROM golang:latest AS builder
RUN mkdir /go/src/ampgo
WORKDIR /go/src/ampgo

COPY ampgoserver.go .
COPY ampgosetup.go .
COPY artist.go .
COPY album.go .
COPY songs.go .
COPY mp3.go .
COPY go.mod .
COPY go.sum .
RUN export GOPATH=/go/src/ampgo
RUN go get -v /go/src/ampgo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/ampgo

# FROM arm32v6/alpine:latest
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/ampgo/main .
RUN \
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./static && \
  mkdir ./assets && \
  mkdir ./assets/css && \
  mkdir ./assets/templates && \
  chmod -R +rwx ./static && \
  chmod -R +rwx ./assets

COPY assets/animals.jpg ./assets/

COPY assets/css/loginstyles.css ./assets/css/

COPY assets/templates/home.html ./assets/templates/
COPY assets/templates/intro.html ./assets/templates/
COPY assets/templates/artist.html ./assets/templates/
COPY assets/templates/album.html ./assets/templates/
COPY assets/templates/song.html ./assets/templates/
COPY assets/templates/playlist.html ./assets/templates/

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

