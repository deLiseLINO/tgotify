package postgres

import (
	models "tgotify/storage"
)

func (db *Gormdb) CreateClient(client models.Client) error {
	return db.Create(&client).Error
}

// Returns Clients with enabled tokents and enabled chats
func (db *Gormdb) EnabledClients(uid uint) ([]models.Client, error) {
	clients := []models.Client{}
	err := db.
		Joins("JOIN chats ON clients.id = chats.client_id").
		Where("clients.user_id = ? AND clients.enabled = ? AND chats.enabled = ?", uid, true, true).
		Preload("Chats").
		Find(&clients).Error
	return clients, err
}
