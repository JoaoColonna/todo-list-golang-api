package models

import "time"

type Tb_Task struct {
	tsk_id            int       `json:"tsk_id"`
	tsk_name          string    `json:"tsk_name"`
	tsk_description   string    `json:"tsk_description"`
	tsk_creation_date time.Time `json:"tsk_creation_date"`
	tsk_update_date   time.Time `json:"task_update_date"`
	tsk_deadline_date time.Time `json:"taks_deadline_date"`
	tsk_status        string    `json:"tsk_status"`
	tsk_color         string    `json:"tsk_color"`
	// usr_id id `json:"usr_id"`
	tskpr_id int `json:"tskpr_id"`
}

type Tb_Category struct {
	cat_id   int    `json:"cat_id"`
	cat_name string `json:"cat_name"`
	usr_id   int    `json:"usr_id"`
}

type task_category struct {
	tsk_id int `json:"tsk_id"`
	cat_id int `json:"cat_id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
