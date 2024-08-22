package entity

import "time"

type Accounts struct {
	Kode_account         string `json:"kode_account"`
	Kode_konsumen        string `json:"kode_konsumen"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	Status_account       string `json:"status_account"`
	Verification_status  string `json:"verification_status"`
	Verification_code    string `json:"verification_code"`
	Verification_expired time.Time
	Jwt_token            string `json:"jwt_token"`
	Jwt_token_expired_at time.Time
	Device_id            string `json:"device_id"`
	Auditable
}
