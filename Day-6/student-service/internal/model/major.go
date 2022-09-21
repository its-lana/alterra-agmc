package model

type Major struct {
	Name string `json:"name" gorm:"varchar;not_null;unique"`
	Common
}
