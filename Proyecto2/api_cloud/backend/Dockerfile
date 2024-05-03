# Usa una imagen base de Node.js
FROM node:16-alpine

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia el package.json y package-lock.json (si lo hay) al contenedor
COPY package*.json ./

# Instala las dependencias
RUN npm install

# Copia el resto de los archivos del backend al directorio de trabajo
COPY . .

# Expone el puerto en el que se ejecutará el servidor
EXPOSE 2024

# Comando para iniciar el servidor
CMD ["npm", "start"]