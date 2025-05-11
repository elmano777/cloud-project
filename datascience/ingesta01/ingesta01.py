import mysql.connector
import boto3
import csv


HOST = "172.31.86.125" # IPv4 privada de "MV Bases de Datos"
PORT = "8005"
USERNAME = "root"
PASSWORD = "mi_password"
DATABASE_NAME = "bd_api_employees"  


# Cambiar el nombre del bucket
BUCKET = "ingesta-bd-storage"
UPLOAD_FILES = [
    ("cursosData.csv", "cursos"),
    ("estudianteCursosData.csv", "estudiante_cursos")
]


for filename, table in UPLOAD_FILES:
    mydb = mysql.connector.connect(
        host=HOST,
        port=PORT,
        user=USERNAME,
        password=PASSWORD,
        database=DATABASE_NAME
    )  

    cursor = mydb.cursor()

    cursor.execute(f"SELECT * FROM {table}")

    columnas = [desc[0] for desc in cursor.description]
    result = cursor.fetchall()

    with open(filename, mode="w", newline="", encoding="utf-8") as file:
        writer = csv.writer(file)
        writer.writerow(columnas)
        writer.writerows(result)


    cursor.close()
    mydb.close()

    s3 = boto3.client('s3')
    response = s3.upload_file(filename, BUCKET, filename)

    print(f"Ingesta de {table} en archivo {filename} completada")