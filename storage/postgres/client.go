package postgres

import (
	models "tgotify/storage"
)

func (db *Gormdb) CreateClient(client models.Client) error {
	return db.Create(&client).Error
}

// Returns Clients with enabled tokens and enabled chats for user
func (db *Gormdb) EnabledClientsForUser(uid uint) ([]models.Client, error) {
	clients := []models.Client{}
	err := db.
		Joins("JOIN chats ON clients.id = chats.client_id").
		Where("clients.user_id = ? AND clients.enabled = ? AND chats.enabled = ?", uid, true, true).
		Preload("Chats").
		Find(&clients).Error
	return clients, err
}

func (db *Gormdb) EnabledClients() ([]models.EnabledClientsResponse, error) {
	var clients []models.EnabledClientsResponse
	err := db.Table("clients").
		Select("id, token").
		Where("clients.enabled = ?", true).
		Find(&clients).Error
	return clients, err
}

func (db *Gormdb) Clients(uid uint) (*[]models.ClientResponse, error) {
	clients := []models.ClientResponse{}
	err := db.Table("clients").
		Select("id, name, token, enabled").
		Where("clients.user_id = ?", uid).
		Find(&clients).Error
	return &clients, err
}

func (db *Gormdb) DeleteClient(uid uint, id uint) error {
	return db.Where("user_id = ? AND id = ?", uid, id).Delete(&models.Client{}).Error
}

func (db *Gormdb) UpdateClient(uid uint, client models.ClientResponse) error {
	var err error
	if client.Name != "" {
		err = db.Model(&models.Client{}).Where("user_id = ? AND id = ?", uid, client.ID).Update("name", client.Name).Error
	}
	if client.Token != "" {
		err = db.Model(&models.Client{}).Where("user_id = ? AND id = ?", uid, client.ID).Update("token", client.Token).Error
	}
	if client.Enabled != "" {
		err = db.Model(&models.Client{}).Where("user_id = ? AND id = ?", uid, client.ID).Update("enabled", client.Enabled).Error
	}
	return err
}
