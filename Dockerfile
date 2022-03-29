# syntax=docker/dockerfile:1
FROM golang:1.18 AS builder
WORKDIR /go/src/github.com/adntgv/kv-storage-service/
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/adntgv/kv-storage-service/server ./
CMD ["./server"] 
