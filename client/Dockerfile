FROM golang:alpine AS build

# Use a minimal base image for Go
FROM golang:1.19-alpine as builder

WORKDIR /tmp

RUN apk update --no-cache && apk add curl --no-cache
RUN curl -sSf https://atlasgo.sh | sh

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

COPY --from=builder /usr/local/bin/atlas/ /usr/local/bin/atlas/
# Copy the compiled Go binary from the builder image
COPY --from=builder /app/app /app/
# Set the working directory inside the container
WORKDIR /app
# Expose the port that the application will listen on
EXPOSE $PORT
# Run the Go binary
ENTRYPOINT ["./app"]
