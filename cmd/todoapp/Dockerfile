# NULTISTAGE DEPLOY 


# PART 1: GO IMAGE 
FROM golang:1.25.6-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/todoapp/exe /app/cmd/todoapp

# PART 2: PRODACTION EXE IMAGE
FROM alpine:3.23
WORKDIR /app 
COPY --from=builder /app/cmd/todoapp/exe /app
COPY --from=builder /app/public /app/public
ENV PROJECT_ROOT=/app
CMD ["/app/exe"]