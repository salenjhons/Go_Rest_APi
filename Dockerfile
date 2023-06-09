FROM golang:1.18 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0  GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest as PRODUCTION
COPY --from=builder /app .
CMD ["./app"]