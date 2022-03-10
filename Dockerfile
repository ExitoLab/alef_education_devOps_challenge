FROM golang:alpine as base
RUN apk --no-cache add ca-certificates git

# RUN mkdir /build
WORKDIR /opt/api

ADD . /opt/api
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN go build -o main .

# runner image
FROM alpine:latest
WORKDIR /app/
COPY --from=base /opt/api .

EXPOSE 5000
CMD ["/app/main"]