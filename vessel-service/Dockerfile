# vessel-service/Dockerfile
FROM golang:alpine as build_base
RUN apk --no-cache add git
WORKDIR /app/shippy-service-vessel
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build_base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-vessel .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-vessel .
CMD ["./shippy-service-vessel"]