# Build stage
FROM golang:1.19

# Install the MySQL client utilities
RUN apt-get update && apt-get install -y default-mysql-client

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o server .

# Run the application
CMD ["./server"]
