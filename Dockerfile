FROM golang:latest AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/server/main.go

FROM golang:latest AS production
COPY --from=builder /app .
RUN chmod +x ./app

CMD ["./app"]