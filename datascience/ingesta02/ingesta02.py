# PostgreSQL tables
import psycopg2
import boto3
import csv


HOST = "172.31.86.125" # IPv4 privada de "MV Bases de Datos"
PORT = "5342"
USERNAME = "postgres"
PASSWORD = "postgres"
DATABASE_NAME = "usuario"  


# Cambiar el nombre del bucket
BUCKET = "ingesta-bd-storage"
UPLOAD_FILES = [
    ("usuariosData.csv", "usuarios"),
    ("ejerce.csv", "ejerce")
]



for filename, table in UPLOAD_FILES:
    conn = psycopg2.connect(
        host=HOST,
        port=PORT,
        user=USERNAME,
        password=PASSWORD,
        database=DATABASE_NAME
    )  

    cursor = conn.cursor()
    cursor.execute(f"SELECT * FROM {table}")
    columnas = [desc[0] for desc in cursor.description]
    result = cursor.fetchall()

    with open(filename, mode="w", newline='', encoding="utf-8") as file:
        writer = csv.writer(file)
        writer.writerow(columnas)
        writer.writerow(result)

    cursor.close()
    conn.commit()
    conn.close()

    s3 = boto3.client('s3')
    response = s3.upload_file(filename, BUCKET, filename)

    print(f"Ingesta de {table} en {filename} completada")