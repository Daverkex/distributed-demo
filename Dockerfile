# Build the application from source

FROM golang AS build-stage

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download dependencies
RUN go mod download
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ starter/starter.go
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ worker/worker.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12:debug AS release-stage

# Copy the binary from the build stage into the release stage
COPY --from=build-stage /app/bin .
