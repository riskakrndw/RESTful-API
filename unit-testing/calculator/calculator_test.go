package calculator

import "testing"

func TestAddAllPositiveNumbers(t *testing.T) {
	expected := 19
	a := 4
	b := 15
	actual := Addition(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	expected := -11
	a := 4
	b := -15
	actual := Addition(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestAddAllNegativeNumbers(t *testing.T) {
	expected := -12
	a := -7
	b := -5
	actual := Addition(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestSubAllPositiveNumbers(t *testing.T) {
	expected := 27
	a := 56
	b := 29
	actual := Subtraction(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestSubNegativeNumbers(t *testing.T) {
	expected := -27
	a := 56
	b := 83
	actual := Subtraction(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestSubAllNegativeNumbers(t *testing.T) {
	expected := -1
	a := -5
	b := -4
	actual := Subtraction(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestDivAllPositiveNumbers(t *testing.T) {
	expected := 0.75
	a := 1500
	b := 2000
	actual := Division(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %f, get %f", a, b, expected, actual)
	}
}

func TestDivNegativeNumbers(t *testing.T) {
	expected := -0.75
	a := -1500
	b := 2000
	actual := Division(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %f, get %f", a, b, expected, actual)
	}
}

func TestDivAllNegativeNumbers(t *testing.T) {
	var expected float64
	expected = 2
	a := -2000
	b := -1000
	actual := Division(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %f, get %f", a, b, expected, actual)
	}
}

func TestDivZeroByPositiveNumber(t *testing.T) {
	var expected float64
	expected = 0
	a := 0
	b := 1000
	actual := Division(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %f, get %f", a, b, expected, actual)
	}
}

func TestMulAllPositiveNumbers(t *testing.T) {
	expected := 150000
	a := 1500
	b := 100
	actual := Multiplication(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestMulNegativeNumbers(t *testing.T) {
	expected := -3000000
	a := -1500
	b := 2000
	actual := Multiplication(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}

func TestMulAllNegativeNumbers(t *testing.T) {
	expected := 2000
	a := -20
	b := -100
	actual := Multiplication(a, b)
	if expected != actual {
		t.Errorf("add(%d, %d) should be %d, get %d", a, b, expected, actual)
	}
}
