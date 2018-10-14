package decimal_test

import (
	"testing"

	"github.com/nu11ptr/decimal"
)

func TestNewAndString(t *testing.T) {
	tests := []struct {
		name, input, output string
		prec                int
		ok                  bool
	}{
		{"Zero", "0", "0.00", 2, true},
		{"Positive", "12345.67890", "12345.67890", 5, true},
		{"Negative", "-12345.67890", "-12345.67890", 5, true},
		{"Bad", "bad", "", -1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d, ok := decimal.New(test.input)
			if ok != test.ok {
				t.Error("Got", ok, "Expected", test.ok)
			}
			if ok {
				if d.String(test.prec) != test.output {
					t.Error("Got", d.String(test.prec), "Expected", test.output)
				}
			}
		})
	}
}

func TestInt(t *testing.T) {
	d := decimal.NewInt(123)
	if d.String(2) != "123.00" {
		t.Error("Got", d.String(0), "Expect", "123")
	}
}

func TestLT(t *testing.T) {
	if !decimal.NewInt(1).LT(decimal.NewInt(2)) {
		t.Error("Expected 1 to be less than 2")
	}
}

func TestLTE(t *testing.T) {
	if !decimal.NewInt(1).LTE(decimal.NewInt(2)) {
		t.Error("Expected 1 to be less than 2")
	}
	if !decimal.NewInt(2).LTE(decimal.NewInt(2)) {
		t.Error("Expected 2 to be equal to 2")
	}
}

func TestEq(t *testing.T) {
	if !decimal.NewInt(2).Eq(decimal.NewInt(2)) {
		t.Error("Expected 2 to be equal to 2")
	}
}

func TestGT(t *testing.T) {
	if !decimal.NewInt(2).GT(decimal.NewInt(1)) {
		t.Error("Expected 2 to be greater than 1")
	}
}

func TestGTE(t *testing.T) {
	if !decimal.NewInt(2).GTE(decimal.NewInt(1)) {
		t.Error("Expected 2 to be greater than 1")
	}
	if !decimal.NewInt(2).GTE(decimal.NewInt(2)) {
		t.Error("Expected 2 to be equal to 2")
	}
}

func TestAdd(t *testing.T) {
	actual := decimal.NewMust("100.100").Add(decimal.NewMust("200.200")).String(3)
	expected := decimal.NewMust("300.300").String(3)
	if actual != expected {
		t.Error("Got", actual, "Expected", expected)
	}
}

func TestSub(t *testing.T) {
	actual := decimal.NewMust("300.300").Sub(decimal.NewMust("200.200")).String(3)
	expected := decimal.NewMust("100.100").String(3)
	if actual != expected {
		t.Error("Got", actual, "Expected", expected)
	}
}

func TestMult(t *testing.T) {
	actual := decimal.NewMust("300.300").Mult(decimal.NewMust("200.200")).String(2)
	expected := decimal.NewMust("60120.06").String(2)
	if actual != expected {
		t.Error("Got", actual, "Expected", expected)
	}
}

func TestDiv(t *testing.T) {
	actual := decimal.NewMust("300.300").Div(decimal.NewMust("200.200")).String(1)
	expected := decimal.NewMust("1.5").String(1)
	if actual != expected {
		t.Error("Got", actual, "Expected", expected)
	}
}
