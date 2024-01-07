package postgres

import (
	"fmt"
	models "tgotify/storage"
)

func (db *Gormdb) CreateChat(chat models.Chat) error {
	return db.Create(&chat).Error
}

func (db *Gormdb) GetToken(chatID int) (string, error) {
	var client models.Client

	var chat models.Chat
	if err := db.Where("chat_id = ?", chatID).First(&chat).Error; err != nil {
		return "", err
	}

	if err := db.Model(&chat).Association("Client").Find(&client); err != nil {
		return "", err
	}

	return client.Token, nil
}

func (db *Gormdb) Enable(chatID int) error {
	var chat models.Chat

	if err := db.Where("chat_id = ?", chatID).First(&chat).Error; err != nil {
		return fmt.Errorf("failed to find Chat with chatID %d: %w", chatID, err)
	}

	chat.Enabled = true

	if err := db.Save(&chat).Error; err != nil {
		return fmt.Errorf("failed to update Chat with chatID %d: %w", chatID, err)
	}

	return nil
}

func (db *Gormdb) Disable(chatID int) error {
	var chat models.Chat

	if err := db.Where("chat_id = ?", chatID).First(&chat).Error; err != nil {
		return fmt.Errorf("failed to find Chat with chatID %d: %w", chatID, err)
	}

	chat.Enabled = false

	if err := db.Save(&chat).Error; err != nil {
		return fmt.Errorf("failed to update Chat with chatID %d: %w", chatID, err)
	}

	return nil
}
