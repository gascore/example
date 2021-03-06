package examples

import "github.com/gascore/gas"

// Example application #3
//
// 'if-and-for-directive' shows how you can use component.Directive.If
func IfAndFor() *gas.C {
	root := &IfAndForRoot{
		Show: false,
		Arr:  []string{"click", "here", "if you want to see some magic"},
	}
	c := &gas.C{
		Root: root,
		NotPointer: true,
	}
	root.c = c

	return c
}

type IfAndForRoot struct {
	c *gas.C

	Show bool
	Arr  []string
}

func (root *IfAndForRoot) Add(el string) {
	root.Arr = append(root.Arr, el)
}

func (root *IfAndForRoot) Render() *gas.E {
	return gas.NE(
		&gas.E{},
		gas.NE(
			&gas.E{
				Attrs: func() gas.Map {
					return gas.Map{
						"id": "block-if",
					}
				},
			},
			gas.NE(
				&gas.E{
					Handlers: map[string]gas.Handler{
						"click": func(e gas.Event) {
							root.Show = !root.Show
							root.c.Update()
						},
					},
					Tag: "button",
					Attrs: func() gas.Map {
						return gas.Map{
							"id":    "if__button",
							"class": "button outline",
							"style": "margin-right: .4em;",
						}
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
				Attrs: func() gas.Map {
					return gas.Map{
						"id": "block-for",
					}
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
									"click": func(e gas.Event) {
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
