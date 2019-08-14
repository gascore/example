package examples

import "github.com/gascore/gas"

// Example application #2
//
// 'clicker' shows how you can add handlers and use external components
func Clicker() *gas.E {
	root := &ClickerRoot{
		click: 0,
	}

	c := &gas.C{
		Root: root,
	}

	root.c = c

	return c.Init()
}

type ClickerRoot struct {
	c     *gas.C
	click int
}

func (root *ClickerRoot) addClick() {
	root.click++
	root.c.Update()
}

func (root *ClickerRoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Handlers: map[string]gas.Handler{
					"click.left": func(e gas.Event) {
						root.addClick()
					},
					"keyup.control": func(e gas.Event) {
						root.addClick()
					},
					"keyup.a": func(e gas.Event) {
						root.addClick()
					},
					"keyup.s": func(e gas.Event) {
						root.addClick()
					},
					"keyup.d": func(e gas.Event) {
						root.addClick()
					},
					"keyup.f": func(e gas.Event) {
						root.addClick()
					},
				},
				Tag: "button",
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "clicker__button",
						"class": "button outline",
					}
				},
			},
			"Click me!"),
		newNumberViewer(root.click))
}
