# --- STAGE 1: Builder ---
# This stage compiles your Go application
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are unchanged
RUN go mod download

# Copy the rest of your application's source code
COPY . .

# Build the Go application
# -o main: specifies the output file name as 'main'
# ./: builds the package in the current directory
# CGO_ENABLED=0: Disables cgo, which creates a statically linked executable
#                This is crucial for creating a truly portable and small image.
RUN CGO_ENABLED=0 go build -o main .

# --- STAGE 2: Runner ---
# This stage creates the final, minimal image that contains only the compiled binary
FROM alpine:latest

# Set the working directory for the final image
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your Go application listens on
EXPOSE 8080

# Command to run the executable when the container starts
CMD ["./main"]