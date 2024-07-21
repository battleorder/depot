package db

import (
	"time"

	"github.com/google/uuid"
)

type TraceableModel struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdat,omitempty"`
	UpdatedAt time.Time `json:"updatedat,omitempty"`
}
