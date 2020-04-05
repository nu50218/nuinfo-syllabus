FROM golang:1.14-alpine3.11 as builder

WORKDIR /workdir

COPY . .

# https://github.com/google/go-github/issues/1049
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

RUN GOOS=linux CGO_ENABLED=0 go build -o app main.go

FROM alpine:3.11

COPY --from=builder /workdir/app /app

ENTRYPOINT ["/app"]
