package estudiantecurso

import "time"

type EstudianteCurso struct {
	EstudianteID int
	CursoCodigo  string
	InscritoEn   time.Time
}
