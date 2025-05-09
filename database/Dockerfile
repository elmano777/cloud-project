# Usa la imagen oficial de Python
FROM python:3.11-slim

# Crear el directorio de trabajo
WORKDIR /app

# Copiar archivos de configuración
COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

COPY seed.py seed.py
COPY .env .env

# Esperar a que las bases de datos estén listas antes de ejecutar el script
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh
