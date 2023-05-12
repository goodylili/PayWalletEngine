# FROM golang:1.19-alpine AS build
# ADD . /src
# WORKDIR /src
# # RUN go get -d -v -t
# RUN GOOS=linux GOARCH=amd64 go build -v -o paywalletengine

# FROM alpine:3.17.2
# # EXPOSE 8080
# CMD ["paywalletengine"]
# ENV VERSION 1.1.4
# COPY --from=build /src/paywalletengine /usr/local/bin/paywalletengine
# RUN chmod +x /usr/local/bin/paywalletengine

FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o paywalletengine

# Set the correct permissions for the binary
RUN chmod +x paywalletengine

CMD ["./paywalletengine"]