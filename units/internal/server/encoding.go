package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Validator interface {
	Valid(ctx context.Context) (problems map[string]string)
}

func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return errors.Join(errors.New("encode json"), err)
	}
	return nil
}

func decodeValid[T Validator](r *http.Request) (T, map[string]string, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, nil, errors.Join(errors.New("decode json"), err)
	}
	if problems := v.Valid(r.Context()); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}
	return v, nil, nil
}
