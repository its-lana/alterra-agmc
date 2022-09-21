package model

type Class struct {
	Name string `json:"name" gorm:"varchar;not_null;unique"`
	Common
}
