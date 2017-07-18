FROM golang:latest

ENV GOPATH=/go

RUN mkdir -p /go/src/djlobosearch
WORKDIR /go/src/djlobosearch
ADD . .

RUN apt-get update && apt-get install -y unzip
ADD https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip /tmp/protoc.zip
RUN unzip /tmp/protoc.zip -d /usr
RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN make build

EXPOSE 8080

CMD GIN_MODE=release ./search
