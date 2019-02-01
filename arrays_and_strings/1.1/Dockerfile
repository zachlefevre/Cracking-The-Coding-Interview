FROM golang:latest as builder
LABEL maintainer="@zachlefevre"
WORKDIR /go/src/github.com/zachlefevre/Cracking-The-Coding-Interview/tst
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
CMD ["go", "test"]