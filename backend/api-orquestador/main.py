from fastapi import FastAPI, HTTPException, Query
from fastapi.middleware.cors import CORSMiddleware
import httpx
import os

app = FastAPI()

# Configurar CORS para permitir solicitudes desde cualquier dominio
app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://proyecto4-s3.s3-website-us-east-1.amazonaws.com", "http://localhost:4200,  http://proyecto4-s3.s3-website-us-east-1.amazonaws.com"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


USUARIOS_SERVICE_URL = os.getenv("USUARIOS_SERVICE_URL", "http://mi-nestjs-container:3000")
CURSOS_SERVICE_URL = os.getenv("CURSOS_SERVICE_URL", "http://mi-fiber-container:8070")
NOTAS_SERVICE_URL = os.getenv("NOTAS_SERVICE_URL", "http://mi-fastapi-container:8080")


@app.get("/estudiante/{id_estudiante}/cursos")
async def obtener_estudiante_cursos(
    id_estudiante: int,
    limit: int = Query(10, ge=1),
    page: int = Query(1, ge=1)
):
    # Obtener información del estudiante
    estudiante_url = f"{USUARIOS_SERVICE_URL}/usuarios/{id_estudiante}"
    async with httpx.AsyncClient() as client:
        estudiante_response = await client.get(estudiante_url)
        if estudiante_response.status_code != 200:
            raise HTTPException(status_code=500, detail="Error al obtener información del estudiante")

        estudiante = estudiante_response.json()

    # Obtener cursos del estudiante con paginación
    cursos_url = f"{CURSOS_SERVICE_URL}/estudiante-cursos/estudiante/{id_estudiante}"
    params = {"limit": limit, "page": page}
    async with httpx.AsyncClient() as client:
        cursos_response = await client.get(cursos_url, params=params)
        if cursos_response.status_code != 200:
            raise HTTPException(status_code=500, detail="Error al obtener cursos del estudiante")

        cursos_inscritos = cursos_response.json()

    # Obtener detalles de cada curso
    detalles_cursos = []
    for curso in cursos_inscritos:
        codigo_curso = curso["CursoCodigo"]
        curso_detalle_url = f"{CURSOS_SERVICE_URL}/cursos/{codigo_curso}"
        async with httpx.AsyncClient() as client:
            curso_response = await client.get(curso_detalle_url)
            if curso_response.status_code == 200:
                detalles_cursos.append(curso_response.json())
            else:
                detalles_cursos.append({
                    "codigo": codigo_curso,
                    "error": "Información del curso no disponible"
                })

    return {
        "estudiante": estudiante,
        "cursos": detalles_cursos,
        "pagination": {
            "limit": limit,
            "page": page
        }
    }


@app.get("/estudiante/{id_estudiante}/curso/{codigo_curso}/notas")
async def obtener_notas_curso_estudiante(
    id_estudiante: int,
    codigo_curso: int,
    limit: int = Query(10, ge=1),
    page: int = Query(1, ge=1)
):
    # Obtener información del estudiante
    estudiante_url = f"{USUARIOS_SERVICE_URL}/usuarios/{id_estudiante}"
    async with httpx.AsyncClient() as client:
        estudiante_response = await client.get(estudiante_url)
        if estudiante_response.status_code != 200:
            raise HTTPException(status_code=500, detail="Error al obtener información del estudiante")

        estudiante = estudiante_response.json()

    # Obtener notas del estudiante
    notas_url = f"{NOTAS_SERVICE_URL}/notas"
    params = {
        "id_estudiante": id_estudiante,
        "codigo_curso": codigo_curso,
        "limit": limit,
        "page": page
    }
    async with httpx.AsyncClient() as client:
        notas_response = await client.get(notas_url, params=params)
        if notas_response.status_code != 200:
            raise HTTPException(status_code=404, detail="No se encontraron notas para este estudiante en el curso especificado")

        notas = notas_response.json()

    return {
        "estudiante": estudiante,
        "notas": notas,
        "pagination": {
            "limit": limit,
            "page": page
        }
    }
