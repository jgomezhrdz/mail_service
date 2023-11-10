FROM golang:alpine AS build

# Use a minimal base image for Go
FROM golang:1.19-alpine as builder
# Set the working directory inside the container
WORKDIR /app
# Copy go.mod and go.sum to cache dependencies
COPY go.mod ./
# Download and cache Go dependencies
RUN go mod tidy
# Copy the rest of the application code
COPY . .
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/api/main.go

#------------------------------------------------------------------------------------------
# Use a lightweight base image for the final image
FROM alpine:3.14

ARG PORT=8080

# Set a non-root user for better security
RUN adduser -D -g '' appuser
USER appuser

# Copy the compiled Go binary from the builder image
COPY --from=builder /app/app /app/
# Set the working directory inside the container
WORKDIR /app
# Expose the port that the application will listen on
EXPOSE $PORT
# Run the Go binary
ENTRYPOINT ["./app"]
