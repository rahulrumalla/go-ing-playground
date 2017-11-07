package fibonacci

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	res := Generate(10)
	fmt.Printf("%+v\n", res)
}
