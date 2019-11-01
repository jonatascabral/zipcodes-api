package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Code string
	Address string
	Number string
	City string
	State string
	Country string
}

func (z Address) New(
	code string,
	address string,
	number string,
	city string,
	state string,
	country string) *Address {
	return &Address{
		Code: code,
		Address: address,
		Number: number,
		City: city,
		State: state,
		Country: country}
}