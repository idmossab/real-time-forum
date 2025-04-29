package server

import(

)

import "real_time_forum/internal/model"

var TestUser = &model.User{
	Id:        1,
	NickName:  "devGuru42",
	Age:       28,
	Gender:    "Male",
	FirstName: "Alex",
	LastName:  "Johnson",
	Email:     "alex.johnson@example.com",
	Password:  "securePassword123!",
}



func init(){

}