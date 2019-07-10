package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// Example application #5
//
// 'watchersAndBinds' shows how you can use component.Watchers and element.Binds for create forms
func WatchersAndBinds() *gas.E {
	root := &WathcersAndBindsRoot{
		Text:     "",
		Color:    "#000",
		Range:    0,
		CheckBox: false,
	}

	c := &gas.C{
		Root: root,
		Watchers: map[string]gas.Watcher{
			"text": func(val interface{}, e gas.Object) (string, error) {
				if val == nil { // default value
					return root.Text, nil
				}

				root.Text = val.(string)
				return root.Text, nil
			},
			"color": func(val interface{}, e gas.Object) (string, error) {
				if val == nil { // default value
					return root.Color, nil
				}

				root.Color = val.(string)
				return root.Color, nil
			},
			"range": func(val interface{}, e gas.Object) (string, error) {
				if val == nil { // default value
					return fmt.Sprintf("%d", root.Range), nil
				}

				root.Range = val.(int)
				return fmt.Sprintf("%d", root.Range), nil
			},
			"checkbox": func(val interface{}, e gas.Object) (string, error) {
				if val == nil { // default value
					return "false", nil
				}

				root.CheckBox = val.(bool)
				return fmt.Sprintf("%t", root.CheckBox), nil
			},
		},
	}
	root.c = c

	return c.Init()
}

type WathcersAndBindsRoot struct {
	c *gas.C

	Text     string
	Color    string
	Range    int
	CheckBox bool
}

func (root *WathcersAndBindsRoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id":    "model__text",
					"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
				},
			},
			"Your text: ", root.Text,
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Watcher: "text",
					Tag:     "input",
				},
			),
		),
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id":    "model__color",
					"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
				},
			},
			"Your color: ",
			gas.NE(
				&gas.E{
					Tag: "span",
					Attrs: map[string]string{
						"style": "color:" + root.Color,
					},
				},
				root.Color,
			),
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Watcher: "color",
					Tag:     "input",
					Attrs: map[string]string{
						"type": "color",
					},
				},
			),
		),
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id":    "model__range",
					"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
				},
			},
			gas.NE(
				&gas.E{
					Tag: "span",
					Attrs: map[string]string{
						"style": "display: flex;",
					},
				},
				"Generated color: ",
				gas.NE(
					&gas.E{
						Tag: "div",
						Attrs: map[string]string{
							"style": fmt.Sprintf("width: 24px; height: 18px; margin: 0 18px; border-radius: 4px; background-color: rgb(%d, %d, %d)", root.Range, 255-root.Range, root.Range),
						},
					},
				),
				"  and range: ", root.Range,
			),
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Watcher: "range",
					Tag:     "input",
					Attrs: map[string]string{
						"type": "range",
						"min":  "0",
						"max":  "255",
					},
				},
			),
		),
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id":    "model__checkbox",
					"style": "border: 1px solid #dedede; margin-bottom: 8px; padding: 4px 16px;",
				},
			},
			"Your checkbox: ", root.CheckBox,
			gas.NE(&gas.E{Tag: "br"}),
			gas.NE(
				&gas.E{
					Watcher: "checkbox",
					Tag:     "input",
					Attrs: map[string]string{
						"type": "checkbox",
					},
				},
			),
		),
	)
}
