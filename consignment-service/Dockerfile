FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . .

ENV GO111MODULE=on
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
