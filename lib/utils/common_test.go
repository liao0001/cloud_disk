package utils

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		uid := uuid.New()
		fmt.Println(uid.ID(), uid.String(), uid.String(), len(uid.String()))
	}
}
