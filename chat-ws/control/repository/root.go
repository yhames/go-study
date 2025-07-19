package repository

import (
	"chat-ws-control/config"
	"chat-ws-control/repository/kafka"
	"chat-ws-control/types/table"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type Repository struct {
	config *config.Config
	db     *sql.DB

	Kafka *kafka.Kafka
}

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

func (r *Repository) GetAvailableServers() ([]*table.ServerInfo, error) {
	qs := "SELECT ip, available FROM chatting.server_info WHERE available = true"
	cursor, err := r.db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var servers []*table.ServerInfo
	for cursor.Next() {
		var server table.ServerInfo
		if err := cursor.Scan(&server.Ip, &server.Available); err != nil {
			return nil, err
		}
		servers = append(servers, &server)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return servers, nil
}
