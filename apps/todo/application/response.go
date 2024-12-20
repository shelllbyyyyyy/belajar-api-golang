package application

import "time"

type TodoResponse struct {
	Id       	string 		`json:"id"`
	UserId      string 		`json:"user_id"`
	Name 		string		`json:"name"`
	Status    	string		`json:"status"`
	IsArchived  bool		`json:"is_archived"`
	CreatedAt 	time.Time 	`json:"created_at"`
	StartedAt 	time.Time 	`json:"started_at"`
	PausedAt 	time.Time 	`json:"paused_at"`
	FinishedAt 	time.Time 	`json:"finished_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}