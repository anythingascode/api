package endpoints

import (
	"fmt"
	"testing"
)

func TestPostType(t *testing.T) {
	for k, v := range PostEndPoints {
		if fmt.Sprintf("%T", k) != "string" {
			t.Fatalf("Endpoints key type should be string, found %v", fmt.Sprintf("%T", v))
		}
		if fmt.Sprintf("%T", v) != "func(*gin.Context)" {
			t.Fatalf("Endpoints key type should be func, found %v", fmt.Sprintf("%T", v))
		}
	}
}
