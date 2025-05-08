package group

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/leothemighty/oliveplay/utils/db"
	"github.com/leothemighty/oliveplay/utils/serializers"
	"github.com/leothemighty/oliveplay/utils/types"
)

func CreateGroup(ctx context.Context, req *http.Request) (types.GroupResponse, error) {
	var gr types.GroupRequest
	if err := json.NewDecoder(req.Body).Decode(&gr); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		return types.GroupResponse{}, err
	}

	create_params, err := serializers.DeserializeGroup(gr)
	if err != nil {
		log.Printf("Failed to deserialize group: %v", err)
		return types.GroupResponse{}, err
	}

	database := db.OpenDatabase()
	group, err := database.Queries.CreateGroup(context.Background(), create_params)
	if err != nil {
		log.Printf("Failed to create group: %v", err)
		return types.GroupResponse{}, err
	}

	group_response, err := serializers.SerializeGroup(group)
	if err != nil {
		log.Printf("Failed to serialize group: %v", err)
		return types.GroupResponse{}, err
	}

	db.CloseDatabase(database)

	return group_response, nil
}
