from database import db
from models import Calificacion
from bson import ObjectId

def convert_objectid(doc):
    if "_id" in doc and isinstance(doc["_id"], ObjectId):
        doc["_id"] = str(doc["_id"])
    return doc

coleccion = db["calificacion"]

async def crear_calificacion(data: Calificacion):
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
    result = await coleccion.delete_one({"id_nota": id_nota})
    return result.deleted_count == 1
