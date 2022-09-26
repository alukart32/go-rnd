package prototype

import (
	"testing"
)

func TestShirtsClone(t *testing.T) {
	shirtCache := NewShirtsCache()
	if shirtCache == nil {
		t.Fatal("received cache was nil")
	}

	whiteShirtPrototype, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if whiteShirtPrototype.(*Shirt) == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := whiteShirtPrototype.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt1 couldn't be done successfully")
	}
	shirt1.SKU = "abbcc"

	whiteShirtPrototype2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatal(err)
	}

	shirt2, ok := whiteShirtPrototype2.(*Shirt)
	if !ok {
		t.Fatal("type assertion for shirt2 couldn't be done successfully")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU's of shirt1 and shirt2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("shirt 1 cannot be equal to Shirt 2")
	}

	t.Logf("log: %s", shirt1.GetInfo())
	t.Logf("log: %s", shirt2.GetInfo())

	t.Logf("log: the memory positions of the shirts are different %p != %p \n\n", &shirt1, &shirt2)
}
