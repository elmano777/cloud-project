#!/bin/bash

# Asegúrate de tener instalado swag
# go install github.com/swaggo/swag/cmd/swag@latest

# Generar la documentación de Swagger
echo "Generando documentación de Swagger..."
swag init -g app/boostrap/app.go -o ./docs

echo "Documentación generada en la carpeta docs/"
echo "Puedes acceder a la interfaz de Swagger en: http://localhost:8070/swagger"
