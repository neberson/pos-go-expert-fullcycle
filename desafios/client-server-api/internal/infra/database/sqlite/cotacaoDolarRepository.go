package database

import (
	"client-server-api/pkg/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

const (
	scriptCreateTable = `CREATE TABLE IF NOT EXISTS Cotacao (
		id VARCHAR(255) PRIMARY KEY,
		code VARCHAR(255),
		codein VARCHAR(255),
		name VARCHAR(255),
        high VARCHAR(255),
        low VARCHAR(255),		
		varbid VARCHAR(255),
		pctChange VARCHAR(255),
		bid VARCHAR(255),
		ask VARCHAR(255),
		timestamp VARCHAR(255),
		create_date VARCHAR(255)
	);`
)

func CreateTable() error {
	connectDB, err := NewConnectDB()
	if err != nil {
		return err
	}
	_, err = connectDB.db.Exec(scriptCreateTable)
	if err != nil {
		return fmt.Errorf("erro ao criar a tabela: %w", err)
	}

	return nil
}

func InsertCotacao(cotacao *entity.Cotacao) error {
	connectDB, err := NewConnectDB()
	if err != nil {
		return err
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	query := ` INSERT INTO Cotacao 
	           		( id, code, codein, name, high, low, varbid, pctChange, bid, ask, timestamp, create_date)
				VALUES
					( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			`
	stmt, err := connectDB.db.PrepareContext(ctxDB, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctxDB, uuid.New().String(),
		&cotacao.Usdbrl.Code,
		&cotacao.Usdbrl.Codein,
		&cotacao.Usdbrl.Name,
		&cotacao.Usdbrl.High,
		&cotacao.Usdbrl.Low,
		&cotacao.Usdbrl.VarBid,
		&cotacao.Usdbrl.PctChange,
		&cotacao.Usdbrl.Bid,
		&cotacao.Usdbrl.Ask,
		&cotacao.Usdbrl.Timestamp,
		&cotacao.Usdbrl.CreateDate)

	if err != nil {
		return err
	}

	return nil
}
