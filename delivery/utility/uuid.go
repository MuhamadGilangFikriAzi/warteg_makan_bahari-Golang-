package utility

import "github.com/google/uuid"

func GetUuid() string {
	return uuid.New().String()
}
