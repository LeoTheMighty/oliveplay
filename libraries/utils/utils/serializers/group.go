package serializers

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/leothemighty/oliveplay/utils/utils/db"
	"github.com/leothemighty/oliveplay/utils/utils/types"
)

func SerializeGroup(group db.CreateGroupRow) (types.GroupResponse, error) {
	var description string
	if group.Description.Valid {
		description = group.Description.String
	}

	var image *string
	if group.Image.Valid {
		image = &group.Image.String
	}

	return types.GroupResponse{
		ID:          strconv.Itoa(int(group.ID)),
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		Name:        group.Name,
		Description: description,
		Image:       image,
		Longitude:   group.Longitude.Float64,
		Latitude:    group.Latitude.Float64,
	}, nil
}

func SerializeGetCloseGroups(longitude string, latitude string, radius string) (db.GetGroupsWithinRadiusParams, error) {
	longitude_float, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return db.GetGroupsWithinRadiusParams{}, err
	}

	latitude_float, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return db.GetGroupsWithinRadiusParams{}, err
	}

	radius_float, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		return db.GetGroupsWithinRadiusParams{}, err
	}

	return db.GetGroupsWithinRadiusParams{
		StMakepoint:   longitude_float,
		StMakepoint_2: latitude_float,
		StDwithin:     radius_float,
	}, nil
}

func DeserializeCreateGroup(group_request types.GroupRequest) (db.CreateGroupParams, error) {
	description := pgtype.Text{}
	if group_request.Description != "" {
		description.String = group_request.Description
		description.Valid = true
	}

	image := pgtype.Text{}
	if group_request.Image != nil {
		image.String = *group_request.Image
		image.Valid = true
	}

	longitude, err := strconv.ParseFloat(group_request.Longitude, 64)
	if err != nil {
		return db.CreateGroupParams{}, err
	}

	latitude, err := strconv.ParseFloat(group_request.Latitude, 64)
	if err != nil {
		return db.CreateGroupParams{}, err
	}

	return db.CreateGroupParams{
		Name:        group_request.Name,
		Description: description,
		Image:       image,
		Longitude:   longitude,
		Latitude:    latitude,
	}, nil
}

func DeserializeGetCloseGroups(groups []db.GetGroupsWithinRadiusRow) ([]types.GroupResponse, error) {
	group_responses := []types.GroupResponse{}
	for _, group := range groups {
		group_response := types.GroupResponse{
			ID:          strconv.Itoa(int(group.ID)),
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
			Name:        group.Name,
			Description: group.Description.String,
			Image:       group.Image.String,
			Longitude:   group.Longitude.Float64,
			Latitude:    group.Latitude.Float64,
		}

		group_responses = append(group_responses, group_response)
	}
	return group_responses, nil
}
