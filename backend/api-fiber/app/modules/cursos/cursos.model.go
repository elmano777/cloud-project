package cursos

import (
	"time"
)

type Curso struct {
	Codigo    int
	Nombre    string
	Horario   string
	Ciclo     string
	CreatedAt time.Time
}
