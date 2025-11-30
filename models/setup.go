package models

// Perhatikan: Nama Struct harus Huruf Besar biar bisa dipanggil dari luar
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `gorm:"type:bit" json:"is_done"`
	UserID      uint   `json:"user_id"`
}