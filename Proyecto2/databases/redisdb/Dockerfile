# Usa la imagen base oficial de Redis
FROM redis:latest

# Copia el archivo de configuración de Redis al contenedor
COPY database_init.conf /usr/local/etc/redis/database_init.conf

# Inicia Redis con el archivo de configuración especificado y luego ejecuta los comandos de inicialización
CMD ["redis-server", "/usr/local/etc/redis/database_init.conf"]
