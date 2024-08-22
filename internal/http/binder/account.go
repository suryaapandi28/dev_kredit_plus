package binder

type AccountreateRequest struct {
	Kode_account         string `form:"kode_account" validate:"required"`
	Kode_konsumen        string `form:"kode_konsumen" validate:"required"`
	Email                string `form:"email" validate:"required"`
	Password             string `form:"password" validate:"required"`
	Status_account       string `form:"status_account" validate:"required"`
	Verification_status  string `form:"verification_status" validate:"required"`
	Verification_code    string `form:"verification_code" validate:"required"`
	Verification_expired string `form:"verification_expired" validate:"required"`
	Jwt_token            string `form:"jwt_token" validate:"required"`
	Jwt_token_expired_at string `form:"jwt_token_expired_at" validate:"required"`
	Device_id            string `form:"device_id" validate:"required"`
}
