# Usa una imagen de Node.js como base
FROM node:16.20.2-alpine

# Establece el directorio de trabajo en la aplicación
WORKDIR /usr/src/app

# Copia los archivos del paquete y el bloque de bloqueo para instalar las dependencias
COPY package*.json ./

# Instala las dependencias
RUN npm install

# Copia los archivos del resto de la aplicación
COPY . .

# Compila la aplicación
RUN npm run build

# Expone el puerto en el que se ejecuta la aplicación
EXPOSE 3000

# Inicia la aplicación80:8080 back_tarea1
CMD ["npm", "start"]

#Notas:
# Crear imagen: docker build -t front_tarea1 .
#docker build -t front_tarea1 .
# Correr imagen: docker run -p 3000:3000 front_tarea1
#docker run -p 3000:3000 front_tarea1