FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod init github.com/jagmal/shippy/consignment-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build  -o shippy-service-consignment -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-consignment .

CMD ["./shippy-service-consignment"]





#FROM alpine:latest

#RUN mkdir /app
#WORKDIR /app
#RUN CGO_ENABLED=0 GOOS=linux go build  -o shippy-service-consignment -a -installsuffix cgo main.go repository.go handler.go datastore.go
#ADD shippy-service-consignment /app/shippy-service-consignment

#CMD ["./shippy-service-consignment"]
