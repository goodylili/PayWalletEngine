# Builder stage
FROM golang:1.19.1 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o paywalletengine cmd/server/main.go

# Final stage
FROM alpine:latest



# Copy the built application binary from the builder stage to the current directory of the final stage
COPY --from=builder /app/paywalletengine .

 # Set the entrypoint command to run the "appName" binary when the container starts
ENTRYPOINT ["./paywalletengine"]


