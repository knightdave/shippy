FROM golang:alpine as builder
RUN apk --no-cache add git
WORKDIR /app/shippy-cli-consignment
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build  -o shippy-cli-consignment -a -installsuffix cgo .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json

COPY --from=builder /app/shippy-cli-consignment/shippy-cli-consignment .


CMD ["./shippy-cli-consignment"]