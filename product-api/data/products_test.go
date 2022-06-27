package data

import "testing"

func TestChecksValidation(t *testing.T) {
	P := Product{
		Name:  "Book2",
		Price: 10,
		SKU:   "abc-def-ijk",
	}
	err := P.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
