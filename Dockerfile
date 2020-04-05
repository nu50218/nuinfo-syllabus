FROM golang:1.14.1-alpine3.11 as builder

WORKDIR /workdir

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o app main.go

FROM alpine:3.11

COPY --from=builder /workdir/app /app

ENTRYPOINT ["/app"]
