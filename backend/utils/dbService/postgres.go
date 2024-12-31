package dbService

import (
	"context"
	"gorm.io/gorm"
)

type DBService struct {
	DB *gorm.DB
}

func (service *DBService) Create(ctx context.Context, model interface{}, condition interface{}) (interface{}, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Commit()

	// get or create a new record
	if err := tx.Where(condition).FirstOrCreate(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return model, nil
}

func (service *DBService) Get(ctx context.Context, condition interface{}) (interface{}, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Commit()

	// get or create a new record
	if err := tx.Where(condition).First(condition).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return condition, nil
}

func (service *DBService) Update(ctx context.Context, model interface{}, condition interface{}) (interface{}, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Commit()
	if err := tx.Where(condition).Updates(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return model, nil
}

func (service *DBService) UpdateAndGet(ctx context.Context, model interface{}, condition interface{}) (interface{}, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Commit()

	// update
	if err := tx.Where(condition).Updates(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// get
	user, err := service.Get(ctx, condition)
	if err != nil {
		return nil, err
	}

	return user, nil
}
