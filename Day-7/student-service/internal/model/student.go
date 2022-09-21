package model

type Student struct {
	Fullname string `json:"fullname" gorm:"varchar;not_null"`
	Email    string `json:"email" gorm:"varchar;not_null;unique"`
	Password string `json:"password" gorm:"varchar;not_null"`
	ClassID  uint   `json:"class_id"`
	Class    Class
	MajorID  uint `json:"major_id"`
	Major    Major
	Common
}
