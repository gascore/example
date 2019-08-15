package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

func getHooker() *gas.E {
	root := &hooker{
		Show: true,
	}

	c := &gas.C{
		Root: root,
		Hooks: gas.Hooks{
			BeforeCreated: func() (bool, error) {
				fmt.Println("Hooker: BeforeCreated")
				return false, nil
			},
			Created: func() error {
				fmt.Println("Hooker: Created")
				return nil
			},
			Mounted: func() error {
				fmt.Println("Hooker: Mounted")
				return nil
			},
			BeforeUpdate: func() error {
				fmt.Println("Hooker: BeforeUpdate")
				return nil
			},
			Updated: func() error {
				root.Counter++
				fmt.Println("Hooker: Updated")
				return nil
			},
			BeforeDestroy: func() error {
				fmt.Println("Hooker: BeforeDestroy")
				return nil
			},
		},
	}

	root.c = c
	return c.Init()
}

type hooker struct {
	c *gas.C

	Show    bool
	Counter int
}

func (root *hooker) Render() []interface{} {
	return gas.CL(
		"You have updated app state ", root.Counter, "th times",
		gas.NE(
			&gas.E{
				Handlers: map[string]gas.Handler{
					"click": func(e gas.Event) {
						root.Show = !root.Show
						go root.c.Update()
					},
				},
				Tag: "button",
				Attrs: func() gas.Map {
					return gas.Map{
						"id":    "hooks__button",
						"class": "button outline",
						"style": "margin: 0 .4em;",
					}
				},
			},
			func() interface{} {
				if root.Show {
					return gas.NE(
						&gas.E{
							Tag: "i",
						},
						"Show text",
					)
				} else {
					return gas.NE(
						&gas.E{
							Tag: "b",
						},
						"Hide text",
					)
				}
			}(),
		),
		func() interface{} {
			if root.Show {
				return "Visible text"
			} else {
				return "Hidden text"
			}
		}(),
	)
}
