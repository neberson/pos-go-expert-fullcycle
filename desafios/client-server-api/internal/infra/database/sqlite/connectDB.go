package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "modernc.org/sqlite"
)

type ConnectDB struct {
	db *sql.DB
}

var (
	once       sync.Once
	InstanceDB *ConnectDB
)

const (
	VariableDbPath = "SQLITE_DB_PATH"
)

func NewConnectDB() (*ConnectDB, error) {
	var err error
	once.Do(func() {
		dbPath := os.Getenv(VariableDbPath)
		if dbPath == "" {
			err = fmt.Errorf("a variável de ambiente '%v' não está configurada", VariableDbPath)
			return
		}

		if _, fileErr := os.Stat(dbPath); os.IsNotExist(fileErr) {
			err = fmt.Errorf("o arquivo para o banco de dados SQLite não foi encontrado no caminho '%s'", dbPath)
			return
		}

		dbOnce, dbErr := sql.Open("sqlite", dbPath)
		if dbErr != nil {
			err = dbErr
			return
		}

		InstanceDB = &ConnectDB{db: dbOnce}
	})

	return InstanceDB, err
}

func DbClose() {
	if InstanceDB.db != nil {
		InstanceDB.db.Close()
	}
}
