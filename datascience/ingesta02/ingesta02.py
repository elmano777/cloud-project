import psycopg2
import boto3


HOST = "172.31.86.125" # IPv4 privada de "MV Bases de Datos"
PORT = "5342"
USERNAME = "root"
PASSWORD = "utec"
DATABASE_NAME = "bd_api_employees"  


# Cambiar el nombre del bucket
BUCKET = "ingesta-bd-storage"
UPLOAD_FILES = [
    ("usuarioData.csv", "usuario"),
    ("ejerceData.csv", "ejerce")
]


def get_user_data(table_name: str):
    conn = psycopg2.connect(
        host=HOST,
        port=PORT,
        user=USERNAME,
        password=PASSWORD,
        database=DATABASE_NAME
    )  
    cursor = conn.cursor()
    cursor.execute(f"SELECT * FROM {table_name}")
    result = cursor.fetchall()
    cursor.close()
    conn.commit()
    conn.close()
    return result


for filename, table in UPLOAD_FILES:
    with open(filename, "w") as file:
        res = get_user_data(table)


    s3 = boto3.client('s3')
    response = s3.upload_file(filename, BUCKET, filename)

    print(f"Ingesta de {table} completada")