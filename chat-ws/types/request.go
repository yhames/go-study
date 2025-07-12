package types

type BodyRoomRequest struct {
	Name string `json:"name" binding:"required"`
}

type FormRoomRequest struct {
	Name string `form:"name" binding:"required"`
}
