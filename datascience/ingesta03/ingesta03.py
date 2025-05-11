import pymongo
import boto3
import csv


HOST = "172.31.86.125" # IPv4 privada de "MV Bases de Datos"
PORT = "27017"
DATABASE_NAME = "calificaciones"  


# Cambiar el nombre del bucket
BUCKET = "ingesta-bd-storage"
UPLOAD_FILES = [
    ("notaData.csv", "nota"),
    ("calificacionData.csv", "calificacion")
]


for filename, table in UPLOAD_FILES:
    client = pymongo.MongoClient(f"mongodb://{HOST}:{PORT}")  
    db = client[DATABASE_NAME]
    collection = db[table]

    data = list(collection.find())
    columnas = set()

    for doc in data:
        data.update(doc.keys())
    
    columnas = sorted(list(columnas))
    
    with open(filename, mode="w", newline="", encoding="utf-8") as file:
        writer = csv.DictWriter(file, fieldnames=columnas)
        writer.writeheader()
        for doc in data:
            doc.pop("_id", None)
            writer.writerow(doc)


    s3 = boto3.client('s3')
    response = s3.upload_file(filename, BUCKET, filename)

    print(f"Ingesta de {table} completada")