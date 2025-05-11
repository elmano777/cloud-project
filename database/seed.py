import os
import psycopg2
import mysql.connector
from pymongo import MongoClient
from dotenv import load_dotenv
from faker import Faker
import random
import bcrypt

load_dotenv()

faker = Faker()

# Configuración PostgreSQL (Usuarios)
pg_conn = psycopg2.connect(
    host=os.getenv("POSTGRES_HOST"),
    port=os.getenv("POSTGRES_PORT"),
    user=os.getenv("POSTGRES_USER"),
    password=os.getenv("POSTGRES_PASSWORD"),
    dbname=os.getenv("POSTGRES_DB")
)
pg_cursor = pg_conn.cursor()

# Configuración MySQL (Cursos)
mysql_conn = mysql.connector.connect(
    host=os.getenv("MYSQL_HOST"),
    port=os.getenv("MYSQL_PORT"),
    user=os.getenv("MYSQL_USER"),
    password=os.getenv("MYSQL_PASSWORD"),
    database=os.getenv("MYSQL_DB")
)
mysql_cursor = mysql_conn.cursor()

# Configuración MongoDB (Calificaciones)
mongo_client = MongoClient(f"mongodb://{os.getenv('MONGO_HOST')}:{os.getenv('MONGO_PORT')}")
# Configuración MongoDB (Calificaciones)
mongo_client = MongoClient(
    f"mongodb://{os.getenv('MONGO_HOST', 'localhost')}:{os.getenv('MONGO_PORT', '27017')}"
)
mongo_db = mongo_client[os.getenv("MONGO_DB", "calificaciones")]
calificaciones_collection = mongo_db["calificacion"]

# Crear 10,000 profesores y 10,000 estudiantes
emails_generados = set()
print("Seeding PostgreSQL (Usuarios)...")
for i in range(2000):
    nombre = faker.first_name()
    apellido = faker.last_name()
    # Asegurarse de que el correo electrónico sea único
    email = faker.email()
    while email in emails_generados:
        email = faker.email()  # Generar uno nuevo si ya existe
    emails_generados.add(email)
    password = bcrypt.hashpw(faker.password().encode(), bcrypt.gensalt()).decode()
    rol = "profesor" if i < 10000 else "estudiante"
    activo = True
    pg_cursor.execute(
        """
        INSERT INTO usuarios (nombre, apellido, email, password, rol, activo)
        VALUES (%s, %s, %s, %s, %s, %s)
        """,
        (nombre, apellido, email, password, rol, activo)
    )
pg_conn.commit()

# Crear 20,000 cursos
print("Seeding MySQL (Cursos)...")
cursos_insertados = set()  # Set to track inserted course codes
for i in range(2000):
    codigo = i + 1000
    if codigo in cursos_insertados:  # Skip if the course code already exists
        continue
    cursos_insertados.add(codigo)
    nombre = faker.catch_phrase()
    horario = f"{faker.day_of_week()} {random.randint(8, 18)}:00-{random.randint(19, 23)}:00"
    ciclo = f"{random.randint(2021, 2025)}-{random.randint(1, 2)}"
    mysql_cursor.execute(
        """
        INSERT INTO cursos (codigo, nombre, horario, ciclo)
        VALUES (%s, %s, %s, %s)
        """,
        (codigo, nombre, horario, ciclo)
    )
mysql_conn.commit()

# Crear 20,000 calificaciones
print("Seeding MongoDB (Calificaciones)...")
for i in range(2000):
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

print("Datos insertados correctamente.")
# Cerrar conexiones
pg_cursor.close()
pg_conn.close()
mysql_cursor.close()
mysql_conn.close()
mongo_client.close()
