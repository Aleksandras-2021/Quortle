package repository

import (
	"gorm.io/gorm"
)

// GenericRepository is a reusable repository for any GORM model
type GenericRepository[T any] struct {
	db *gorm.DB
}

// NewGenericRepository creates a new generic repository
func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: db}
}

// Select retrieves a record by ID
func (r *GenericRepository[T]) Select(id uint) (*T, error) {
	var model T
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

// Insert inserts a new record
func (r *GenericRepository[T]) Insert(entity *T) error {
	return r.db.Create(entity).Error
}

// Update updates an existing record
func (r *GenericRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// DeleteByID deletes a record by ID
func (r *GenericRepository[T]) Delete(id uint) error {
	var model T
	return r.db.Delete(&model, id).Error
}
