package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/supabase-community/postgrest-go"
)

type TraceableModel struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdat,omitempty"`
	UpdatedAt time.Time `json:"updatedat,omitempty"`
}

type Paginated[T any] struct {
	Data      []T     `json:"data"`
	NextToken *string `json:"next_token"`
	PerPage   int     `json:"per_page"`
}

func Paginate(qb *postgrest.FilterBuilder, field string, perPage int, token ...string) *postgrest.FilterBuilder {
	b := qb.Order("id", &postgrest.OrderOpts{Ascending: true}).Limit(perPage+1, "")

	if len(token) > 0 {
		nextToken := token[0]
		b = b.Gt("id", nextToken)
	}

	return b
}
