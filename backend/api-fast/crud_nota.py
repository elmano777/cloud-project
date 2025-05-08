import datetime
from database import db
from models import Nota
from bson import ObjectId

coleccion = db["nota"]

def convert_objectid(doc):
    if "_id" in doc and isinstance(doc["_id"], ObjectId):
        doc["_id"] = str(doc["_id"])
    return doc

async def crear_nota(data: Nota):
    nota_dict = data.dict()

    # Verificar si ya existe una nota con el mismo id_nota
    nota_existente = await obtener_nota(nota_dict["id_nota"])
    if nota_existente:
        return {"error": "Esta nota ya fue asignada"}

    # Convertir date a datetime para compatibilidad con MongoDB
    if isinstance(nota_dict["fecha"], datetime.date) and not isinstance(nota_dict["fecha"], datetime.datetime):
        nota_dict["fecha"] = datetime.datetime.combine(nota_dict["fecha"], datetime.time.min)

    result = await coleccion.insert_one(nota_dict)
    created = await coleccion.find_one({"_id": result.inserted_id})
    return convert_objectid(created)

async def obtener_notas():
    notas = []
    async for doc in coleccion.find({}):
        notas.append(convert_objectid(doc))
    return notas

async def obtener_nota(id_nota: int):
    nota = await coleccion.find_one({"id_nota": id_nota})
    if nota:
        return convert_objectid(nota)
    return None

async def eliminar_nota(id_nota: int):
    result = await coleccion.delete_one({"id_nota": id_nota})
    return result.deleted_count == 1

async def obtener_notas_filtradas(id_estudiante: int, codigo_curso: int):
    # Filtrar las notas según estudiante y curso
    notas = []
    async for doc in coleccion.find({"id_estudiante": id_estudiante, "codigo_curso": codigo_curso}):
        notas.append(convert_objectid(doc))
    return notas
