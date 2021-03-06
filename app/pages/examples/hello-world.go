package examples

import (
	"github.com/gascore/gas"
)

type HelloRoot struct {
	Hello string
}

func (root *HelloRoot) Render() *gas.E {
	return gas.NE(
		&gas.E{},
		gas.NE(
			&gas.E{
				Tag: "h1",
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "hello-world",
						"class": "greeting h1",
					}
				},
			},
			root.Hello),
		gas.NE(
			&gas.E{
				Tag: "i",
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "italiano",
						"class": "greeting",
						"style": "margin-right: 12px;",
					}
				},
			},
			"Ciao mondo!"),
	)
}

// Example application #1
//
// 'hello-world' shows how you can create components, component.Data and component.Attributes
func Hello() *gas.C {
	c := &gas.C{
		Root: &HelloRoot{
			Hello: "Hello world!",
		},
		NotPointer: true,
	}

	return c
}
