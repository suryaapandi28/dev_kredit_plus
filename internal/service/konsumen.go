package service

import (
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/repository"
)

type KonsumenService interface {
	CreateKonsumen(konsumen *entity.Konsumens) (*entity.Konsumens, error)
	CreateKonsumenLimit(konsumenlimit *entity.Tenor_limits) (*entity.Tenor_limits, error)
	FindNIKKonsumen(nik string) (*entity.Konsumens, error)
	CheckLimitTenor(kode_konsumen string) (*entity.Konsumens, error)
	FindLimitTenor(kode_konsumen string) (*entity.Konsumens, error)
}

type konsumenService struct {
	konsumenRepository repository.KonsumenRepository
}

func NewKonsumenService(konsumenRepo repository.KonsumenRepository) *konsumenService {
	return &konsumenService{konsumenRepository: konsumenRepo}
}
func (s *konsumenService) CreateKonsumen(konsumen *entity.Konsumens) (*entity.Konsumens, error) {
	return s.konsumenRepository.CreateKonsumen(konsumen)
}
func (s *konsumenService) CreateKonsumenLimit(konsumenlimit *entity.Tenor_limits) (*entity.Tenor_limits, error) {
	return s.konsumenRepository.CreateKonsumenLimit(konsumenlimit)
}

func (s *konsumenService) FindNIKKonsumen(nik string) (*entity.Konsumens, error) {
	return s.konsumenRepository.FindNIKKonsumen(nik)
}
func (s *konsumenService) CheckLimitTenor(kode_konsumen string) (*entity.Konsumens, error) {
	return s.konsumenRepository.CheckLimitTenor(kode_konsumen)
}
func (s *konsumenService) FindLimitTenor(kode_konsumen string) (*entity.Konsumens, error) {
	return s.konsumenRepository.FindLimitTenor(kode_konsumen)
}
