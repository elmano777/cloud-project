package utils

import (
	"encoding/json"
	"time"
)

// CustomDate es un tipo personalizado para manejar fechas en formato "YYYY-MM-DD".
type CustomDate time.Time

// UnmarshalJSON implementa la interfaz json.Unmarshaler para analizar fechas.
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}
