import pymongo
import boto3
import pandas
import os
from dotenv import load_dotenv


load_dotenv()


HOST = "172.31.86.125" # IPv4 privada de "MV Bases de Datos"
PORT = "5342"
USERNAME = "root"
PASSWORD = "utec"
DATABASE_NAME = "calificaciones"  


# Cambiar el nombre del bucket
BUCKET = "ingesta-bd-storage"
UPLOAD_FILES = [
    ("notaData.csv", "nota"),
    ("calificacionData.csv", "calificacion")
]


def get_user_data(table_name: str):
    conn = pymongo.MongoClient(
        f""
    )  
    db = conn[DATABASE_NAME]
    collection = conn[table_name]

    data = []
    for doc in collection.find():
        data.append(doc)
    
    return data


for filename, table in UPLOAD_FILES:
    with open(filename, "w") as file:
        res = get_user_data(table)


    s3 = boto3.client('s3')
    response = s3.upload_file(filename, BUCKET, filename)

    print(f"Ingesta de {table} completada")