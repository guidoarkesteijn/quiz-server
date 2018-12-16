FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
#Install Git.
RUN apk update && apk add --no-cache git
# Fetch dependencies.
# Using go get.
RUN go get -d -v
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]