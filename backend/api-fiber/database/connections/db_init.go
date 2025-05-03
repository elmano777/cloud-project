package connections

import (
	"context"
	"log"

	"github.com/elmano777/api-go/database/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pgxPool *pgxpool.Pool
var sqlcQueries *generated.Queries

func InitDB() (*pgxpool.Pool, *generated.Queries, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/Fitshare5?sslmode=disable")
	if err != nil {
		return nil, nil, err
	}

	queries := generated.New(pool)
	pgxPool = pool
	sqlcQueries = queries

	log.Println("✅ Conexión a la base de datos establecida correctamente")
	return pool, queries, nil
}

func GetDBConnection() *pgxpool.Pool {
	return pgxPool
}

func GetSQLCQueries() *generated.Queries {
	return sqlcQueries
}
