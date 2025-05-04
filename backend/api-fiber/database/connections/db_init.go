package connections

import (
	"api-fiber/database/generated"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	DB      *sql.DB
	Queries *generated.Queries
)

// LoadEnv carga las variables del archivo .env
func LoadEnv() error {
	// Intenta cargar desde .env en el directorio raíz
	err := godotenv.Load()
	if err != nil {
		// Intenta cargar desde diferentes ubicaciones relativas comunes
		paths := []string{"./.env", "../.env", "../../.env"}
		loaded := false

		for _, path := range paths {
			if err := godotenv.Load(path); err == nil {
				loaded = true
				log.Printf("Variables de entorno cargadas desde: %s", path)
				break
			}
		}

		if !loaded {
			return fmt.Errorf("error al cargar archivo .env: %v", err)
		}
	} else {
		log.Println("Variables de entorno cargadas desde el archivo .env")
	}
	return nil
}

func InitDB() (*sql.DB, *generated.Queries, error) {
	// Cargar variables de entorno si aún no se han cargado
	if err := LoadEnv(); err != nil {
		log.Printf("Advertencia: %v", err)
		log.Println("Continuando con las variables de entorno actuales")
	}

	// Obtener variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD") // Cambiado de DB_PASS a DB_PASSWORD para coincidir con tus variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Verificar que todas las variables estén disponibles
	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil, nil, fmt.Errorf("faltan variables de entorno necesarias para la conexión a la BD: USER=%s, HOST=%s, PORT=%s, NAME=%s",
			dbUser, dbHost, dbPort, dbName)
	}

	// Crear DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		dbUser, dbPass, dbHost, dbPort, dbName)

	log.Printf("Intentando conectar a MySQL con DSN: %s:%s@tcp(%s:%s)/%s",
		dbUser, "[PASSWORD]", dbHost, dbPort, dbName)

	var err error
	// Abrir conexión
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("error al abrir conexión a la base de datos: %v", err)
	}

	// Configurar el pool de conexiones
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Verificar la conexión con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = DB.PingContext(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error al verificar conexión a la base de datos: %v", err)
	}

	// Inicializar el objeto Queries
	Queries = generated.New(DB)

	log.Println("✅ Conexión a la base de datos MySQL establecida correctamente")

	return DB, Queries, nil
}

func GetDBConnection() *sql.DB {
	return DB
}

func GetSQLCQueries() *generated.Queries {
	return Queries
}
