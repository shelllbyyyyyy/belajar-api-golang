package domain

import "time"

type CreateToDoPayload struct {
	UserId string
	Name   string
}

type UpdateToDoPayload struct {
	Name       *string
	Status     *Status
	StartedAt  *time.Time
	PausedAt   *time.Time
	FinishedAt *time.Time
	IsArchived *bool
	UpdatedAt  time.Time
	Id		   string
}
