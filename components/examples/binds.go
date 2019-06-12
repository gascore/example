package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// Binds Example application #10
//
// 'binds' shows how you can use component.Binds
func Binds() *gas.Component {
	return gas.NC(
		&gas.C{
			Data: map[string]interface{}{
				"foo": int(0),
			},
			Attrs: map[string]string{
				"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
			},
		},
		func(this *gas.C) []interface{} {
			return gas.CL(
				gas.NE(
					&gas.C{
						Tag: "div",
						Attrs: map[string]string{
							"style": "display: flex;",
						},
					},
					fmt.Sprintf("Your range: %d", this.Get("foo").(int)),
					gas.NE(
						&gas.C{
							Binds: map[string]gas.Bind{
								"style": func() string {
									foo, ok := this.Get("foo").(int)
									this.WarnIfNot(ok)

									// REMEMBER!
									// Bind attributes appends over normal attributes, you will lose your normal attribute value
									return fmt.Sprintf("width: 48px; height: 36px; margin: 0 18px; border-radius: 4px; background-color: rgb(%d, %d, %d)", foo, 255-foo, foo)
								},
							},
							Tag: "div",
						}),
					gas.NE(
						&gas.C{
							Attrs: map[string]string{
								"style": "color: darkgray;",
							},
							Tag: "i",
						},
						"// color: rgb(x, 255-x, x)",
					),
				),
				gas.NE(&gas.C{Tag: "br"}),
				gas.NE(
					&gas.C{
						Model: gas.ModelDirective{
							Data:      "foo",
							Component: this,
						},
						Tag: "input",
						Attrs: map[string]string{
							"type": "range",
							"min":  "0",
							"max":  "255",
						},
					}),
			)
		})
}
