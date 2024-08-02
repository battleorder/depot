package server

import (
	"net/http"
	"strconv"

	"github.com/battleorder/depot/units/internal/db"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	storage_go "github.com/supabase-community/storage-go"
)

func handleGetUnits(lgr log.Logger) http.Handler {
	type request struct{}

	type responseItem struct {
		db.Unit
		MemberCount int64 `json:"membercount"`
	}
	type response db.Paginated[responseItem]

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := GetSupabaseClient(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			level.Error(lgr).Log("msg", "failed to supabase client for user", "err", err)
			return
		}

    perPageRaw := r.URL.Query().Get("per_page")
    perPage := int64(10)
    if perPageRaw != "" {
      perPage, err = strconv.ParseInt(perPageRaw, 10, 64)
      if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid per_page value."))
        level.Debug(lgr).Log("msg", "failed to get per_page", "err", err)
        return
      }
    }

    tokenQ := []string{}
    token := r.URL.Query().Get("token")
    if token != "" {
      tokenQ = append(tokenQ, token)
    }

		u, err := db.ListUnits(client, int(perPage), tokenQ...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			level.Error(lgr).Log("msg", "failed to get units", "err", err)
			return
		}

		res := response{
      PerPage: int(perPage),
      NextToken: u.NextToken,
    }

		for _, unit := range u.Data {
      if unit.Avatar != nil {
        avatarUrl := client.Storage.GetPublicUrl("units_avatars", *unit.Avatar, storage_go.UrlOptions{
          Transform: &storage_go.TransformOptions{
            Width:  64,
            Height: 64,
          },
        }).SignedURL
        unit.Avatar = &avatarUrl
      }

			mc, err := db.GetMemberCount(client, unit.Id.String())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				level.Error(lgr).Log("msg", "failed to get unit member count", "err", err)
				return
			}

			ri := responseItem{
				Unit:        unit,
				MemberCount: mc,
			}
			res.Data = append(res.Data, ri)
		}

		_ = encode(w, http.StatusOK, res)
	})
}
