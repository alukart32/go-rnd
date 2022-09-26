// src: https://blog.ralch.com/articles/design-patterns/golang-composite/
//
// The Composite Pattern provides the following units:
// - Component is an interface for all components, including composite ones. It declares the interface for objects in the composition
// - Leaf represents leaf objects in the composition implements all Component methods
// - Composite represents a composite Component that has children. Usually implements all Componenet methods and methods to manipulate children.

// image editors, which compose different shapes and layers into hierarchy.
package composite

// VisualElement that is drawn on the screen
// "Component interface"
// type VisualElement interface {
// 	Draw(drawer *Drawer) error
// }

// The editor supports two kind of shapes (circle and square). Each of the structs that represents
// the coresponding shape obeys the VisualElement interface by implementing a Draw function
// that has exactly the same signiture exposed in the interface.
// "Leaf"
// type Square struct {
// 	Location Point
// 	Side     float64
// }

// Draw draws a square
// func (square *Square) Draw(drawer *Drawer) error {
// 	return drawer.DrawRect(Rect{
// 		Location: square.Location,
// 		Size: Size{
// 			Height: square.Side,
// 			Width:  square.Side,
// 		},
// 	})
// }

// Circle represents a circle shape
// "Leaf"
// type Circle struct {
// 	// Center of the circle
// 	Center Point
// 	// Radius of the circle
// 	Radius float64
// }

// Draw draws a circle
// func (circle *Circle) Draw(drawer *Drawer) error {
// 	rect := Rect{
// 		Location: Point{
// 			X: circle.Center.X - circle.Radius,
// 			Y: circle.Center.Y - circle.Radius,
// 		},
// 		Size: Size{
// 			Width:  2 * circle.Radius,
// 			Height: 2 * circle.Radius,
// 		},
// 	}

// 	return drawer.DrawEllipseInRect(rect)
// }

// In order to allow composition and drawing of multiple shapes on the screen,
// a Layer compose thoes object. It contains an array of VisualElement. It is responsible
// to interate over the elements and draw each of them. As you can see the actual struct uses
// the VisualElement interface as a contract to support different shapes no matter what is their type.
// "Composition"
// type Layer struct {
// 	Elements []VisualElement
// }

// Draw draws a layer
// func (layer *Layer) Draw(drawer *Drawer) error {
// 	for _, element := range layer.Elements {
// 		if err := element.Draw(drawer); err != nil {
// 			return err
// 		}
// 		fmt.Println()
// 	}

// 	return nil
// }

// func ImageEditorComposition() {
// 	circle := &Circle{
// 		Center: Point{X: 100, Y: 100},
// 		Radius: 50,
// 	}

// 	square := &Square{
// 		Location: Point{X: 50, Y: 50},
// 		Side:     20,
// 	}

// 	layer := &Layer{
// 		Elements: []VisualElement{
// 			circle,
// 			square,
// 		},
// 	}

// 	layer.Draw(&Drawer{})
// }
