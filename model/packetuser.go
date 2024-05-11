package model

type UserPacket struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	PacketID int `json:"packet_id"`
}
