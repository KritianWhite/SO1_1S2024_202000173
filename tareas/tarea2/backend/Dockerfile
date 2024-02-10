# Usa una imagen base de Node.js
FROM node:16.20.2-alpine

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el package.json e instala las dependencias
COPY package*.json ./
RUN npm install

# Copia el resto del código fuente de la aplicación
COPY . .

# Expone el puerto 3000 para que la aplicación sea accesible desde fuera del contenedor
EXPOSE 8000

# Comando para iniciar la aplicación
CMD ["npm", "start"]
