FROM golang:1.23.0 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum primero para aprovechar la caché de Docker
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o prueba-go ./cmd/main.go

# Etapa 2: Creación de la imagen ligera utilizando Distroless
FROM gcr.io/distroless/base-debian12

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/prueba-go /app/prueba-go

COPY .env .env


# Exponer el puerto en el que corre tu aplicación (ajusta según tu configuración)
EXPOSE 8080

# Comando para ejecutar el binario
CMD ["/app/prueba-go"]