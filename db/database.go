package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Respuestas struct {
	ID        int
	Comando   string
	Respuesta string
}

var db *sql.DB

func SQLiteConn() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}

	return db
}

func (res *Respuestas) GetMsg(cmd string) (*Respuestas, error) {
	db := SQLiteConn()
	q := `SELECT id, comando, respuesta FROM messages WHERE comando = ?`

	err := db.QueryRow(q, cmd).Scan(&res.ID, &res.Comando, &res.Respuesta)
	defer db.Close()
	if err != nil {
		return &Respuestas{}, err
	}

	return res, nil
}
