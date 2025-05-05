from fastapi import FastAPI, HTTPException
from models import Calificacion, Nota
import crud_calificaciones, crud_nota

app = FastAPI()

# --- Endpoints Calificacion ---
@app.post("/calificacion/")
async def crear_calificacion(data: Calificacion):
    return await crud_calificaciones.crear_calificacion(data)

@app.get("/calificacion/")
async def listar_calificaciones():
    return await crud_calificaciones.obtener_calificaciones()

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
    return {"mensaje": "Eliminada correctamente"}

# --- Endpoints Nota ---
@app.post("/nota/")
async def crear_nota(data: Nota):
    return await crud_nota.crear_nota(data)

@app.get("/nota/")
async def listar_notas():
    return await crud_nota.obtener_notas()

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
