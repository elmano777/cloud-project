from database import db
from models import Calificacion
from bson import ObjectId

def convert_objectid(doc):
    if "_id" in doc and isinstance(doc["_id"], ObjectId):
        doc["_id"] = str(doc["_id"])
    return doc

coleccion = db["calificacion"]

async def crear_calificacion(data: Calificacion):
    # Verificar si ya existe una calificación con el mismo id_nota
    existing = await obtener_calificacion(data.id_nota)
    if existing:
        return "calificacion ya existe"  # O manejar de otra manera si la calificación ya existe

    # Si no existe, crear la nueva calificación
    result = await coleccion.insert_one(data.dict())
    created = await coleccion.find_one({"_id": result.inserted_id})
    return convert_objectid(created)

async def obtener_calificaciones():
    calificaciones = []
    async for cal in coleccion.find({}):  # <-- typo fixed
        calificaciones.append(convert_objectid(cal))
    return calificaciones

async def obtener_calificacion(id_nota: int):
    cal = await coleccion.find_one({"id_nota": id_nota})
    if cal:
        return convert_objectid(cal)
    return None

async def eliminar_calificacion(id_nota: int):
    cal_deleted = await db.calificacion.delete_one({"id_nota": id_nota})
    nota_deleted = await db.nota.delete_one({"id_nota": id_nota})
    print(cal_deleted.deleted_count, nota_deleted.deleted_count)

    if cal_deleted.deleted_count == 0:
        return False  # No se encontró calificación

    return True
