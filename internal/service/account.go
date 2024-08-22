package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"
	"github.com/suryaapandi28/dev_kredit_plus/internal/repository"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/email"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/encrypt"
	"github.com/suryaapandi28/dev_kredit_plus/pkg/token"
)

type AccountService interface {
	Registered(Accounts *entity.Accounts) (*entity.Accounts, error)
	Verification_account(verifycode *entity.Accounts) (*entity.Accounts, error)
	FindAccountByID(kode_account string) (*entity.Accounts, error)
	LoginAccount(kode_account string, email string, password string) (string, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
	tokenUseCase      token.TokenUseCase
	encryptTool       encrypt.EncryptTool
	emailSender       *email.EmailSender
}

func NewAccountService(accountRepo repository.AccountRepository, tokenUseCase token.TokenUseCase,
	encryptTool encrypt.EncryptTool, emailSender *email.EmailSender) *accountService {
	return &accountService{accountRepository: accountRepo, tokenUseCase: tokenUseCase, encryptTool: encryptTool, emailSender: emailSender}
}
func (s *accountService) Registered(Accounts *entity.Accounts) (*entity.Accounts, error) {
	emailAddr := Accounts.Email
	err := s.emailSender.SendWelcomeEmail(emailAddr, Accounts.Email, "")

	if err != nil {
		return nil, err
	}

	err = s.emailSender.SendVerifyAccount(Accounts.Email, Accounts.Verification_code)
	if err != nil {
		return nil, err
	}
	Newaccount, err := s.accountRepository.Registered(Accounts)
	if err != nil {
		return nil, err
	}
	return Newaccount, nil
}
func (s *accountService) Verification_account(verifycode *entity.Accounts) (*entity.Accounts, error) {
	return s.accountRepository.Verification_account(verifycode)
}

func (s *accountService) FindAccountByID(kode_account string) (*entity.Accounts, error) {
	return s.accountRepository.FindAccountByID(kode_account)
}
func (s *accountService) LoginAccount(kode_account string, email string, password string) (string, error) {
	account, err := s.accountRepository.FindAccountByID(kode_account)
	if err != nil {
		return "", errors.New("wrong input email/password")
	}

	// Lanjutkan dengan pembuatan token dan logika lainnya
	expiredTime := time.Now().Local().Add(24 * time.Hour)

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	expiredTimeInJakarta := expiredTime.In(location)
	// Buat claims JWT
	claims := token.JwtCustomClaims{
		ID:    account.Kode_account,
		Email: account.Email,
		Role:  "user",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "kreditplus",
			ExpiresAt: jwt.NewNumericDate(expiredTimeInJakarta),
		},
	}

	// Generate JWT token
	jwtToken, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("there is an error in the system")
	}

	// nyimpen token JWT dan waktu ke database
	account.Jwt_token = jwtToken
	account.Jwt_token_expired_at = expiredTime

	// update token JWT dan waktu di database
	if err := s.accountRepository.UpdateAccountJwtToken(kode_account, jwtToken, expiredTime); err != nil {
		return "", errors.New("failed to update user token info")
	}

	// double check buat bandingin JWT yang dibuat dengan JWT yang tersimpan di database
	if account.Jwt_token != jwtToken {
		return "", errors.New("JWT token mismatch")
	}

	return jwtToken, nil
}
