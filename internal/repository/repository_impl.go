package repository

import (
	"errors"
	"fmt"
	"github.com/MorZLE/GoParseTSV/config"
	"github.com/MorZLE/GoParseTSV/constants"
	"github.com/MorZLE/GoParseTSV/internal/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRepositoryImpl(cnf *config.Config) (Repository, error) {
	db, err := gorm.Open(postgres.Open(cnf.DB), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.Debug().AutoMigrate(&model.Guid{}, &model.Err{}, &model.ParseFile{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &repositoryImpl{db: db}, nil
}

type repositoryImpl struct {
	db *gorm.DB
}

// Get получает все guid по уникальному номеру
func (r *repositoryImpl) Get(guid string) ([]model.Guid, error) {
	var guidAPI []model.Guid
	if err := r.db.Where("unit_guid = ?", guid).Find(&guidAPI).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrNotFound
		}
		return nil, fmt.Errorf("error get guid: %w", err)
	}
	return guidAPI, nil
}

// Set записывает guid  в базу
func (r *repositoryImpl) Set(guid []model.Guid) error {
	if err := r.db.Create(&guid).Error; err != nil {
		return fmt.Errorf("error create guid: %w", err)
	}
	return nil
}

// SetError записывает ошибку в базу
func (r *repositoryImpl) SetError(filename string, err error) error {
	var guidErr model.Err
	guidErr.File = filename
	guidErr.Err = fmt.Sprintf("%v", err)

	if err := r.db.Create(&guidErr).Error; err != nil {
		return fmt.Errorf("error create guidErr: %w", err)
	}
	return nil
}

// SetFileName записывает имя обработанного файла в базу
func (r *repositoryImpl) SetFileName(filename string) error {
	if err := r.db.Create(&model.ParseFile{File: filename}).Error; err != nil {
		return fmt.Errorf("error create guid: %w", err)
	}
	return nil
}

// GetFileName получает названия обработанных файлов из базы
func (r *repositoryImpl) GetFileName() ([]model.ParseFile, error) {
	var filenames []model.ParseFile
	if err := r.db.Find(&filenames).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return filenames, nil
		}
		return nil, fmt.Errorf("error get filenames: %w", err)
	}
	return filenames, nil
}
