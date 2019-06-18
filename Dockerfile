FROM golang:1.12-alpine3.9
RUN mkdir /go/src/msisdn
RUN apk add --no-cache git \
    && go get github.com/golang/dep/cmd/dep \
	&& apk del git

COPY . /go/src/msisdn

WORKDIR /go/src/msisdn
RUN dep ensure
RUN go test -v ./...
RUN go build ./swagger/cmd/msisdn-server
ENTRYPOINT ["./msisdn-server"]