package repository

import (
	"fmt"
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRepositoryImpl(cnf *config.Config) (Repository, error) {
	db, err := gorm.Open(postgres.Open(cnf.DB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.Debug().AutoMigrate(&model.Guid{}, &model.Err{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return &repositoryImpl{db: db}, nil
}

type repositoryImpl struct {
	db *gorm.DB
}

func (r *repositoryImpl) Get(interface{}) error {
	return nil
}
func (r *repositoryImpl) Set(guid []model.Guid) error {
	if err := r.db.Create(&guid).Error; err != nil {
		return fmt.Errorf("error create guid: %w", err)
	}
	return nil
}

func (r *repositoryImpl) SetError(filename string, err error) error {
	var guidErr model.Err
	guidErr.File = filename
	guidErr.Err = err
	if err := r.db.Create(&guidErr).Error; err != nil {
		return fmt.Errorf("error create guidErr: %w", err)
	}
	return nil
}
