# Use the official Golang image to create a build stage.
FROM golang:latest as builder

# Set the working directory.
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . .

# Download all the dependencies.
RUN go mod download

# Build the application.
RUN go build -o main .

# Use a minimal base image to package the application.
FROM alpine:latest

# Set the working directory.
WORKDIR /app

# Copy the built executable from the builder stage.
COPY --from=builder /app/main .

# Run the application.
CMD ["./main"]
