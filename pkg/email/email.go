package email

import (
	"fmt"

	"github.com/suryaapandi28/dev_kredit_plus/internal/entity"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	Config *entity.Config
}

func NewEmailSender(config *entity.Config) *EmailSender {
	return &EmailSender{Config: config}
}

func (e *EmailSender) SendEmail(to []string, subject, body string) error {
	from := "kreditplus@smartlms.my.id"
	password := e.Config.SMTP.Password
	smtpHost := e.Config.SMTP.Host

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	dialer := gomail.NewDialer(smtpHost, 587, from, password)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func (e *EmailSender) SendWelcomeEmail(to, name, message string) error {
	subject := "Selamat Datang | Kredit Plus"
	body := fmt.Sprintf("Dear %s,\nTerima Kasih Telah mendaftarkan Di Kredit Plus\n\nKREDIT PLUS\n'Solusi Pembiayaan Terpercaya Untuk Semua Kebutuhan Anda'", name)
	return e.SendEmail([]string{to}, subject, body)
}

func (e *EmailSender) SendResetPasswordEmail(to, name, resetCode string) error {
	subject := "Reset Password | Depublic"
	body := fmt.Sprintf("Dear %s,\nPlease use the following code to reset your password: %s\n\nDepublic Team", name, resetCode)
	return e.SendEmail([]string{to}, subject, body)
}

func (e *EmailSender) SendVerificationEmail(to, name, code string) error {
	subject := "Verify Your Email | Depublic"
	body := fmt.Sprintf("Dear %s,\nPlease use the following code to verify your email: %s\n\nDepublic Team", name, code)
	return e.SendEmail([]string{to}, subject, body)
}

func (e *EmailSender) SendVerifyAccount(to, verifycode string) error {
	subject := "Verify Your Email | Kredit Plus"
	body := fmt.Sprintf("Hallo \nTerima Kasih telah mendaftar di Kredit Plus\n\nUntuk mengaktifkan akun anda dan memulai menggunakan fitur kami. Silahkan masukkan kode veritifikasi dibawah ini :\n\nKode : %s\n\nPenting : Tautan ini hanya berlaku selama 5 Menit. Jika Anda tidak meminta verifikasi ini, abaikan email ini.\n\nSalam,\nTim Kredit Plus",
		verifycode)
	return e.SendEmail([]string{to}, subject, body)
}
