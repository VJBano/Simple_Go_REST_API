# Use the official Go image as the base
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Set the working directory to the cmd folder
WORKDIR /app/cmd

# Build the Go app
RUN go build -o /app/main .

EXPOSE 4040

# Set the command to run the Go app
CMD ["/app/main"]
