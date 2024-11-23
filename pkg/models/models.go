package models

import "time"

type Tb_User struct {
	Usr_id       int    `json:"usr_id"`
	Usr_name     string `json:"usr_name"`
	Usr_email    string `json:"usr_email"`
	Usr_password string `json:"usr_password"`
}

type UserDTO struct {
	Usr_name     string `json:"usr_name"`
	Usr_email    string `json:"usr_email"`
	Usr_password string `json:"usr_password"`
}

type UserResponse struct {
	Usr_id    int    `json:"usr_id"`
	Usr_name  string `json:"usr_name"`
	Usr_email string `json:"usr_email"`
}

type Tb_Task struct {
	Tsk_id            int       `json:"tsk_id"`
	Tsk_name          string    `json:"tsk_name"`
	Tsk_description   string    `json:"tsk_description"`
	Tsk_creation_date time.Time `json:"tsk_creation_date"`
	Tsk_update_date   time.Time `json:"tsk_update_date"`
	Tsk_deadline_date time.Time `json:"tks_deadline_date"`
	Tsk_color         string    `json:"tsk_color"`
	Usr_id            int       `json:"usr_id"`
	Tskpr_id          int       `json:"tskpr_id"`
	Tskst_id          int       `json:"tsksk_id"`
}

type Task_Request struct {
	Tsk_name          string    `json:"tsk_name"`
	Tsk_description   string    `json:"tsk_description"`
	Tsk_deadline_date string 	`json:"tks_deadline_date"`
	Tsk_color         string    `json:"tsk_color"`
	Usr_id            int       `json:"usr_id"`
	Tskpr_id          int       `json:"tskpr_id"`
	Tskst_id          int       `json:"tsksk_id"`
}

type Tb_Category struct {
	Cat_id   int    `json:"cat_id"`
	Cat_name string `json:"cat_name"`
	Usr_id   int    `json:"usr_id"`
}

type Tb_Task_category struct {
	Tsk_id int `json:"tsk_id"`
	Cat_id int `json:"cat_id"`
}

type Tb_Task_status struct {
	Tskst_id   int    `json:"tskst_id"`
	Tskst_name string `json:"tskst_name"`
}

type Tb_Task_priority struct {
	Tskpr_id   int    `json:"tskpr_id"`
	Tskpr_name string `json:"tskpr_name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginResponse struct {
	Token string `json:"token"`
	UserResponse
}