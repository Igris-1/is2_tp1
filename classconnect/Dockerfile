# Usa una imagen oficial de Golang
FROM golang:1.24-alpine

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos de dependencias
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el código fuente
COPY . .

# Compila la aplicación
RUN go build -o app ./cmd/main.go

# Define el puerto
EXPOSE 8080

# Comando para ejecutar el servicio
CMD ["/app/app"]
