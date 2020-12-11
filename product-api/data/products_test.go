package data

import "testing"

func TestChecksValidation(t *testing.T) {

	p := &Product{
		// Name:  "Tester",
		// Price: 69.69,
		// SKU:   "a-b-c",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
