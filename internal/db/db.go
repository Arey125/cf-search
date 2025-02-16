package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const ADD_PROBLEMS_TABLE = `
CREATE TABLE IF NOT EXISTS problems (
    id INTEGER PRIMARY KEY,
    contest_id INTEGER,
    problemset_name TEXT,
    ind TEXT,
    name TEXT,
    type TEXT,
    points REAL,
    rating INTEGER,
    tags TEXT
);
`

func Connect(dsn string) *sql.DB {
    db, err := sql.Open("sqlite3", dsn)
    if err != nil {
        panic(err)
    }

    _, err = db.Exec(ADD_PROBLEMS_TABLE)
    if err != nil {
        panic(err)
    }

    return db
}
