package api

import (
	"log"

	"github.com/battleorder/depot/units/internal/db"
	"github.com/gofiber/fiber/v2"
	storage_go "github.com/supabase-community/storage-go"
)

type UnitSummary struct {
	db.Unit
	MemberCount int64 `json:"membercount"`
}

func ListUnits(c *fiber.Ctx) error {
	units, err := db.ListUnits(c.UserContext())
	if err != nil {
		return err
	}

	apiUnits := []UnitSummary{}
	for _, unit := range units {
		if unit.Avatar != nil {
			avatarUrl := db.Client.Storage.GetPublicUrl("units_avatars", *unit.Avatar, storage_go.UrlOptions{
				Transform: &storage_go.TransformOptions{
					Width:  64,
					Height: 64,
				},
			}).SignedURL
			unit.Avatar = &avatarUrl
		}

		mc, err := db.GetMemberCount(c.UserContext(), unit.Id.String())
		if err != nil {
			return err
		}

		us := UnitSummary{
			Unit:        unit,
			MemberCount: mc,
		}
		apiUnits = append(apiUnits, us)
	}

	return c.JSON(apiUnits)
}

type CreateUnitRequest struct {
	Slug        string `json:"slug"`
	DisplayName string `json:"displayName"`
	Description string `json:"description,omitempty"`
	Tagline     string `json:"tagline,omitempty"`
}

type CreateUnitResponse struct {
	Unit   *db.Unit   `json:"unit"`
	Member *db.Member `json:"member"`
	Ranks  []db.Rank  `json:"ranks"`
}

func CreateUnit(c *fiber.Ctx) error {
	user, _ := GetAuthUser(c)

	var body CreateUnitRequest
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	unit, member, ranks, err := db.CreateUnit(
		c.UserContext(),
		user.ID.String(),
		body.Slug,
		body.DisplayName,
		body.Description,
		body.Tagline,
	)
	if err != nil {
		log.Printf("Failed to create unit: %v", err)
		return err
	}

	res := CreateUnitResponse{
		Unit:   unit,
		Member: member,
		Ranks:  ranks,
	}

	return c.JSON(res)
}

func GetUnit(c *fiber.Ctx) error {
	unitId := c.Params("unitId")

	unit, err := db.GetUnit(c.UserContext(), unitId)
	if err != nil {
		return err
	}

	if unit == nil {
		return fiber.ErrNotFound
	}

	if unit.Avatar != nil {
		avatarUrl := db.Client.Storage.GetPublicUrl("units_avatars", *unit.Avatar, storage_go.UrlOptions{
			Transform: &storage_go.TransformOptions{
				Width:  64,
				Height: 64,
			},
		}).SignedURL
		unit.Avatar = &avatarUrl
	}

	mc, err := db.GetMemberCount(c.UserContext(), unit.Id.String())
	if err != nil {
		return err
	}

	us := UnitSummary{
		Unit:        *unit,
		MemberCount: mc,
	}

	return c.JSON(us)
}
