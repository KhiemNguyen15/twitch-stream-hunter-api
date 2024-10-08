# Stage 1: Build stage
FROM golang:1.23.2-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

# Stage 2: Final stage
FROM alpine:edge
WORKDIR /app
COPY --from=build /app/api .
EXPOSE 8080
ENTRYPOINT ["/app/api"]
