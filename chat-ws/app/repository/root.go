package repository

import (
	"chat-ws/app/config"
	"chat-ws/app/repository/kafka"
	"chat-ws/app/types/schema"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"log"
	"strings"
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

func (s *Repository) ServerSet(ip string, available bool) error {
	qs := "INSERT INTO chatting.server_info(`ip`, `available`) VALUES (?, ?) ON DUPLICATE KEY UPDATE available = VALUES(available)"
	_, err := s.db.Exec(qs, ip, available)
	return err
}

func (s *Repository) InsertChatting(user, message, roomName string) error {
	log.Println("Insert Chatting with wss", "from", user, "message", message, "room", roomName)
	_, err := s.db.Exec("INSERT INTO chatting.chat(room, name, message) VALUES (?, ?, ?)", roomName, user, message)
	return err
}

func (s *Repository) CreateRoom(name string) error {
	qs := "INSERT INTO chatting.room(name) VALUES (?)"
	_, err := s.db.Exec(qs, name)
	return err
}

func (s *Repository) FindRoomByName(name string) (*schema.Room, error) {
	var d schema.Room
	qs := query([]string{"SELECT id, name, created_at, updated_at FROM", room, "WHERE name = ?"})
	err := s.db.QueryRow(qs, name).Scan(&d.Id, &d.Name, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Room not found
		}
		return nil, err // Other error
	}
	return &d, nil
}

func (s *Repository) FindRoomAll() ([]*schema.Room, error) {
	qs := query([]string{"SELECT id, name, created_at, updated_at FROM", room})
	cursor, err := s.db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(); err != nil {
			log.Printf("Error closing cursor: %v", err)
			panic(err)
		}
	}()

	var rooms []*schema.Room
	for cursor.Next() {
		var d schema.Room
		if err := cursor.Scan(&d.Id, &d.Name, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, err
		}
		rooms = append(rooms, &d)
	}
	return rooms, nil
}

func (s *Repository) FindChatByRoomName(roomName string) ([]*schema.Chat, error) {
	qs := query([]string{"SELECT id, room, name, message, send_time FROM", chat, "WHERE room = ? ORDER BY send_time DESC LIMIT 10"})
	cursor, err := s.db.Query(qs, roomName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(); err != nil {
			log.Printf("Error closing cursor: %v", err)
			panic(err)
		}
	}()

	var chats []*schema.Chat
	for cursor.Next() {
		var d schema.Chat
		if err := cursor.Scan(&d.Id, &d.Room, &d.Name, &d.Message, &d.SendTime); err != nil {
			return nil, err
		}
		chats = append(chats, &d)
	}
	return chats, nil
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
