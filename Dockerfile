# Build the application from source
FROM golang AS build-stage

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ starter/starter.go
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ worker/worker.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12:debug AS release-stage
COPY --from=build-stage /app/bin .
