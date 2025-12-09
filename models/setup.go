package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"default:'user'" json:"role"`
}

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline 	string `json:"deadline"`
	IsDone      bool   `gorm:"type:bit" json:"is_done"`
	UserID      uint   `json:"user_id"`
}