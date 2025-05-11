import os
from pymongo import MongoClient
from dotenv import load_dotenv
from faker import Faker
import random

load_dotenv()
faker = Faker()

# Configuración MongoDB (Calificaciones)
mongo_client = MongoClient(
    f"mongodb://{os.getenv('MONGO_HOST', 'localhost')}:{os.getenv('MONGO_PORT', '27017')}"
)
mongo_db = mongo_client[os.getenv("MONGO_DB", "calificaciones")]
calificaciones_collection = mongo_db["calificacion"]

# Crear 20,000 calificaciones
print("Seeding MongoDB (Calificaciones)...")
for i in range(20000):
    calificacion = {
        "id_nota": i + 1,
        "codigo_curso": random.randint(1000, 20000),
        "id_estudiante": random.randint(10001, 20000),  # Solo estudiantes
        "profesor_id": random.randint(1, 10000)  # Solo profesores
    }

    # Verificar si ya existe una calificación con el mismo id_nota
    if calificaciones_collection.find_one({"id_nota": calificacion["id_nota"]}):
        continue  # Skip if the record already exists

    calificaciones_collection.insert_one(calificacion)

# Cerrar conexión MongoDB
mongo_client.close()

print("Seeding de MongoDB completado.")
