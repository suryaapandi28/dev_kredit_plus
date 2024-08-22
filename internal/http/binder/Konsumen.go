package binder

type KonsumenCreateRequest struct {
	Kode_konsumen string `form:"kode_konsumen" validate:"required"`
	NIK           string `form:"nik" validate:"required"`
	Full_name     string `form:"full_name" validate:"required"`
	Legal_name    string `form:"legal_name" validate:"required"`
	Tempat_lahir  string `form:"tempat_lahir"  validate:"required"`
	Tanggal_lahir string `form:"tanggal_lahir"  validate:"required"`
	Gaji          string `form:"gaji"  validate:"required"`
	Foto_ktp      string `form:"foto_ktp"  validate:"required"`
	Foto_selfie   string `form:"foto_selfie"  validate:"required"`
}
