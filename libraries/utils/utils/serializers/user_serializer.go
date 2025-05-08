package serializers

import (
	"github.com/leothemighty/oliveplay/utils/db"
)

type UserSerializer struct {
	ID int32
}

func UserSerializerFromCreateUserRow(row db.GetGroupsWithinRadiusRow) (UserSerializer, error) {
	return UserSerializer{
		ID: row.ID,
	}, nil
}

func UserSerializerFromFetchUserRow() {
}
