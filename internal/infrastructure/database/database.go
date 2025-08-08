package database

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"task-be/internal/domain"
	"task-be/internal/infrastructure/config"
	appLogger "task-be/internal/infrastructure/logger"
)

func NewDatabase(ctx context.Context, cfg *config.Config) *gorm.DB {
	log := appLogger.GetLogger()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		log.Error("Failed to connect to database", "error", err, "host", cfg.Database.Host, "port", cfg.Database.Port, "dbname", cfg.Database.Name)
		panic("Failed to connect to database")
	}

	log.Info("Database connected successfully", "host", cfg.Database.Host, "port", cfg.Database.Port, "dbname", cfg.Database.Name)

	err = db.AutoMigrate(&domain.Task{})
	if err != nil {
		log.Error("Failed to migrate database", "error", err)
		panic("Failed to migrate database")
	}

	log.Info("Database migration completed successfully")

	return db
}
