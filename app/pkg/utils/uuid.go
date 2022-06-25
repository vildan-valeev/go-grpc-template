package utils

import "github.com/google/uuid"

func ReplaceNilUUID(id uuid.UUID) string {
	if id != uuid.Nil {
		return id.String()
	}

	return ""
}

func ReplaceStringToUUID(id string) uuid.UUID {
	res, _ := uuid.Parse(id)
	return res
}
