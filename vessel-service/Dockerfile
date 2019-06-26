FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD vessel-service /app/vessel-service

CMD ["./vessel-service"]


# vessel-service/Dockerfile
# FROM golang:alpine as builder
# 
# RUN apk --no-cache add git
# 
# WORKDIR /app
# 
# COPY . .
# 
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-vessel
# 
# FROM alpine:latest
# 
# RUN apk --no-cache add ca-certificates
# 
# RUN mkdir /app
# WORKDIR /app
# COPY --from=builder /app .
# 
# CMD ["./shippy-service-vessel"]
# 