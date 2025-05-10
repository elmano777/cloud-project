#!/bin/bash

# Asegúrate de tener instalado swag
# go install github.com/swaggo/swag/cmd/swag@latest

# Generar la documentación de Swagger
echo "Generando documentación de Swagger..."
swag init -g app/bootstrap/app.go -o ./docs --parseDependency --parseInternal
echo "Documentación generada en la carpeta docs/"
echo "Puedes acceder a la interfaz de Swagger en: http://localhost:8070/swagger"
