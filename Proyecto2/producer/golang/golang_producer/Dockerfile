# Usa la imagen oficial de Golang como base
FROM golang:1.22-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /grcpserver

# Copia el archivo go.mod y go.sum al directorio de trabajo
COPY go.mod .
COPY go.sum .

# Descarga las dependencias del módulo
RUN go mod download

# Copia el resto de los archivos al directorio de trabajo
COPY . .

# Compila el código Go en un binario llamado "app"
RUN go build -o app .

# Expone el puerto 8080 en el contenedor
EXPOSE 5001

# Comando por defecto para ejecutar la aplicación
CMD ["./app"]