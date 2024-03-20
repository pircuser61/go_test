package some_func

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	x := Sum(1, 2)
	want := 3
	if x != want {
		t.Errorf("Sum(1,2) want %d got %d", want, x)
	}
}

func TestSumTable(t *testing.T) {
	type testItem struct {
		name string
		x    int
		y    int
		want int
	}
	testSet := []testItem{
		{"positive", 1, 2, 3},
		{"zero", 0, 0, 0},
		{"negative", -1, 1, 0},
	}

	for _, ti := range testSet {
		t.Run(ti.name, func(t *testing.T) {
			got := Sum(ti.x, ti.y)
			if ti.want != got {
				t.Errorf("Sum want %d got %d", ti.want, got)
			}
		})
	}
}

func TestSumParallelTable(t *testing.T) {
	type testItem struct {
		name string
		x    int
		y    int
		want int
	}
	testSet := []testItem{
		{"positive", 1, 2, 3},
		{"zero", 0, 0, 0},
		{"1 negative", -1, 1, 0},
		{"2 negative", -1, -1, -2},
	}

	for _, ti := range testSet {
		ti := ti
		t.Run(ti.name, func(t *testing.T) {
			t.Parallel()
			got := LongCalc(ti.x, ti.y)
			if ti.want != got {
				t.Errorf("Sum want %d got %d", ti.want, got)
			}
		})
	}
}

func TestError(t *testing.T) {
	x, err := Div(1, 0)
	if err == nil {
		t.Errorf("want Error got %d", x)
	}
}

func TestSlice(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{-2, -2, -2}
	want := []int{-1, 0, 1}

	got, err := SSum(x, y)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("SSum wrong result")
	}
}
