package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// Watchers application #8
//
// 'watchers' shows how you can use component.Watchers
func Watchers() *gas.C {
	return gas.NC(
		&gas.Component{
			Data: map[string]interface{}{
				"show":               true,
				"watcherIsTriggered": false,
			},
			Watchers: map[string]gas.Watcher{
				"show": func(this *gas.Component, new interface{}, old interface{}) error {
					this.ConsoleLog(fmt.Sprintf("Watcher is triggered! New value: %t, old value: %t", new, old))

					err := this.SetValueImm("watcherIsTriggered", true)
					if err != nil {
						this.WarnError(err)
						return err
					}

					return nil
				},
			},
			Attrs: map[string]string{
				"id": "if",
			},
		},
		func(this *gas.Component) []interface{} {
			return gas.CL(
				gas.NE(
					&gas.Component{
						Handlers: map[string]gas.Handler{
							"click": func(c *gas.Component, e gas.Object) {
								this.SetValue("show", !this.Get("show").(bool))
							},
						},
						Tag: "button",
						Attrs: map[string]string{
							"id": "if__button",
						},
					},
					gas.NE(
						&gas.C{
							Show: func(c *gas.C) bool {
								return this.Get("show").(bool)
							},
							Tag: "i",
						},
						"Show text"),
					gas.NE(
						&gas.C{
							Show: func(c *gas.C) bool {
								return !this.Get("show").(bool)
							},
							Tag: "i",
						},
						"Hidden text",
					),
				),
				gas.NE(
					&gas.Component{
						If: func(c *gas.Component) bool {
							return !this.Get("show").(bool)
						},
						Tag: "i",
					},
					"Hidden text",
				),
				gas.NE(
					&gas.Component{
						If: func(this2 *gas.Component) bool {
							watcherIsTriggered, ok := this.Get("watcherIsTriggered").(bool)
							this.WarnIfNot(ok)
							return watcherIsTriggered
						},
						Tag: "strong",
						Attrs: map[string]string{
							"style": "color: red;",
						},
					},
					"Watcher is triggered!",
				),
			)
		})
}
