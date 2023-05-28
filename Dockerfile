FROM golang:1.20.4-alpine

RUN apk update && apk add vim bash 

WORKDIR /usr/local/protoc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v23.2/protoc-23.2-linux-x86_64.zip
RUN unzip protoc-23.2-linux-x86_64.zip
RUN echo 'export PATH=$PATH:/usr/local/protoc/bin' >> ~/.bashrc

WORKDIR /go/src

