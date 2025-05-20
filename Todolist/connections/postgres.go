package connections

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func CreatePostgresSession(dsn string) *pgxpool.Pool {
	count := 0
	for {
		db, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			count++
		} else {
			log.Default().Println("connected to postgres")
			return db
		}
		if count == 5 {
			log.Default().Println("could not connect to postgres after 5 attempts")
			log.Default().Println("waiting 5 seconds before retrying...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}
