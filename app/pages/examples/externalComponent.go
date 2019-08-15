package examples

import "github.com/gascore/gas"

type clickerExternalComponent struct {
	click int
}

func (data *clickerExternalComponent) Render() []interface{} {
	return gas.CL(
		"You've clicked button: ",
		gas.NE(
			&gas.E{
				Tag: "i",
				Attrs: func() gas.Map {
					return gas.Map{
						"id": "number-viewer",
					}
				},
			},
			data.click, " times",
		),
	)
}

func newNumberViewer(click int) interface{} {
	ec := &clickerExternalComponent{
		click: click,
	}

	c := &gas.C{
		NotPointer: true,
		Root:       ec,
	}

	return c.Init()
}
