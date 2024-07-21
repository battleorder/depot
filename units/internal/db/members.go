package db

import (
	"context"

	"github.com/google/uuid"
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

func CreateMember(ctx context.Context, unitId, userId, rankId, displayName string, isAdmin bool) (*Member, error) {
	body := createMemberBody{
		UnitID:      unitId,
		UserID:      userId,
		RankID:      rankId,
		DisplayName: displayName,
		IsAdmin:     isAdmin,
	}

	var members []Member
	_, err := Client.From(membersTable).Insert(&body, false, "", "representation", "exact").ExecuteTo(&members)
	if err != nil {
		return nil, err
	}

	return &members[0], nil
}

func GetMemberCount(ctx context.Context, unitId string) (int64, error) {
	_, count, err := Client.From(membersTable).
		Select("id", "exact", false).
		Eq("unitid", unitId).
		Execute()
	if err != nil {
		return 0, err
	}

	return count, nil
}
