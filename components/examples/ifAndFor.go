package examples

import "github.com/gascore/gas"

// Example application #3
//
// 'if-and-for-directive' shows how you can use component.Directive.If
func IfAndFor() *gas.E {
	root := &IfAndForRoot{
		Show: false,
		Arr:  []string{"click", "here", "if you want to see some magic"},
	}
	c := &gas.C{NotPointer: true, Root: root}
	root.c = c

	return c.Init()
}

type IfAndForRoot struct {
	c *gas.C

	Show bool
	Arr  []string
}

func (root *IfAndForRoot) Add(el string) {
	root.Arr = append(root.Arr, el)
}

func (root *IfAndForRoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id": "block-if",
				},
			},
			gas.NE(
				&gas.E{
					Handlers: map[string]gas.Handler{
						"click": func(e gas.Object) {
							root.Show = !root.Show
							root.c.Update()
						},
					},
					Tag: "button",
					Attrs: map[string]string{
						"id": "if__button",
					},
				},
				// here is how you can use IF
				func() interface{} {
					if root.Show {
						return "Show text"
					} else {
						return "Hide text"
					}
				}(),
			),
			func() interface{} {
				if root.Show {
					return gas.NE(
						&gas.E{
							Tag: "i",
						},
						"Public text",
					)
				} else {
					return gas.NE(
						&gas.E{
							Tag: "b",
						},
						"Hidden text",
					)
				}
			}(),
		),
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id": "block-for",
				},
			},
			gas.NE(
				&gas.E{
					Tag: "ul",
				},
				func() []*gas.E {
					var arr []*gas.E
					for i, el := range root.Arr {
						arr = append(arr, gas.NE(
							&gas.E{
								Handlers: map[string]gas.Handler{
									"click": func(e gas.Object) {
										root.Add("Hello!") // hello, Annoy-o-Tron
										root.c.Update()
									},
								},
								Tag: "li",
							},
							i+1, ": ", el,
						))
					}
					return arr
				}(),
			),
		),
	)
}
