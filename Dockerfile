# Use the golang:1.22 base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and verify the Go module dependencies
RUN go mod download && go mod verify && go install github.com/cosmtrek/air@latest

# Copy the entire project to the working directory
COPY . .

# CMD [ "go", "run", "main.go" ]
CMD ["air", "-c", ".air.toml"]



# SKIP THE BUILDING OF THE GO APPLICATION DURING DEVELOPMENT
# Build the Go application and create an executable named hotel-reservation in /bin
# RUN go build -v -o /bin/hotel-reservation

# Run the executable when the container starts
# CMD ["/bin/hotel-reservation"]
