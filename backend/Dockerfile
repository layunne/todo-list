FROM golang:alpine3.12 as builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:3.12

COPY --from=builder app/app .

CMD ["./app"]
