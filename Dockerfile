# Builder Image
FROM golang:1.20.4-bullseye as builder

WORKDIR /api

# Download all project dependencies
COPY go.mod .
COPY go.sum .
RUN ["go", "mod", "download"]

# Build the HTTP Server Binary
COPY . .
RUN ["go", "build"]

# Runtime Image
FROM debian:bullseye

WORKDIR /api

EXPOSE 8080

COPY --from=builder /api/go-todo ./

CMD ["./go-todo"]
