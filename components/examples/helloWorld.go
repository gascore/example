package examples

import (
	"github.com/gascore/gas"
)

// HelloWorld application #1
//
// 'hello-world' shows how you can create components, component.Data and component.Attributes
func HelloWorld() *gas.C {
	return gas.NC(
		&gas.C{
			Data: map[string]interface{}{
				"hello": "Hello world!",
			},
		},
		func(this *gas.C) []interface{} {
			return gas.CL(
				gas.NE(
					&gas.C{
						Tag: "h1",
						Attrs: map[string]string{
							"id":    "hello-world",
							"class": "greeting h1",
						},
					},
					this.Get("hello")),
				gas.NE(
					&gas.C{
						Tag: "i",
						Attrs: map[string]string{
							"id":    "italiano",
							"class": "greeting",
							"style": "margin-right: 12px;",
						},
					},
					"Ciao mondo!"))
		})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
