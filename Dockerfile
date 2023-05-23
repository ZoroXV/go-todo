FROM golang:1.20.4-bullseye as builder
WORKDIR /api
COPY . /api
RUN ["go", "build"]

FROM debian:bullseye
WORKDIR /api
COPY --from=builder /api/go-todo ./
CMD ["./go-todo"]
