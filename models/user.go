package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Location string `json:"location" form:"location" gorm:"type:text"`
	Role     string `json:"role" form:"role" gorm:"type: varchar(255)"`
	Image    string `json:"image"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
