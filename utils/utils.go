package utils

import "github.com/google/uuid"

func Guid() string {
	return uuid.New().String()
}
