# Use an official Go image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Expose the port your app will run on
EXPOSE 8080

# Command to build and run the app
CMD ["go", "run", "main.go"]
