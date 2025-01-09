package dbService

import (
	"context"
	"gorm.io/gorm"
	"reflect"
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
	if err := tx.Debug().Where(condition).Create(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return model, nil
}

func (service *DBService) GetOrCreate(ctx context.Context, model interface{}, condition interface{}) (interface{}, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer tx.Commit()

	// get or create a new record
	if err := tx.Debug().Where(condition).FirstOrCreate(model).Error; err != nil {
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
	model := reflect.New(reflect.TypeOf(condition).Elem()).Interface()

	if err := tx.Debug().Where(condition).First(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return model, nil
}

func (service *DBService) GetPaginated(ctx context.Context, condition interface{}, page, pageSize int) (any, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// commit or rollback depending on presence of errors
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// calculate offset
	offset := (page - 1) * pageSize

	// get all
	sliceType := reflect.SliceOf(reflect.TypeOf(condition).Elem())
	models := reflect.New(sliceType).Interface()

	if err := tx.Debug().Offset(offset).Limit(pageSize).Where(condition).Order("created_at DESC").Find(models).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return models, nil
}

func (service *DBService) GetAll(ctx context.Context, condition interface{}) (any, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// commit or rollback depending on presence of errors
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// get all
	sliceType := reflect.SliceOf(reflect.TypeOf(condition).Elem())
	models := reflect.New(sliceType).Interface()

	if err := tx.Debug().Where(condition).Order("created_at DESC").Find(models).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return models, nil
}

func (service *DBService) Update(ctx context.Context, model interface{}, condition interface{}) (interface{}, error) {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Commit()
	if err := tx.Debug().Where(condition).Updates(model).Error; err != nil {
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
	if err := tx.Debug().Where(condition).Updates(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// get
	if err := tx.Where(condition).First(model).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return model, nil
}

func (service *DBService) SoftDelete(ctx context.Context, condition interface{}) error {
	tx := service.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Commit()

	// delete record
	model := reflect.New(reflect.TypeOf(condition).Elem()).Interface()

	if err := tx.Debug().Where(condition).Delete(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
