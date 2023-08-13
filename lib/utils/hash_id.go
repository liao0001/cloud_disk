package utils

import "github.com/google/uuid"

func NewHashID() string {
	return uuid.NewString()
}
