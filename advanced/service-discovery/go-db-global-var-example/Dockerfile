# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

#Turn on GO modules
ENV GO111MODULE=on

# Copy local code to the container image.
WORKDIR /app

#Copy over module files
COPY go.mod .
#Get generated when using specific package versions
#COPY go.sum .

#Download dependencies in mod file
RUN go mod download

#Copy everything over to /app
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o mysql-example

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/mysql-example /mysql-example

# Run the web service on container startup.
CMD ["/mysql-example"]