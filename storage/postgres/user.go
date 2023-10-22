package postgres

import (
	models "tgotify/storage"
)

// UserByName retrieves a user by their username from the database.
func (db *Gormdb) UserByName(name string) (*models.User, error) {
	user := new(models.User)
	err := db.Where("name = ?", name).First(user).Error
	if user.Name == name {
		return user, err
	}
	return nil, err
}

// CreateUser creates a new user in the database.
func (db *Gormdb) CreateUser(user *models.User) error {
	return db.Create(user).Error
}

// GetUserByID retrieves a user by their ID from the database.
func (db *Gormdb) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	err := db.DB.First(user, id).Error
	if user.ID == id {
		return user, err
	}
	return nil, err
}

// DeleteUser deletes a user by their ID from the database.
func (db *Gormdb) DeleteUser(id uint) error {
	err := db.Delete(&models.User{}, id).Error
	return err
}
