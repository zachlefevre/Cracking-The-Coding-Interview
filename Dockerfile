FROM golang:latest as builder
LABEL maintainer="@zachlefevre"
WORKDIR /go/src/github.com/zachlefevre/Cracking-The-Coding-Interview/tst
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
CMD ["go", "test"]