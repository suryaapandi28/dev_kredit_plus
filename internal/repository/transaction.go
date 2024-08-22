package repository

import (
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/cache"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateKonsumen(Konsumens *entity.Konsumens) (*entity.Konsumens, error)
}
type transactionRepository struct {
	db        *gorm.DB
	cacheable cache.Cacheable
}

func NewTransactionRepository(db *gorm.DB, cacheable cache.Cacheable) TransactionRepository {
	return &transactionRepository{db: db, cacheable: cacheable}
}

func (r *transactionRepository) CreateKonsumen(Konsumens *entity.Konsumens) (*entity.Konsumens, error) {

	if err := r.db.Create(&Konsumens).Error; err != nil {
		return Konsumens, err
	}
	return Konsumens, nil
}
