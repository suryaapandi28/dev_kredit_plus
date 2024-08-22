package repository

import (
	"errors"
	"time"

	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/cache"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Registered(Accounts *entity.Accounts) (*entity.Accounts, error)
	Verification_account(verifycode *entity.Accounts) (*entity.Accounts, error)
	FindAccountByID(kode_account string) (*entity.Accounts, error)
	UpdateAccountJwtToken(kode_account string, token string, expiresAt time.Time) error
}
type accountRepository struct {
	db        *gorm.DB
	cacheable cache.Cacheable
}

func NewAccountRepository(db *gorm.DB, cacheable cache.Cacheable) *accountRepository {
	return &accountRepository{db: db, cacheable: cacheable}
}
func (r *accountRepository) Registered(Accounts *entity.Accounts) (*entity.Accounts, error) {

	if err := r.db.Create(&Accounts).Error; err != nil {
		return Accounts, err
	}
	return Accounts, nil
}
func (r *accountRepository) Verification_account(verifycode *entity.Accounts) (*entity.Accounts, error) {
	fields := make(map[string]interface{})

	if verifycode.Verification_code != "" {
		fields["verification_code"] = " "
	}
	if verifycode.Verification_expired.Format("2006-01-02 15:04:05") != "" {
		fields["verification_expired"] = verifycode.Verification_expired
	}
	if verifycode.Verification_status != "" {
		fields["verification_status"] = verifycode.Verification_status
	}
	if verifycode.Status_account != "" {
		fields["status_account"] = verifycode.Status_account
	}

	if err := r.db.Model(verifycode).Where("kode_account = ?", verifycode.Kode_account).Updates(fields).Error; err != nil {
		return verifycode, err
	}

	return verifycode, nil
}

func (r *accountRepository) FindAccountByID(kode_account string) (*entity.Accounts, error) {
	account := new(entity.Accounts)
	if err := r.db.Where("kode_account = ?", kode_account).Take(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (u *accountRepository) UpdateAccountJwtToken(kode_account string, token string, expiresAt time.Time) error {
	result := u.db.Model(&entity.Accounts{}).
		Where("kode_account = ? AND deleted_at IS NULL", kode_account).
		Updates(map[string]interface{}{
			"jwt_token":            token,
			"jwt_token_expired_at": expiresAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found or already deleted")
	}

	return nil
}
