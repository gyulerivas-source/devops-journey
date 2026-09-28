# Laboratorio Central de Infraestructura Cloud y DevOps

Repositorio práctico de administración de sistemas, ingeniería de servidores web, contenerización y despliegue continuo.

## Arquitectura del Proyecto

- **`01-nginx-proxy/`**: Reverse Proxy, terminación SSL/TLS, balanceo de carga (`upstream`) y diagnóstico de códigos de estado HTTP.
- **`02-docker-compose/`**: Orquestación multicontenedor, persistencia con Named Volumes (PostgreSQL) y resolución dinámica de DNS (`127.0.0.11`).
- **`03-multistage-build/`**: Optimización de artefactos de producción y reducción de superficie de ataque (Go + Alpine) reduciendo peso en más de un 95%.

## Requisitos de Ejecución
- Linux (Ubuntu / WSL2)
- Docker Engine & Docker Compose
- Nginx
