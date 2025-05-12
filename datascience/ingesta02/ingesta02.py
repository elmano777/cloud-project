# PostgreSQL tables
import psycopg2
import boto3
import csv


HOST = "" # IPv4 publica
PORT = "5432"
USERNAME = "postgres"
PASSWORD = "utec"
DATABASE_NAME = "postgres"  


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
        dbname=DATABASE_NAME,
        user=USERNAME,
        password=PASSWORD,
    )  

    cursor = conn.cursor()
    cursor.execute(f"SELECT * FROM {table}")
    columnas = [desc[0] for desc in cursor.description]
    result = cursor.fetchall()

    with open(filename, mode="w", newline='', encoding="utf-8") as file:
        writer = csv.writer(file)
        writer.writerow(columnas)
        writer.writerows(result)

    cursor.close()
    conn.commit()
    conn.close()

    s3 = boto3.client('s3')
    response = s3.upload_file(filename, BUCKET, filename)

    print(f"Ingesta de {table} en {filename} completada")