package models

import (
	"time"
)

//User struct
type User struct {
	Firstname   string
	Lastname    string
	Phonenumber string
	password    string
	Username    string
	Createdat   time.Time
}
