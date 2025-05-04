package group

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/leothemighty/oliveplay/utils/utils/db"
	"github.com/leothemighty/oliveplay/utils/utils/serializers"
	"github.com/leothemighty/oliveplay/utils/utils/types"
)

func GetCloseGroups(context context.Context, req *http.Request) ([]types.GroupResponse, error) {
	longitude := req.URL.Query().Get("longitude")
	latitude := req.URL.Query().Get("latitude")
	radius := req.URL.Query().Get("radius")

	if longitude == "" || latitude == "" || radius == "" {
		return nil, errors.New("longitude, latitude, and radius are required")
	}

	database := db.OpenDatabase()
	defer db.CloseDatabase(database)

	get_params, err := serializers.SerializeGetCloseGroups(longitude, latitude, radius)
	if err != nil {
		log.Printf("Failed to serialize get close groups: %v", err)
		return nil, err
	}

	groups, err := database.Queries.GetGroupsWithinRadius(context, get_params)
	if err != nil {
		log.Printf("Failed to get close groups: %v", err)
		return nil, err
	}

	group_responses, err := serializers.DeserializeGetCloseGroups(groups)
	if err != nil {
		log.Printf("Failed to serialize get close groups: %v", err)
		return nil, err
	}

	return groups, nil
}
