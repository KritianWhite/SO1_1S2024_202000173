
# Etapa de construcción del servidor
FROM rust:1.77.0 as client-build-env
WORKDIR /app/client
ADD . /app/client
RUN cargo build --release

# Etapa final para el contenedor
FROM rust:1.77.0

# Copia el servidor compilado
COPY --from=client-build-env /app/client/target/release /client

# Exponer el puerto 5003
EXPOSE 5003

# CMD para ejecutar tanto el cliente como el servidor
CMD ["/client/client"]