package db

import (
	"log"

	"github.com/supabase-community/supabase-go"
)

const unitsTable = "units"

type Unit struct {
	TraceableModel
	Slug        string  `json:"slug"`
	DisplayName string  `json:"displayname"`
	Tagline     *string `json:"tagline"`
	Description *string `json:"description"`
	Avatar      *string `json:"avatar"`
}

func GetUnit(client *supabase.Client, unitid string) (*Unit, error) {
	var unit Unit
	rows, err := client.From(unitsTable).
		Select("*", "exact", false).
		Eq("id", unitid).
		Single().
		ExecuteTo(&unit)
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, nil
	}

	return &unit, nil
}

func ListUnits(client *supabase.Client, perPage int, nextToken ...string) (Paginated[Unit], error) {
	units := []Unit{}

  p := Paginated[Unit]{
    Data: units,
    PerPage: perPage,
  }

	_, err := Paginate(
    client.From(unitsTable).Select("*", "exact", false),
    "id",
    perPage,
    nextToken...,
  ).ExecuteTo(&units)
	if err != nil {
		return p, err
	}

  p.Data = units
  if len(units) > 0 {
    next := units[len(units)-1].Id.String()
    p.NextToken = &next
  }

	return p, nil
}

type createUnitBody struct {
	Slug        string `json:"slug"`
	DisplayName string `json:"displayname"`
	Description string `json:"description,omitempty"`
	Tagline     string `json:"tagline,omitempty"`
	OwnerId     string `json:"ownerid"`
}

func CreateUnit(client *supabase.Client, userId, slug, displayName, description, tagline string) (*Unit, *Member, []Rank, error) {
	defaultRanks := []createRankBody{
		{Slug: "Pvt", DisplayName: "Private", RankOrder: 0},
		{Slug: "Cpl", DisplayName: "Corporal", RankOrder: 1},
		{Slug: "Sgt", DisplayName: "Sergeant", RankOrder: 2},
		{Slug: "Lt", DisplayName: "Lieutenant", RankOrder: 3},
		{Slug: "Cmdr", DisplayName: "Commander", RankOrder: 4},
	}

	var units []Unit
	_, err := client.From(unitsTable).Insert(&createUnitBody{
		Slug:        slug,
		DisplayName: displayName,
		Description: description,
		Tagline:     tagline,
		OwnerId:     userId,
	}, false, "", "representation", "exact").ExecuteTo(&units)
	if err != nil {
		log.Printf("Failed to insert unit: %v", err)
		return nil, nil, []Rank{}, err
	}
	unit := &units[0]

	var endRanks []Rank
	for _, r := range defaultRanks {
		rank, err := CreateRank(
			client,
			unit.Id.String(),
			r.Slug,
			r.DisplayName,
			r.RankOrder,
		)
		if err != nil {
			return nil, nil, []Rank{}, err
		}
		endRanks = append(endRanks, *rank)
	}

	member, err := CreateMember(
		client,
		unit.Id.String(),
		userId,
		endRanks[len(endRanks)-1].Id.String(),
		"",
		true,
	)
	if err != nil {
		return nil, nil, []Rank{}, err
	}

	return unit, member, endRanks, nil
}
