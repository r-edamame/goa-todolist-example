FROM golang:1.20.4-alpine

RUN apk update && apk add vim bash 

WORKDIR /usr/local/protoc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v23.2/protoc-23.2-linux-x86_64.zip
RUN unzip protoc-23.2-linux-x86_64.zip
RUN echo 'export PATH=$PATH:/usr/local/protoc/bin' >> ~/.bashrc

RUN go install goa.design/goa/v3/cmd/goa@v3
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /go/src
