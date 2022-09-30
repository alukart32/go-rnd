package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator(t *testing.T) {
	pizza := &PizzaDecorator{}
	completePizza, _ := pizza.AddIngredient()
	expectedPizzaDescr := "Pizza with:"
	if !strings.Contains(completePizza, expectedPizzaDescr) {
		t.Errorf("When calling the add ingredient of the pizza decorator it must return the text %sthe expected text, not '%s'", completePizza, expectedPizzaDescr)
	}
}

func TestPizzaDecoratorAddOnion(t *testing.T) {
	onion := &PizzaOnion{}
	onionPizzaDescr, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the onion decorator without"+
			"an IngredientAdd on its Ingredient field must return an error, not a string with '%s'", onionPizzaDescr)
	}

	onion = &PizzaOnion{&PizzaDecorator{}}
	onionPizzaDescr, err = onion.AddIngredient()

	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionPizzaDescr, "onion") {
		t.Errorf("When calling the add ingredient of the onion decorator it"+
			"must return a text with the word 'onion', not '%s'", onionPizzaDescr)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &PizzaMeat{}
	meatPizzaDescr, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the meat decorator without"+
			"an IngredientAdd in its Ingredient field must return an error,"+
			"not a string with '%s'", meatPizzaDescr)
	}

	meat = &PizzaMeat{&PizzaDecorator{}}
	meatPizzaDescr, err = meat.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(meatPizzaDescr, "meat") {
		t.Errorf("When calling the add ingredient of the meat decorator it"+
			"must return a text with the word 'meat', not '%s'", meatPizzaDescr)
	}
}

func TestPizzaWithOnionAndCheezeToppings(t *testing.T) {
	pizzaOnionWithCheeze := &PizzaWithCheeze{&PizzaCheezeDecorator{&PizzaOnion{&PizzaDecorator{}}}}
	expectedDescr := "Pizza with: onion, and toppings cheeze"
	actual, err := pizzaOnionWithCheeze.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if expectedDescr != actual {
		t.Errorf("expected: %s; actual: %s", expectedDescr, actual)
	}
}
