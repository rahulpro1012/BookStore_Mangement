# Start from a base image containing the Go runtime
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /app/main ./cmd

# Expose port 9010 to the outside world
EXPOSE 9010

# Command to run the executable
CMD ["/app/main"]
