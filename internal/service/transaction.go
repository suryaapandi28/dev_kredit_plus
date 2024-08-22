package service

import (
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/repository"
)

type TransactionService interface {
	CreateKonsumen(konsumen *entity.Konsumens) (*entity.Konsumens, error)
	CreateKonsumenLimit(konsumenlimit *entity.Tenor_limits) (*entity.Tenor_limits, error)
	FindNIKKonsumen(nik string) (*entity.Konsumens, error)
	CheckLimitTenor(kode_konsumen string) (*entity.Konsumens, error)
	FindLimitTenor(kode_konsumen string) (*entity.Konsumens, error)
}

type transacctionService struct {
	konsumenRepository repository.KonsumenRepository
}

func NewTransactionService(konsumenRepo repository.KonsumenRepository) *transacctionService {
	return &transacctionService{konsumenRepository: konsumenRepo}
}
func (s *transacctionService) CreateKonsumen(konsumen *entity.Konsumens) (*entity.Konsumens, error) {
	return s.konsumenRepository.CreateKonsumen(konsumen)
}
