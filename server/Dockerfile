# Use the official Go image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port for the application (if needed)
EXPOSE 8080

# Define the command to run your application when the container starts
CMD ["./myapp"]
