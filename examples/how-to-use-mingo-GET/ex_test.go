package examples

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {}

func TestGetEndPoints(t *testing.T) {
	endp, err := GetEndp()

	fmt.Println(err)
	fmt.Println(endp)
}
