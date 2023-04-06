FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main

FROM alpine
RUN apk add --no-cache ca-certificates
wORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
