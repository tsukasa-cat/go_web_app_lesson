FROM golang:latest

ENV GOPATH /go

RUN apt-get update -qq && \
    apt-get install -y vim
RUN mkdir /app
WORKDIR /app

RUN go get github.com/oxequa/realize
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
# 
# ADD . .
# RUN go build