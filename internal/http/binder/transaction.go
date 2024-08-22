package binder

type TransactionCreateRequest struct {
	Kode_transaction string `form:"kode_transaction" validate:"required"`
	Kode_konsumen    string `form:"kode_konsumen" validate:"required"`
	OTR              string `form:"otr" validate:"required"`
	Admin_fee        string `form:"admin_fee" validate:"required"`
	Jumlah_cicilan   string `form:"jumlah_cicilan"  validate:"required"`
	Jumlah_bunga     string `form:"jumlah_bunga"  validate:"required"`
	Nama_aset        string `form:"nama_aset"  validate:"required"`
}
