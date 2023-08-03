# Start from the official Go image
FROM golang:latest

# Install the PostgreSQL client
RUN apt-get update && apt-get install -y postgresql-client

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Install migration tool
COPY dbconfig.yml ./
RUN go install github.com/rubenv/sql-migrate/...@latest

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port on which the Go application will listen
EXPOSE 8080

# Command to run the Go application
CMD ["./app run"]
