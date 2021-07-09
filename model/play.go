package model

type CreateRoomRequest struct {
	Password string `json:"password" binding:"max:20"`
}

type CreateRoomResponse struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
}


