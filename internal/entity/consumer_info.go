package entity

import "time"

type ConsumerInfo struct {
	Id           int       `json:"id,omitempty"`
	NIK          string    `json:"nik,omitempty"`
	Email        string    `json:"email,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	FullName     string    `json:"full_name,omitempty"`
	LegalName    string    `json:"legal_name,omitempty"`
	TempatLahir  string    `json:"tempat_lahir,omitempty"`
	TanggalLahir string    `json:"tanggal_lahir,omitempty"`
	Gaji         uint      `json:"gaji,omitempty"`
	FotoKTP      string    `json:"foto_ktp,omitempty"`
	FotoSelfie   string    `json:"foto_selfie,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type ConsumerInfoRequest struct {
	NIK          string `json:"nik,omitempty"`
	Email        string `json:"email,omitempty"`
	Gender       string `json:"gender,omitempty"`
	FullName     string `json:"full_name,omitempty"`
	LegalName    string `json:"legal_name,omitempty"`
	TempatLahir  string `json:"tempat_lahir,omitempty"`
	TanggalLahir string `json:"tanggal_lahir,omitempty"`
	Gaji         uint   `json:"gaji,omitempty"`
	FotoKTP      string `json:"foto_ktp,omitempty"`
	FotoSelfie   string `json:"foto_selfie,omitempty"`
}
