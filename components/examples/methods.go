package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// Methods Example application #5
//
// 'methods' shows how you can use component.Methods
func Methods() *gas.C {
	return gas.NC(
		&gas.Component{
			Data: map[string]interface{}{
				"show":   true,
				"number": 1,
			},
			// What the difference between Methods and Computed?
			// Methods will do business things.
			// Computed will return value from data, libraries, /dev/random, e.t.c. with some changes (or just raw)
			Methods: map[string]gas.Method{
				"toggle": func(this *gas.Component, values ...interface{}) (interface{}, error) {
					this.SetValue("show", !this.Get("show").(bool))

					if this.Get("show").(bool) {
						this.SetValue("number", this.Get("number").(int)+1)
					}

					return nil, nil
				},
				"number": func(this *gas.Component, values ...interface{}) (interface{}, error) {
					this.ConsoleLog(fmt.Sprintf("Some values: %s", values[0].(string)))

					currentNumber, ok := this.Get("number").(int)
					this.WarnIfNot(ok) // it's good practise to your data for valid type
					explanation := fmt.Sprintf("You showed hidden text: %d times", currentNumber)
					return explanation, nil
				},
			},
			Attrs: map[string]string{
				"id": "M&C",
			},
		},
		func(this *gas.Component) []interface{} {
			return gas.CL(
				methodsGetButton(this.Get("show").(bool), this.PocketMethod("toggle")),
				methodsGetHiddenText(this.Get("show").(bool), this.PocketMethod("number")))
		})
}

func methodsGetButton(show bool, toggleMethod gas.PocketMethod) *gas.Component {
	return gas.NE(
		&gas.Component{
			Handlers: map[string]gas.Handler{
				"click": func(this *gas.Component, e gas.Object) {
					_, err := toggleMethod()
					this.WarnError(err)
				},
			},
			Tag: "button",
			Attrs: map[string]string{
				"id": "M&C__button",
			},
		},
		gas.NE(
			&gas.C{
				If: func(p *gas.C) bool {
					return show
				},
			},
			"Show text"),
		gas.NE(
			&gas.C{
				If: func(p *gas.C) bool {
					return !show
				},
			},
			"Hide text"))
}

func methodsGetHiddenText(show bool, getNumber gas.PocketMethod) *gas.Component {
	return gas.NC(
		&gas.Component{
			If: func(c *gas.Component) bool {
				return !show
			},
			Tag: "i",
		},
		func(this *gas.Component) []interface{} {
			n, err := getNumber("something for computed")
			this.WarnError(err)

			return []interface{}{
				"Hidden text",
				fmt.Sprintf("  (%s)", n),
			}
		},
	)
}
