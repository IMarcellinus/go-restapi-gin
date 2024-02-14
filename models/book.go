package models

type Book struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	NamaBuku  string `gorm:"type:varchar(300)" json:"nama_buku"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}
