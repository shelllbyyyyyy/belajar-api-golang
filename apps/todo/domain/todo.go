package domain

import (
	"api/first-go/common"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Pending 	Status = "pending"
	Working 	Status = "working"
	Paused 		Status = "paused"
	Completed 	Status = "completed"
)

type Todo struct {
	Id        	string    `db:"id"`
	UserId		string 	  `db:"user_id"`
	Name      	string 	  `db:"name"`
	Status 		Status    `db:"status"`
	IsArchived 	bool 	  `db:"is_archived"`
	CreatedAt 	time.Time `db:"created_at"`
	StartedAt 	*time.Time `db:"started_at"`
	PausedAt 	*time.Time `db:"paused_at"`
	FinishedAt 	*time.Time `db:"finished_at"`
	UpdatedAt 	time.Time `db:"updated_at"`
}

func NewTodo(payload CreateToDoPayload) (*Todo, error) {
	if payload.UserId == "" || payload.Name == "" {
		return nil, common.ErrInputCannotBeNull
	}

	return &Todo{
		Id: uuid.NewString(),
		UserId: payload.UserId,
		Name: payload.Name,
		Status: Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (t *Todo) Update(name string) (err error) {
	if name == "" {
		return common.ErrInputCannotBeNull
	}

	t.Name = name
	t.UpdatedAt = time.Now()

	return
}

func (t *Todo) Validate() (err error) {
	if len(t.Name) < 4 || len(t.Name) > 100 {
		err = common.ErrToDoNameInvalidLength
		return 
	}

	return
}

func (t *Todo) Complete() (err error) {
	if t.IsArchived {
		return common.ErrToDoIsArchived
	}

	now := time.Now()

	t.Status = Completed
	t.FinishedAt = &now
	t.UpdatedAt = time.Now()

	return
}

func (t *Todo) Paused() (err error) {
	if t.IsArchived {
		return common.ErrToDoIsArchived
	}

	now := time.Now()

	t.Status = Paused
	t.PausedAt = &now
	t.UpdatedAt = time.Now()

	return
}

func (t *Todo) Working() (err error) {
	if t.IsArchived {
		return common.ErrToDoIsArchived
	}

	now := time.Now()
	
	t.Status = Working
	t.StartedAt = &now
	t.UpdatedAt = time.Now()

	return
}

func (t *Todo) Archived() (err error) {
	t.IsArchived = true
	t.UpdatedAt = time.Now()

	return
}

func (t *Todo) UnArchived() (err error) {
	t.IsArchived = false
	t.UpdatedAt = time.Now()

	return
}