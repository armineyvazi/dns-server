# Use a lightweight base image with Go installed
FROM golang:1.23.3-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the application
RUN go build -o dns-server

# Expose port 53 for DNS queries
EXPOSE 5321/udp

# Run the DNS server
CMD ["./dns-server"]
