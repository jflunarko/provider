package model

func (t *User) TableName() string {
	return "user"
}

func (t *Provider) TableName() string {
	return "provider"
}

func (t *Packet) TableName() string {
	return "packet"
}

func (t *UserProvider) TableName() string {
	return "user_provider"
}

func (t *UserPacket) TableName() string {
	return "user_packet"
}
