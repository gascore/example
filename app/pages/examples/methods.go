package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// Example application #4
//
// 'methods' shows how you can call methods from child components.
// Just use interface{Functions} or pass function as argument
func Methods() *gas.E {
	root := &MethodsRoot{
		Show:  false,
		Count: 0,
	}
	c := &gas.C{Root: root}
	root.c = c

	return c.Init()
}

type MethodsRoot struct {
	c *gas.C

	Show  bool
	Count int
}

func (root *MethodsRoot) Toggle() {
	root.Show = !root.Show
	if root.Show {
		root.Count++
	}
	root.c.Update()
}

func (root *MethodsRoot) Render() []interface{} {
	return gas.CL(
		getButton(root.Show, root), // with interface{}
		func() interface{} {
			if root.Show {
				return getHiddenText(
					root.Show,
					func(text string) string { // with function as argument
						fmt.Println("Some text from child component: ", text)
						return fmt.Sprintf("You showed hidden text: %d times", root.Count)
					})
			}
			return nil // nil will be ignored while rendering
		}(),
	)
}

func getButton(show bool, root interface{ Toggle() }) *gas.Element {
	return gas.NE(
		&gas.E{
			Handlers: map[string]gas.Handler{
				"click": func(e gas.Event) {
					root.Toggle()
				},
			},
			Tag: "button",
			Attrs: func() gas.Map {
				return gas.Map{
					"id":    "M&C__button",
					"class": "button outline",
				}
			},
		},
		func() interface{} {
			if show {
				return "Show text"
			} else {
				return "Hide text"
			}
		}(),
	)
}

func getHiddenText(show bool, getNumber func(string) string) *gas.Element {
	return gas.NE(
		&gas.E{
			Tag: "i",
		},
		"Hidden text",
		"  (", getNumber("something for method"), ")",
	)
}
