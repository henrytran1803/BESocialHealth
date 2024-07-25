# Dockerfile
FROM golang:1.21.6-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the entire content of the current directory to the /app directory in the container
COPY . .

# Download Go modules
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/main ./cmd/main.go

ENTRYPOINT ["/out/main"]
