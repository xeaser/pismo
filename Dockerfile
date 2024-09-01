# Inital stage: download modules and build the binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o bin/main cmd/main.go

# Copy the binary and run
FROM alpine:3.20.2
RUN addgroup -S app && adduser -S -G app app
RUN mkdir /app
WORKDIR /app
RUN chown -R app:app /app
COPY --chown=app:app --chmod=0740 --from=builder /app/bin/main .
COPY --chown=app:app config/config.yaml /app/config/config.yaml
USER app
EXPOSE 8080
CMD ["./main"]