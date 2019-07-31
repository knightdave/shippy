FROM golang:alpine as build_base
RUN apk --no-cache add git
WORKDIR /app/shippy-cli-consignment
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build  -o shippy-cli-consignment -a -installsuffix cgo .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json

COPY --from=builder /app/shippy-cli-consignment/shippy-cli-consignment .


CMD ["./shippy-cli-consignment"]