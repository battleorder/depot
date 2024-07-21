package db

import (
	"github.com/google/uuid"
	"github.com/supabase-community/supabase-go"
)

const membersTable = "members"

type Member struct {
	TraceableModel
	UnitID      uuid.UUID `json:"unitid"`
	UserID      uuid.UUID `json:"userid"`
	RankID      uuid.UUID `json:"rankid"`
	DisplayName string    `json:"displayname"`
	IsAdmin     bool      `json:"isadmin"`
}

type createMemberBody struct {
	UnitID      string `json:"unitid"`
	UserID      string `json:"userid"`
	RankID      string `json:"rankid"`
	DisplayName string `json:"displayname,omitempty"`
	IsAdmin     bool   `json:"isadmin"`
}

func CreateMember(client *supabase.Client, unitId, userId, rankId, displayName string, isAdmin bool) (*Member, error) {
	body := createMemberBody{
		UnitID:      unitId,
		UserID:      userId,
		RankID:      rankId,
		DisplayName: displayName,
		IsAdmin:     isAdmin,
	}

	var members []Member
	_, err := client.From(membersTable).Insert(&body, false, "", "representation", "exact").ExecuteTo(&members)
	if err != nil {
		return nil, err
	}

	return &members[0], nil
}

func GetMemberCount(client *supabase.Client, unitId string) (int64, error) {
	_, count, err := client.From(membersTable).
		Select("id", "exact", false).
		Eq("unitid", unitId).
		Execute()
	if err != nil {
		return 0, err
	}

	return count, nil
}
