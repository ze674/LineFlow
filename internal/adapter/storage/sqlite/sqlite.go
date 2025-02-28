package sqlite

import (
	"LineFlow/internal/adapter/storage"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

// GormStorage реализует интерфейс Storage с использованием GORM
type Storage struct {
	db *gorm.DB
}

// New создает новое хранилище, открывая соединение с SQLite и выполняя миграцию модели
func New(dsn string) (*Storage, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	// Автоматическая миграция модели Code
	if err := db.AutoMigrate(&Code{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return &Storage{db: db}, nil
}

// Save сохраняет уникальный код в таблицу codes
func (s *Storage) Save(data string) error {
	op := "storage.sqlite.Save"

	code := NewCode(data) // Создаем новый код

	result := s.db.Create(&code) // Сохраняем код

	if result.Error != nil {
		// Проверяем, содержит ли сообщение об ошибке информацию о нарушении уникальности
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			return errors.New(storage.ErrExists)
		}

		return fmt.Errorf("%s: %w", op, result.Error)

	}
	return nil
}
