package types

type ApiResponse struct {
	Status      int64  `json:"status"`
	Message     string `json:"message"`
	Data        any    `json:"data,omitempty"`
	Description string `json:"description"`
}
