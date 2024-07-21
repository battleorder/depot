package db

import (
	"context"
	"log"
)

const unitsTable = "units"

type Unit struct {
	TraceableModel
	Slug        string `json:"slug"`
	DisplayName string `json:"displayname"`
	Tagline     *string `json:"tagline"`
	Description *string `json:"description"`
	Avatar      *string `json:"avatar"`
}

func GetUnit(ctx context.Context, unitid string) (*Unit, error) {
	var unit Unit
	rows, err := Client.From(unitsTable).
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

func ListUnits(ctx context.Context) ([]Unit, error) {
	units := []Unit{}
	_, err := Client.From(unitsTable).
		Select("*", "exact", false).
		ExecuteTo(&units)
	if err != nil {
		return nil, err
	}

	return units, nil
}

type createUnitBody struct {
	Slug        string `json:"slug"`
	DisplayName string `json:"displayname"`
	Description string `json:"description,omitempty"`
	Tagline     string `json:"tagline,omitempty"`
}

func CreateUnit(ctx context.Context, userId, slug, displayName, description, tagline string) (*Unit, *Member, []Rank, error) {
	defaultRanks := []createRankBody{
		{Slug: "Pvt", DisplayName: "Private", RankOrder: 0},
		{Slug: "Cpl", DisplayName: "Corporal", RankOrder: 1},
		{Slug: "Sgt", DisplayName: "Sergeant", RankOrder: 2},
		{Slug: "Lt", DisplayName: "Lieutenant", RankOrder: 3},
		{Slug: "Cmdr", DisplayName: "Commander", RankOrder: 4},
	}

	var units []Unit
	_, err := Client.From(unitsTable).Insert(&createUnitBody{
		Slug:        slug,
		DisplayName: displayName,
		Description: description,
		Tagline:     tagline,
	}, false, "", "representation", "exact").ExecuteTo(&units)
	if err != nil {
		log.Printf("Failed to insert unit: %v", err)
		return nil, nil, []Rank{}, err
	}
	unit := &units[0]

	var endRanks []Rank
	for _, r := range defaultRanks {
		rank, err := CreateRank(
			ctx,
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
    ctx,
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
