package examples

import (
	"github.com/gascore/gas"
)

type HelloRoot struct {
	Hello string
}

func (root *HelloRoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Tag: "h1",
				Attrs: func() map[string]string {
					return map[string]string{
						"id":    "hello-world",
						"class": "greeting h1",
					}
				},
			},
			root.Hello),
		gas.NE(
			&gas.E{
				Tag: "i",
				Attrs: func() map[string]string {
					return map[string]string{
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
func Hello() *gas.E {
	c := &gas.C{
		Root: &HelloRoot{
			Hello: "Hello world!",
		},
	}

	return c.Init()
}
