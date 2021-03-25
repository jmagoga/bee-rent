package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq" // tirar essa se houver problemas
)

// (3:30) nao entendi esses beforequery and afterquery ainda. EQuimper.

type DBLogger struct {}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	fmt.Println("db >>", db) // 5432
	fmt.Println("ran before return")

	return db
}

// migrate create -ext sql -dir postgres/migrations create_bees   
// migrate -path "postgres/migrations" -database "$POSTGRESQL up"
// comando acima roda a migracao dos arquivos da pasta migrations. aquele
// $POSTGRESQL eh a url que botei no .env e depois meti source .env