// src: https://blog.devgenius.io/go-composite-pattern-393caaa0b105

package composite

import "fmt"

type Component interface {
	Display()
}

// Leaf of the hierarchy
type Employee struct {
	id   string
	name string
}

func (e Employee) Display() {
	fmt.Printf("ID: %s\nName: %s\n", e.id, e.name)
}

func NewEmployee(id, name string) Employee {
	return Employee{
		id:   id,
		name: name,
	}
}

// the organizational hierarchy can be treated similar to a Tree data structure
type Composite struct {
	component Component   // The root of the Tree/Head of the organization
	leaves    []Component // The leave of the Tree/Subordinate of the organization
}

func (c *Composite) Add(leaf Component) {
	c.leaves = append(c.leaves, leaf)
}

func (c Composite) Display() {
	c.component.Display()
	if len(c.leaves) == 0 {
		return
	}
	fmt.Println("===List of Subordinates===")
	for _, leaf := range c.leaves {
		leaf.Display()
	}
	fmt.Println("===End===")
}

func NewComposite(component Component) Composite {
	return Composite{
		component: component,
	}
}

// Build the hierarchy of office of VP of Idealist
// Build the hierarchy of office of VP of Realists
// Build the hierarchy of the company, by adding the Ceo and add two offices of Vice Presidents
func ShowEmployeeHierarchy() {
	ceo := NewEmployee("ID-1", "Socrates")
	vpIdealist := NewEmployee("ID-2", "Plato")
	vpRealist := NewEmployee("ID-3", "Aristotle")
	directorIdealist := NewEmployee("ID-4", "Hegel")
	directorRealist := NewEmployee("ID-5", "Hume")

	directorIdealistOffice := NewComposite(directorIdealist)
	directorRealistOffice := NewComposite(directorRealist)

	vpIdealistOffice := NewComposite(vpIdealist)
	vpIdealistOffice.Add(directorIdealistOffice)

	vpRealistOffice := NewComposite(vpRealist)
	vpRealistOffice.Add(directorRealistOffice)

	company := NewComposite(ceo)
	company.Add(vpRealistOffice)
	company.Add(vpIdealistOffice)

	fmt.Println("=====Display Office of VP of Idealist=====")
	vpIdealistOffice.Display()
	fmt.Println("=====Display Office of VP of Realist=====")
	vpRealistOffice.Display()
	fmt.Println("=====Display Company=====")
	company.Display()
}
