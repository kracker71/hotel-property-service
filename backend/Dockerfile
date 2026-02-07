FROM golang:1.25.6-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build -o /myapp ./cmd/server/main.go

# ---- runtime ----
FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=builder /myapp /myapp
COPY --from=builder /app/config.yaml /config.yaml

EXPOSE 3000
CMD ["/myapp"]
