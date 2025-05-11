import os
import mysql.connector
from dotenv import load_dotenv
from faker import Faker
import random

load_dotenv()
faker = Faker()

# Configuración MySQL (Cursos)
mysql_conn = mysql.connector.connect(
    host=os.getenv("MYSQL_HOST"),
    port=os.getenv("MYSQL_PORT"),
    user=os.getenv("MYSQL_USER"),
    password=os.getenv("MYSQL_PASSWORD"),
    database=os.getenv("MYSQL_DB")
)
mysql_cursor = mysql_conn.cursor()

# Crear 20,000 cursos
print("Seeding MySQL (Cursos)...")
cursos_insertados = set()  # Set to track inserted course codes
for i in range(20000):
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

# Cerrar conexión MySQL
mysql_cursor.close()
mysql_conn.close()

print("Seeding de MySQL completado.")
