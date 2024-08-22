package repository

import (
	"encoding/base64"
	"encoding/json"

	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/cache"
	"gorm.io/gorm"
)

type KonsumenRepository interface {
	CreateKonsumen(Konsumens *entity.Konsumens) (*entity.Konsumens, error)
	CreateKonsumenLimit(konsumenlimit *entity.Tenor_limits) (*entity.Tenor_limits, error)
	FindNIKKonsumen(nik string) (*entity.Konsumens, error)
	CheckLimitTenor(kode_konsumen string) (*entity.Konsumens, error)
	FindLimitTenor(kode_konsumen string) (*entity.Konsumens, error)
}
type konsumenRepository struct {
	db        *gorm.DB
	cacheable cache.Cacheable
}

func NewKonsumenRepository(db *gorm.DB, cacheable cache.Cacheable) KonsumenRepository {
	return &konsumenRepository{db: db, cacheable: cacheable}
}

func (r *konsumenRepository) CreateKonsumen(Konsumens *entity.Konsumens) (*entity.Konsumens, error) {

	if err := r.db.Create(&Konsumens).Error; err != nil {
		return Konsumens, err
	}
	return Konsumens, nil
}
func (r *konsumenRepository) CreateKonsumenLimit(Konsumenslimit *entity.Tenor_limits) (*entity.Tenor_limits, error) {

	if err := r.db.Create(&Konsumenslimit).Error; err != nil {
		return Konsumenslimit, err
	}
	return Konsumenslimit, nil
}

func (r *konsumenRepository) FindNIKKonsumen(nik string) (*entity.Konsumens, error) {
	var konsumens entity.Konsumens
	convertnik := base64.StdEncoding.EncodeToString([]byte(nik))

	konsumensdata := &konsumens
	key := "FindNIK"

	data, _ := r.cacheable.Get(key)

	if data == "" {

		err := r.db.Raw("SELECT * FROM konsumens WHERE NIK = ?", convertnik).Scan(&konsumens).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		return &konsumens, nil

	} else {
		// Data ditemukan di Redis, unmarshal data ke konsumen
		err := json.Unmarshal([]byte(data), &konsumensdata)
		if err != nil {
			return konsumensdata, err
		}
	}
	return &konsumens, nil

}

func (r *konsumenRepository) CheckLimitTenor(kode_konsumen string) (*entity.Konsumens, error) {
	var konsumens entity.Konsumens
	konsumensdata := &konsumens
	key := "FindNIK"

	data, _ := r.cacheable.Get(key)

	if data == "" {

		err := r.db.Raw("SELECT * FROM konsumens WHERE kode_konsumen = ?", kode_konsumen).Scan(&konsumens).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		return &konsumens, nil

	} else {
		// Data ditemukan di Redis, unmarshal data ke konsumen
		err := json.Unmarshal([]byte(data), &konsumensdata)
		if err != nil {
			return konsumensdata, err
		}
	}
	return &konsumens, nil

}

func (r *konsumenRepository) FindLimitTenor(kode_konsumen string) (*entity.Konsumens, error) {
	var konsumens entity.Konsumens
	konsumensdata := &konsumens
	key := "FindNIK"

	data, _ := r.cacheable.Get(key)

	if data == "" {

		err := r.db.Raw("SELECT * FROM tenor_limits WHERE kode_konsumen = ?", kode_konsumen).Scan(&konsumens).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			return nil, err
		}
		return &konsumens, nil

	} else {
		// Data ditemukan di Redis, unmarshal data ke konsumen
		err := json.Unmarshal([]byte(data), &konsumensdata)
		if err != nil {
			return konsumensdata, err
		}
	}
	return &konsumens, nil

}

// func (r *konsumenRepository) UpdateLimit(limitTenor *entity.Tenor_limits) (*entity.Tenor_limits, error) {
// 	fields := make(map[string]interface{})

// 	if limitTenor.Tenor_pertama != "" {
// 		fields["ternor_pertama"] = limitTenor.Tenor_pertama
// 	}
// 	if limitTenor.Tenor_Kedua != "" {
// 		fields["tenor_kedua"] = limitTenor.Tenor_Kedua
// 	}
// 	if limitTenor.Tenor_Ketiga != "" {
// 		fields["tenor_Ketiga"] = limitTenor.Tenor_Ketiga
// 	}
// 	if limitTenor.Tenor_keenam != "" {
// 		fields["tenor_keenam"] = limitTenor.Tenor_keenam
// 	}

// 	if err := r.db.Model(limitTenor).Where("kode_konsumen = ?", limitTenor.Kode_konsumen).Updates(fields).Error; err != nil {
// 		return limitTenor, err
// 	}

// 	return limitTenor, nil
// }

func (r *konsumenRepository) UpdateLimit(limitTenor *entity.Tenor_limits) (*entity.Tenor_limits, error) {
	// Use map to store fields to be updated.
	fields := make(map[string]interface{})

	if limitTenor.Tenor_pertama != "" {
		fields["ternor_pertama"] = limitTenor.Tenor_pertama
	}
	if limitTenor.Tenor_Kedua != "" {
		fields["tenor_kedua"] = limitTenor.Tenor_Kedua
	}
	if limitTenor.Tenor_Ketiga != "" {
		fields["tenor_Ketiga"] = limitTenor.Tenor_Ketiga
	}
	if limitTenor.Tenor_keenam != "" {
		fields["tenor_keenam"] = limitTenor.Tenor_keenam
	}

	// Update the database in one query.
	if err := r.db.Model(limitTenor).Where("kode_konsumen = ?", limitTenor.Kode_konsumen).Updates(fields).Error; err != nil {
		return limitTenor, err
	}

	return limitTenor, nil
}
