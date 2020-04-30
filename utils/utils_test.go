package utils

import (
	"fmt"
	"testing"
)

func TestGetUUid(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("===", GetUuidInt64())
		fmt.Println("===", GetUuidString())
	}
}
