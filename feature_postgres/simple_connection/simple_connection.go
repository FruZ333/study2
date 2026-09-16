package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SimpleConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "posrgres://postgres:matvej12@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключение к базе данных произошло успешно!")
}
