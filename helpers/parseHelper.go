package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ParseRequestToModel(r *http.Request, dto interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding data: %v", err))
	}

	return nil
}
