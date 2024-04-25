# Etapa de construcción
FROM golang:latest AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia el código fuente de tu servicio de gRPC al contenedor
COPY . .

# Descarga e instala las dependencias del módulo Go
RUN go mod download

# Compila el servicio de gRPC
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc_client .

# Etapa de producción
FROM alpine:latest

# Copia el binario compilado de la etapa anterior
COPY --from=builder /go/src/app/grpc_client /app/grpc_client

# Expone el puerto en el que el cliente gRPC estará escuchando
EXPOSE 3000

# Comando para ejecutar el cliente gRPC
CMD ["/app/grpc_client"]
