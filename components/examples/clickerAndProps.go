package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// ClickerAndProps Example application #2
//
// 'clicker&props' shows how you can add handlers, change component.Data and use external components
func ClickerAndProps() *gas.C {
	return gas.NC(
		&gas.C{
			Data: map[string]interface{}{
				"click": 0,
			},
			Methods: map[string]gas.Method{
				"addClick": func(this *gas.C, i ...interface{}) (interface{}, error) {
					currentClick := this.Get("click").(int)
					this.SetValue("click", currentClick+1)
					return nil, nil
				},
			},
			Attrs: map[string]string{
				"id": "clicker&props",
			},
		},
		func(this *gas.C) []interface{} {
			return gas.CL(
				gas.NE(
					&gas.C{
						Handlers: map[string]gas.Handler{
							"click.left": func(this2 *gas.C, e gas.Object) {
								this.Method("addClick")
							},
							// you need to click button once (for target it)
							"keyup.control": func(this2 *gas.C, e gas.Object) {
								this.Method("addClick")
							},
							"keyup.a": func(this2 *gas.C, e gas.Object) {
								this.Method("addClick")
							},
							"keyup.s": func(this2 *gas.C, e gas.Object) {
								this.Method("addClick")
							},
							"keyup.d": func(this2 *gas.C, e gas.Object) {
								this.Method("addClick")
							},
							"keyup.f": func(this2 *gas.C, e gas.Object) {
								this.Method("addClick")
							},
						},
						Tag: "button",
						Attrs: map[string]string{
							"id": "clicker__button", // I love BEM
						},
					},
					"Click me!"),
				gas.NE(
					&gas.C{
						Tag: "i",
						Attrs: map[string]string{
							"id": "needful_wrapper",
						},
					},
					"You clicked button: ",
					CAPGetNumberViewer(this.Get("click").(int))))
		})
}

// CAPGetNumberViewer return very cool number viewer.
// It can be in another directory too.
// For reference from not parent component you can use `values` (they will reload).
func CAPGetNumberViewer(click int) interface{} {
	return gas.NC(
		&gas.Component{
			Tag: "i",
			Attrs: map[string]string{
				"id": "needful_wrapper--number-viewer",
			},
		},
		func(this *gas.Component) []interface{} {
			return gas.CL(
				fmt.Sprintf("%d times", click))
		})
}
