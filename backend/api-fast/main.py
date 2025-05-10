from fastapi import FastAPI, HTTPException
from models import Calificacion, Nota
import crud_calificaciones, crud_nota

app = FastAPI()

# --- Endpoints Calificacion ---
@app.post("/calificacion/")
async def crear_calificacion(data: Calificacion):
    return await crud_calificaciones.crear_calificacion(data)

@app.get("/calificacion/")
async def listar_calificaciones(limit: int = 10, page: int = 1):
    return await crud_calificaciones.obtener_calificaciones(limit, page)

@app.get("/calificacion/{id_nota}")
async def obtener_calificacion(id_nota: int):
    cal = await crud_calificaciones.obtener_calificacion(id_nota)
    if not cal:
        raise HTTPException(status_code=404, detail="No encontrada")
    return cal

@app.delete("/calificacion/{id_nota}")
async def eliminar_calificacion(id_nota: int):
    if not await crud_calificaciones.eliminar_calificacion(id_nota):
        raise HTTPException(status_code=404, detail="No encontrada")
    return {"msg": "Calificación y nota eliminadas"}

# --- Endpoints Nota ---
@app.post("/nota/")
async def crear_nota(data: Nota):
    return await crud_nota.crear_nota(data)

@app.get("/nota/")
async def listar_notas(limit: int = 10, page: int = 1):
    return await crud_nota.obtener_notas(limit, page)

@app.get("/nota/{id_nota}")
async def obtener_nota(id_nota: int):
    nota = await crud_nota.obtener_nota(id_nota)
    if not nota:
        raise HTTPException(status_code=404, detail="No encontrada")
    return nota

@app.delete("/nota/{id_nota}")
async def eliminar_nota(id_nota: int):
    if not await crud_nota.eliminar_nota(id_nota):
        raise HTTPException(status_code=404, detail="No encontrada")
    return {"mensaje": "Eliminada correctamente"}

@app.get("/notas")
async def obtener_notas_filtradas(id_estudiante: int, codigo_curso: int):
    if id_estudiante is None or codigo_curso is None:
        raise HTTPException(status_code=400, detail="Faltan parámetros: id_estudiante y codigo_curso son requeridos")

    notas = await crud_nota.obtener_notas_filtradas(id_estudiante, codigo_curso)
    if not notas:
        raise HTTPException(status_code=404, detail="No se encontraron notas para este estudiante en el curso especificado")
    return notas
