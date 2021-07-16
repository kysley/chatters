FROM golang:1.16-alpine
WORKDIR /usr/src/server
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .
RUN go build -o /borpaserver
EXPOSE 8081
CMD ["/borpaserver"]


FROM node:latest-alpine
WORKDIR /usr/src/client
COPY package.json ./
COPY yarn.lock ./
RUN yarn
COPY . .
RUN yarn build
EXPOSE 5000
ENV HOST=0.0.0.0
#find a better way to start the built app
CMD [ "yarn", "start" ]
