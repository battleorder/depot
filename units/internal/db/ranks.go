package db

import (
	"context"

	"github.com/google/uuid"
)

const ranksTable = "ranks"

type Rank struct {
	TraceableModel
	UnitID      uuid.UUID `json:"unitid"`
	Slug        string    `json:"slug"`
	DisplayName string    `json:"displayname"`
	RankOrder   int       `json:"rankorder"`
}

type createRankBody struct {
	UnitID      string `json:"unitid"`
	Slug        string `json:"slug"`
	DisplayName string `json:"displayname"`
	RankOrder   int    `json:"rankorder"`
}

func CreateRank(ctx context.Context, unitId, slug, displayName string, rankOrder int) (*Rank, error) {
	body := createRankBody{
		UnitID:      unitId,
		Slug:        slug,
		DisplayName: displayName,
		RankOrder:   rankOrder,
	}

	var ranks []Rank
	_, err := Client.From(ranksTable).Insert(&body, false, "", "representation", "exact").ExecuteTo(&ranks)
	if err != nil {
		return nil, err
	}

	return &ranks[0], nil
}
