package authdto

type CreateUser struct {
	Fullname string `json:"fullname" form:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" form:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Location string `json:"location" form:"location" gorm:"type:text"`
	Role     string `json:"role" form:"role" gorm:"type: varchar(255)"`
	Image    string `json:"image"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
