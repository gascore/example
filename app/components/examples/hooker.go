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
				fmt.Println("BeforeCreated")
				return false, nil
			},
			Created: func() error {
				fmt.Println("Created")
				return nil
			},
			Mounted: func() error {
				fmt.Println("Mounted")
				return nil
			},
			BeforeUpdate: func() error {
				fmt.Println("BeforeUpdate")
				return nil
			},
			Updated: func() error {
				root.Counter++
				fmt.Println("Updated")
				return nil
			},
			BeforeDestroy: func() error {
				fmt.Println("BeforeDestroy")
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
		fmt.Sprintf("You have updated app state %dth times", root.Counter),
		gas.NE(
			&gas.E{
				Handlers: map[string]gas.Handler{
					"click": func(e gas.Object) {
						root.Show = !root.Show
						go root.c.Update()
					},
				},
				Tag: "button",
				Attrs: map[string]string{
					"id": "hooks__button",
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
