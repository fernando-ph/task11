package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {

	databaseUrl := "postgres://postgres:Gas1234hhH@localhost:5432/projek"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprint(os.Stderr, "Unable to connect to database: %v/n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to database!")
}
