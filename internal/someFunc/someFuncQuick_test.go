package some_func

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

/*
	type Generator interface {
		// Generate returns a random instance of the type on which it is a
		// method using the size as a size hint.
		Generate(rand *rand.Rand, size int) reflect.Value
	}
*/

func (Vector) Generate(rand *rand.Rand, size int) reflect.Value {
	v := Vector{X: int(rand.Int31()), Y: int(rand.Int31())}
	fmt.Println("generate V", v.X, v.Y)
	return reflect.ValueOf(v)
}

func TestSumQ(t *testing.T) {
	config := quick.Config{MaxCount: 4}

	f := func(v Vector) bool {
		x := v.Len()
		return x == v.X*v.Y && x > v.X && x > v.Y // тест падает при 0 или int overflow
		// например если в генераторе использовать rand.Int вместо rand.Int31
	}

	err := quick.Check(f, &config)
	if err != nil {
		t.Error(err)
	}
}
