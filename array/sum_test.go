package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("got %d expected %d given, %v", result, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	expected := []int{6, 15}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %d expected %d", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %d expected %d", result, expected)
		}
	}

	t.Run("make the sum of som slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		expected := []int{5, 11}

		checkSums(t, result, expected)
	})
	t.Run("make the sum of a empty slice", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{4, 5, 6})
		expected := []int{0, 11}

		checkSums(t, result, expected)
	})
}

