package model

import "time"

type Employee struct {
	Id         uint64    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Username   string    `json:"username" gorm:"unique"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IdNumber   string    `json:"idNumber"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}
