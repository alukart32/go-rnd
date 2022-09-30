package decorator

import (
	"errors"
	"fmt"
)

// component interface
type Ingredient interface {
	AddIngredient() (string, error)
}

// decorator that combine pizza description
type PizzaDecorator struct {
	Ingredient Ingredient
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with:", nil
}

// concrete component for pizza
type PizzaMeat struct {
	Ingredient Ingredient
}

func (m *PizzaMeat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("an AddIngredient is needed in the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}

type PizzaOnion struct {
	Ingredient Ingredient
}

func (o *PizzaOnion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("an AddIngredient is needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

type PizzaCheezeDecorator struct {
	Ingredient Ingredient
}

func (p *PizzaCheezeDecorator) AddIngredient() (string, error) {
	if p.Ingredient == nil {
		return "", errors.New("an AddIngredient is needed in the Ingredient field of the Cheeze Decorator")
	}
	s, err := p.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "and toppings"), nil
}

type PizzaWithCheeze struct {
	Ingredient Ingredient
}

func (p *PizzaWithCheeze) AddIngredient() (string, error) {
	if p.Ingredient == nil {
		return "", errors.New("an AddIngredient is needed in the Ingredient field of the Cheeze toppings")
	}
	s, err := p.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", s, "cheeze"), nil
}
