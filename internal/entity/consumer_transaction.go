package entity

import "time"

type ConsumerTransaction struct {
	Id            int       `json:"id,omitempty"`
	ConsumerId    int       `json:"consumer_id,omitempty"`
	Otr           int       `json:"otr,omitempty"`
	AdminFee      int       `json:"admin_fee,omitempty"`
	JumlahCicilan int       `json:"jumlah_cicilan,omitempty"`
	JumlahBunga   float64   `json:"jumlah_bunga,omitempty"`
	NamaAsset     string    `json:"nama_asset,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
