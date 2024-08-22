package entity

type Konsumens struct {
	Kode_konsumen string `json:"kode_konsumen"`
	NIK           string `json:"nik"`
	Full_name     string `json:"full_name"`
	Legal_name    string `json:"legal_name"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gaji          string `json:"gaji"`
	Foto_ktp      string `json:"foto_ktp"`
	Foto_selfie   string `json:"foto_selfie"`
	Kredit_score  string `json:"kredit_score"`
	Auditable
}

func NewKonsumen(kode_konsumen, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie, kredit_score string) *Konsumens {
	return &Konsumens{
		Kode_konsumen: kode_konsumen,
		NIK:           nik,
		Full_name:     full_name,
		Legal_name:    legal_name,
		Tempat_lahir:  tempat_lahir,
		Tanggal_lahir: tanggal_lahir,
		Gaji:          gaji,
		Foto_ktp:      foto_ktp,
		Foto_selfie:   foto_selfie,
		Kredit_score:  kredit_score,
		Auditable:     UpdateAuditable(),
	}
}

type Tenor_limits struct {
	Kode_tenor    string `json:"kode_tenor"`
	Kode_konsumen string `json:"kode_konsumen"`
	Nama_konsumen string `json:"nama_konsumen"`
	Tenor_pertama string `json:"tenor_pertama"`
	Tenor_Kedua   string `json:"tenor_kedua"`
	Tenor_Ketiga  string `json:"tenor_ketiga"`
	Tenor_keenam  string `json:"tenor_keenam"`
	Auditable
}

func NewKonsumenLimit(kode_tenor, kode_konsumen, nama_konsumen, tenor_pertama, tenor_kedua, tenor_ketiga, tenor_keenam string) *Tenor_limits {
	return &Tenor_limits{
		Kode_tenor:    kode_tenor,
		Kode_konsumen: kode_konsumen,
		Nama_konsumen: nama_konsumen,
		Tenor_pertama: tenor_pertama,
		Tenor_Kedua:   tenor_kedua,
		Tenor_Ketiga:  tenor_ketiga,
		Tenor_keenam:  tenor_keenam,
		Auditable:     UpdateAuditable(),
	}
}
