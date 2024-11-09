package models

import "time"

type Tb_User struct {
	Usr_id int `json:"usr_id"`
	Usr_name string `json:"usr_name"`
	Usr_email string `json:"usr_email"`
	Usr_password string `json:"usr_password"`
}

type Tb_Task struct {
	Tsk_id            int       `json:"tsk_id"`
	Tsk_name          string    `json:"tsk_name"`
	Tsk_description   string    `json:"tsk_description"`
	Tsk_creation_date time.Time `json:"tsk_creation_date"`
	Tsk_update_date   time.Time `json:"task_update_date"`
	Tsk_deadline_date time.Time `json:"taks_deadline_date"`
	Tsk_color         string    `json:"tsk_color"`
	Tskpr_id int `json:"tskpr_id"`
	Tskst_id int `json:"tsksk_id"`
}

type Tb_Category struct {
	cat_id   int    `json:"cat_id"`
	cat_name string `json:"cat_name"`
	usr_id   int    `json:"usr_id"`
}

type Tb_Task_category struct {
	tsk_id int `json:"tsk_id"`
	cat_id int `json:"cat_id"`
}

type Tb_Task_status struct {
	tskst_id int `json:"tskst_id"`
	tskst_name string `json:"tskst_name"`
}

type Tb_Task_priority struct {
	tskpr_id int `json:"tskpr_i"`
	tskpr_name string `json:"tskpr_name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}