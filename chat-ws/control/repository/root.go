package repository

import (
	"chat-ws-control/config"
	"chat-ws-control/repository/kafka"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type Repository struct {
	config *config.Config
	db     *sql.DB

	Kafka *kafka.Kafka
}

const (
	room       = "chatting.room"
	chat       = "chatting.chat"
	serverInfo = "chatting.server_info"
)

func NewRepository(c *config.Config) (*Repository, error) {
	r := &Repository{
		config: c,
	}
	var err error
	if r.db, err = sql.Open(c.DB.Database, c.DB.Url); err != nil {
		return nil, err
	}
	if r.Kafka, err = kafka.NewKafka(r.config); err != nil {
		return nil, err
	}
	return r, nil
}
