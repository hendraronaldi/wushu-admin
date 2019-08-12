package connections

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sonic_blade"
	dbname   = "wushutp"
)

func PostgresConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		return nil
	}

	return db
}

func InsertPostgresData(conn *sql.DB, tableName string, fields []string, data []interface{}) int {
	var columns string
	var colIdx string

	for i, column := range fields {
		if i == len(fields)-1 {
			columns += column
			colIdx += `$` + fmt.Sprint(i+1)
		} else {
			columns += column + `, `
			colIdx += `$` + fmt.Sprint(i+1) + `, `
		}
	}

	sqlStatement := `
	INSERT INTO ` + tableName + ` (` + columns + `)
	VALUES (` + colIdx + `)
	RETURNING ` + tableName + `_id`

	id := 0
	err := conn.QueryRow(sqlStatement, data...).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return id
}
