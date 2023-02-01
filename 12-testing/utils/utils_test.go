package utils

import "testing"

/*
func TestIsPrime(t *testing.T) {
	// arrange
	no := 17
	expectedResult := true

	//act
	actualResult := IsPrime(no)

	//assert
	// t.Log(expectedResult, actualResult)
	if actualResult != expectedResult {
		// t.Fail()
		t.Errorf("IsPrime(%d), expected = %t, got %t", 17, expectedResult, actualResult)
	}
}

func TestIsPrime2(t *testing.T) {
	// arrange
	no := 17
	expectedResult := true

	//act
	actualResult := IsPrime2(no)

	//assert
	// t.Log(expectedResult, actualResult)
	if actualResult != expectedResult {
		// t.Fail()
		t.Errorf("IsPrime(%d), expected = %t, got %t", 17, expectedResult, actualResult)
	}
}

func TestIsPrime3(t *testing.T) {
	// arrange
	no := 17
	expectedResult := true

	//act
	actualResult := IsPrime3(no)

	//assert
	// t.Log(expectedResult, actualResult)
	if actualResult != expectedResult {
		// t.Fail()
		t.Errorf("IsPrime(%d), expected = %t, got %t", 17, expectedResult, actualResult)
	}
}
*/

func TestIsPrime(t *testing.T) {
	testData := []struct {
		no       int
		expected bool
		actual   bool
		name     string
	}{
		{no: 17, name: "IsPrime-17", expected: true},
		{no: 19, name: "IsPrime-19", expected: true},
		{no: 20, name: "IsPrime-20", expected: false},
		{no: 23, name: "IsPrime-23", expected: true},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			td.actual = IsPrime(td.no)
			if td.expected != td.actual {
				t.Errorf("IsPrime(%d), expected = %t, got %t", td.no, td.expected, td.actual)
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(79)
	}
}

func BenchmarkIsPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime2(79)
	}
}

func BenchmarkIsPrime3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime3(79)
	}
}
