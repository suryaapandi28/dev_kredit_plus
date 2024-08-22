package entity

type Transaction_headers struct {
	Kode_transaction string `json:"kode_transaction"`
	Kode_konsumen    string `json:"kode_konsumen"`
	OTR              string `json:"otr"`
	Admin_fee        string `json:"admin_fee"`
	Jumlah_cicilan   string `json:"jumlah_cicilan" `
	Jumlah_bunga     string `json:"jumlah_bunga" `
	Nama_aset        string `json:"nama_aset" `
	Auditable
}
