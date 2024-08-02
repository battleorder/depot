package db

import (
	"github.com/google/uuid"
	"github.com/supabase-community/postgrest-go"
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

func GetLowestRank(client *supabase.Client, unitId string) (*Rank, error) {
	var rank Rank
	_, err := client.From(ranksTable).
		Select("*", "exact", false).
		Eq("unitid", unitId).
		Order("rankorder", &postgrest.OrderOpts{Ascending: true}).
		Limit(1, "").
		Single().
		ExecuteTo(&rank)
	if err != nil {
		return nil, err
	}
	return &rank, nil
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
