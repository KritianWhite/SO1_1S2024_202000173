# Dockerfile
FROM mongo:4.4

# Copiar el script de inicialización al directorio /docker-entrypoint-initdb.d/
COPY ./database_init.js /docker-entrypoint-initdb.d/

# Establecer permisos para el script
RUN chmod +x /docker-entrypoint-initdb.d/database_init.js

EXPOSE 27017