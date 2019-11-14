package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Code string
	Address string
	District string
	Number string
	City string
	State string
	Country string
}

func (z Address) New(
	code string,
	address string,
	district string,
	number string,
	city string,
	state string,
	country string) *Address {
	return &Address{
		Code: code,
		Address: address,
		Number: number,
		District: district,
		City: city,
		State: state,
		Country: country}
}