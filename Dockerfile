FROM golang:1.17-alpine as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/beat-invoice-service ./src

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
WORKDIR /beat-invoice-service
COPY --from=builder /go/bin/beat-invoice-service /beat-invoice-service/backend
ENTRYPOINT ["/fleets-system-service/backend"]
