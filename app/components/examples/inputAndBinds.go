package examples

import (
	"fmt"
	"strconv"

	"github.com/gascore/gas"
)

// Example application #5
//
// 'inputAndBinds' shows how you can use element Handlers and element.Binds for create forms
func InputAndBinds() *gas.E {
	root := &InputAndBindsRoot{
		Text:     "",
		Color:    "#000",
		Range:    0,
		CheckBox: false,
	}

	c := &gas.C{
		Root: root,
	}
	root.c = c

	return c.Init()
}

type InputAndBindsRoot struct {
	c *gas.C

	Text     string
	Color    string
	Range    int
	CheckBox bool
}

func (root *InputAndBindsRoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "model__text",
						"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
					}
				},
			},
			"Your text: ", root.Text,
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Handlers: map[string]gas.Handler {
						"input": func(event gas.Event) {
							root.Text = event.Value()
							go root.c.Update()
						},
					},
					Attrs: func() gas.Map {
						return gas.Map{"value": root.Text}
					},
					Tag:     "input",
				},
			),
		),
		gas.NE(
			&gas.E{
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "model__color",
						"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
					}
				},
			},
			"Your color: ",
			gas.NE(
				&gas.E{
					Tag: "span",
					Attrs: func() gas.Map {
						return gas.Map{
							"style": "color:" + root.Color,
						}
					},
				},
				root.Color,
			),
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Tag: "input",
					Handlers: map[string]gas.Handler {
						"input": func(event gas.Event) {
							root.Color = event.Value()
							go root.c.Update()
						},
					},
					Attrs: func() gas.Map {
						return gas.Map{
							"type": "color",
							"value": root.Color,
						}
					},
				},
			),
		),
		gas.NE(
			&gas.E{
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "model__range",
						"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
					}
				},
			},
			gas.NE(
				&gas.E{
					Tag: "span",
					Attrs: func() gas.Map {
						return gas.Map{
							"style": "display: flex;",
						}
					},
				},
				"Generated color: ",
				gas.NE(
					&gas.E{
						Tag: "div",
						Attrs: func() gas.Map {
							return gas.Map{
								"style": fmt.Sprintf("width: 24px; height: 18px; margin: 0 18px; border-radius: 4px; background-color: rgb(%d, %d, %d)", root.Range, 255-root.Range, root.Range),
							}
						},
					},
				),
				"  and range: ", root.Range,
			),
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Tag: "input",
					Handlers: map[string]gas.Handler {
						"input": func(event gas.Event) {
							root.Range, _ = strconv.Atoi(event.Value())
							go root.c.Update()
						},
					},
					Attrs: func() gas.Map {
						return gas.Map{
							"type": "range",
							"min":  "0",
							"max":  "255",
							"value": strconv.Itoa(root.Range),
						}
					},
				},
			),
		),
		gas.NE(
			&gas.E{
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "model__checkbox",
						"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
					}
				},
			},
			"Your checkbox: ", root.CheckBox,
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Tag: "input",
					Handlers: map[string]gas.Handler {
						"change": func(event gas.Event) {
							root.CheckBox = !root.CheckBox
							go root.c.Update()
						},
					},
					Attrs: func() gas.Map {
						return gas.Map{
							"type": "checkbox",
							"value": func() string {
								if root.CheckBox {
									return "true"
								}
	
								return "false"
							}(),
						}
					},
				},
			),
		),
	)
}
