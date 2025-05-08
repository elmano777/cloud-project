from pydantic import BaseModel
from datetime import date, datetime

class Calificacion(BaseModel):
    id_nota: int
    codigo_curso: int
    id_estudiante: int
    profesor_id: int

class Nota(BaseModel):
    id_nota: int
    valor: int
    fecha: date
    id_estudiante: int
    codigo_curso: int

    class Config:
        json_encoders = {
            date: lambda v: datetime.combine(v, datetime.min.time())
        }
