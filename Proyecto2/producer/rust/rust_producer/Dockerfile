
# Etapa de construcción del servidor
FROM rust:1.77.0 as server-build-env
WORKDIR /app/server
ADD . /app/server
RUN cargo build --release

# Etapa final para el contenedor
FROM rust:1.77.0

# Copia el servidor compilado
COPY --from=server-build-env /app/server/target/release /server

# Exponer el puerto 5004
EXPOSE 5004

# CMD para ejecutar tanto el cliente como el servidor
CMD ["/server/server"]