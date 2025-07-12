package schema

import "time"

type Chat struct {
	Id       int64     `json:"id"`
	Room     string    `json:"room"`
	Name     string    `json:"name"`
	Message  string    `json:"message"`
	SendTime time.Time `json:"send_time"`
}
