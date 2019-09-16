package examples

import "github.com/gascore/gas"

type clickerExternalComponent struct {
	click int
}

func (data *clickerExternalComponent) Render() *gas.E {
	return gas.NE(
		&gas.E{},
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

func newNumberViewer(click int) *gas.C {
	ec := &clickerExternalComponent{
		click: click,
	}

	c := &gas.C{
		NotPointer: true,
		Root:       ec,
	}

	return c
}
