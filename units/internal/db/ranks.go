package db

import (
	"github.com/google/uuid"
	"github.com/supabase-community/supabase-go"
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

func CreateRank(client *supabase.Client, unitId, slug, displayName string, rankOrder int) (*Rank, error) {
	body := createRankBody{
		UnitID:      unitId,
		Slug:        slug,
		DisplayName: displayName,
		RankOrder:   rankOrder,
	}

	var ranks []Rank
	_, err := client.From(ranksTable).Insert(&body, false, "", "representation", "exact").ExecuteTo(&ranks)
	if err != nil {
		return nil, err
	}

	return &ranks[0], nil
}
