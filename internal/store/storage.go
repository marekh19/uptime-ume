package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("record not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Monitors interface {
		Create(context.Context, *Monitor) error
		GetByID(context.Context, string) (*Monitor, error)
		List(context.Context) ([]*Monitor, error)
		Delete(context.Context, string) error
		Update(context.Context, *Monitor) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	PingResults interface {
		Create(context.Context, *PingResult) error
	}
	StatusPages interface {
		Create(context.Context, *StatusPage) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Monitors:    &MonitorStore{db},
		Users:       &UsersStore{db},
		PingResults: &PingResultStore{db},
		StatusPages: &StatusPagesStore{db},
	}
}
