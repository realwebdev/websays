package models

type User struct {
	Projects  []Project  `json:"-" gorm:"foreignKey:UserID"`
	Activitys []Activity `json:"-" gorm:"foreignKey:UserID"`
	ID        uint
	Username  string `json:"username"`
	Email     string `gorm:"typevarchar(100);unique_index" json:"email"`
	Password  string `json:"-"`
}
