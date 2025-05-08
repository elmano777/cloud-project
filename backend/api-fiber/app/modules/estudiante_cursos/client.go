package estudiantecurso

import (
	"fmt"
	"net/http"
)

type UserClient struct {
	BaseURL string
}

// NewUserClient - Crea una nueva instancia del cliente de usuario
func NewUserClient() *UserClient {
	return &UserClient{
		BaseURL: "http://api-nest:3000", // URL del microservicio de usuarios en NestJS
	}
}

// ValidarUsuario - Valida si un usuario (estudiante) existe en el microservicio de usuarios
func (c *UserClient) ValidarUsuario(id int) error {
	url := fmt.Sprintf("%s/usuarios/%d", c.BaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error al conectar con el servicio de usuarios: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("estudiante con ID %d no encontrado", id)
	}
	return nil
}
