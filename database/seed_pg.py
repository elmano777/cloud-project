import os
import psycopg2
from dotenv import load_dotenv
from faker import Faker
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

# Crear 10,000 profesores y 10,000 estudiantes
emails_generados = set()
print("Seeding PostgreSQL (Usuarios)...")
for i in range(20000):
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

# Cerrar conexión PostgreSQL
pg_cursor.close()
pg_conn.close()

print("Seeding de PostgreSQL completado.")
