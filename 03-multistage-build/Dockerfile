# ==========================================
# ETAPA 1: Construcción (Entorno de compilación)
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /src

# Copiar el código fuente
COPY main.go .

# Compilar un binario estático sin dependencias externas
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/servidor-go main.go

# ==========================================
# ETAPA 2: Producción (Artefacto final limpio)
# ==========================================
FROM alpine:3.19

WORKDIR /app

# Copiar ÚNICAMENTE el ejecutable generado en la etapa anterior
COPY --from=builder /bin/servidor-go .

# Por seguridad, no ejecutamos como root
USER nobody:nobody

EXPOSE 8080

CMD ["./servidor-go"]
