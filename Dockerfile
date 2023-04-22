# Use an official Go runtime as a base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /events-app

# Copy the Go application code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o events-app

# Set the command to run the Go application when the container starts
CMD ["./events-app"]