# ---------- STAGE 1: BUILDER ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

# Instalamos las dependencias
RUN go mod download

# Copiamos lo demas
COPY . .

RUN go build -o app ./


# ---------- STAGE 2: RUNTIME ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/app"]
