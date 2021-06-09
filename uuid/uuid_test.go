package uuid

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	id := uuid.New().String()
	fmt.Println("UUID=" + id)
}
