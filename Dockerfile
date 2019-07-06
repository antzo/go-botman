#Dockerfile
#Step 1 Build
FROM golang:1.12.6-alpine AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go get github.com/gorilla/mux
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./src/main.go

#Step 2 Final
FROM alpine:3.8
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080