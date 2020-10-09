FROM golang:latest
ENV GOPATH /go
RUN apt-get update -qq && \
    apt-get install -y vim
RUN mkdir /app
WORKDIR /app
RUN go get github.com/gin-gonic/gin
# ADD . .
# RUN go build