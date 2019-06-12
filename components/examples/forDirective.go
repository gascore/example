package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// ForDirective Example application #4
//
// 'for-directive' shows how you can use component.Directive.For
func ForDirective() *gas.C {
	return gas.NC(
		&gas.C{
			Data: map[string]interface{}{
				"arr": []interface{}{"click", "here", "if you want to see some magic"},
			},
			Tag: "ul",
			Attrs: map[string]string{
				"id": "list",
			},
		},
		func(this *gas.C) []interface{} {
			return gas.CL(
				gas.NE(
					&gas.C{
						Tag: "ul",
					},
					gas.NewFor("arr", this, func(i int, el interface{}) interface{} {
						return gas.NE(
							&gas.C{
								Handlers: map[string]gas.Handler{
									"click": func(c *gas.C, e gas.Object) {
										arr := this.Get("arr").([]interface{})
										arr = append(arr, "Hello!") // hello, Annoy-o-Tron
										this.SetValue("arr", arr)
									},
								},
								Tag: "li",
							},
							fmt.Sprintf("%d: %s", i+1, el))
					}),
					"end of list"))
		})
}
