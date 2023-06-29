# Utiliza una imagen base oficial de PostgreSQL con la versión específica que deseas
FROM postgres:15.3-alpine

# Copia un archivo de inicialización personalizado a la imagen
COPY ./init/scripts/01-init.sql /docker-entrypoint-initdb.d/

# Establece el usuario no privilegiado "postgres" para ejecutar el servidor PostgreSQL
USER postgres

# Inicia el servidor PostgreSQL
CMD ["postgres"]

