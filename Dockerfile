FROM golang:alpine as builder
RUN apk --no-cache add ca-certificates git

# RUN mkdir /build
WORKDIR /opt/api

ADD . /opt/api
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN go build -o main .

EXPOSE 5000

CMD ["/opt/api/main"]