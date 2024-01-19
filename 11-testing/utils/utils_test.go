package utils

import (
	"fmt"
	"testing"
)

func Test_IsPrime_97(t *testing.T) {
	// Arrange
	no := 97
	expectedResult := false

	t.Log("Testing IsPrime(97)")
	// Act
	actualResult := IsPrime(no)

	// Assert
	if actualResult != expectedResult {
		/*
			t.Logf("IsPrime(97), expected = %v, but actual = %v\n", expectedResult, actualResult)
			t.Fail()
		*/
		t.Errorf("IsPrime(97), expected = %v, but actual = %v\n", expectedResult, actualResult)
	}
}

// data drive tests
func Test_IsPrime(t *testing.T) {
	test_data := []struct {
		no             int
		expectedResult bool
	}{
		{no: 13, expectedResult: true},
		{no: 14, expectedResult: false},
		{no: 15, expectedResult: false},
		{no: 17, expectedResult: true},
		{no: 19, expectedResult: true},
	}
	for _, td := range test_data {
		t.Run(fmt.Sprintf("Is_Prime:%d", td.no), func(t *testing.T) {
			t.Parallel()
			actualResult := IsPrime(td.no)

			// Assert
			if actualResult != td.expectedResult {
				t.Errorf("IsPrime(%d), expected = %v, but actual = %v\n", td.no, td.expectedResult, actualResult)
			}
		})
	}

	t.Cleanup(func() {
		if results["prime"] != 4 {
			t.Error("results were not updated")
		}
	})

}

// Benchmarking
func Benchmark_Is_Prime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(97)
	}
}
