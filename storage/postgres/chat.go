package postgres

import models "tgotify/storage"

func (db *Gormdb) CreateChat(chat models.Chat) error {
	return db.Create(&chat).Error
}
