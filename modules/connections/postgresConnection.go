package connections

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func PostgresConnection() *sql.DB {
	var psqlInfo string

	if os.Getenv("DATABASE_URL") != "" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	} else {
		configJSON, err := ioutil.ReadFile("utilities/postgres-const.json")
		if err != nil {
			fmt.Println("read file postgres const err", err)
			return nil
		}
		var m map[string]interface{}
		err = json.Unmarshal(configJSON, &m)
		if err != nil {
			fmt.Println("unmarshal postgres const err", err)
			return nil
		}

		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			m["host"], int(m["port"].(float64)), m["user"], m["password"], m["dbname"])
	}

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
