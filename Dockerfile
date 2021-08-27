FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/borpalive ./

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
VOLUME /usr/bin/borpa-data
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/borpalive --port 8080
