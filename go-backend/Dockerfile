# Imagen base para construir la aplicación
FROM golang:1.19-alpine AS build

# Configuración del entorno
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app/src

# Copiar los archivos de configuración y el código fuente
COPY ./project/ ./

# Descargar las dependencias del módulo
RUN go mod download\
    && go build -o main . 


# Imagen base para ejecutar la aplicación
FROM alpine:3.18

# Configuración del entorno
ENV TZ=UTC

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo ejecutable de la compilación anterior
COPY --from=build /app/src/main main 
# Copiar el archivo con las variables de ambiente 
COPY --from=build /app/src/.env .env 

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 8080

# Establecer un usuario no privilegiado
RUN adduser --disabled-password --gecos "" container_user 
# Utiliza el usuario no privilegiado
USER container_user

# Comando para ejecutar la aplicación
CMD ["./main"]

