package natural

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	value := []string{
		"３",
		"1",
		"10",
		"2",
		"20",
	}
	expected := []string{
		"1",
		"2",
		"３",
		"10",
		"20",
	}
	Sort(value)

	if !reflect.DeepEqual(value, expected) {
		t.Fatalf("expected %v, got %v:", expected, value)
	}
}
