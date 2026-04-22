# ── Build stage ──────────────────────────────────────────────────────────────
# Debian-based image: receives timely security patches and avoids the
# vulnerability backlog that affects golang:*-alpine builder layers.
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Cache dependency downloads separately from source compilation
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source and build a fully static binary.
# CGO_ENABLED=0 + netgo ensures no libc dependency in the output.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
      -ldflags="-s -w -extldflags=-static" \
      -tags netgo \
      -o /app/car-journal \
      ./cmd/main.go

# ── Runtime stage ─────────────────────────────────────────────────────────────
# distroless/static contains only ca-certificates and tzdata — no shell,
# no package manager, no OS vulnerabilities to scan.
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy only the compiled binary from the builder
COPY --from=builder /app/car-journal .

# Expose the default HTTP port (override with HTTP_PORT env var if needed)
EXPOSE 8080

# All configuration is supplied via environment variables at runtime.
# Required vars:
#   DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASS
# Optional overrides (defaults shown):
#   HTTP_PORT=:8080
#   DB_DIALECT=postgres
ENTRYPOINT ["/app/car-journal"]
