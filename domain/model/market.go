package model

type Market struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	Code string `gorm:"unique_index;not_null;" json:"code"`
	FirstName string `json:"first_name"`
}
