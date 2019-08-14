package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

func FunctionalExample() *gas.E {
	f := &gas.F{}

	counter, setCounter := f.UseStateInt(0)
	msg, setMsg := f.UseStateString("")

	// f.UseEffect(func() (func(), error) {
	// 	fmt.Println("UseEffect", counter(), msg())
	// 	return func() {
	// 		fmt.Println("cleaner")
	// 	}, nil
	// })

	return f.Init(false, func() []interface{} {
		return gas.CL(
			gas.NE(
				&gas.E{},
				gas.NE(
					&gas.E{
						Tag: "button",
						Handlers: map[string]gas.Handler{
							"click": func(e gas.Event) { setCounter(counter() + 1) },
						},
						Attrs: func() gas.Map {
							return gas.Map{
								"class": "button outline",
							}
						},
					},
					"+",
				),
				counter(),
				gas.NE(
					&gas.E{
						Tag: "button",
						Handlers: map[string]gas.Handler{
							"click": func(e gas.Event) { setCounter(counter() - 1) },
						},
						Attrs: func() gas.Map {
							return gas.Map{
								"class": "button outline",
							}
						},
					},
					"-",
				),
			),
			gas.NE(
				&gas.E{},
				gas.NE(
					&gas.E{
						Tag: "button",
						Handlers: map[string]gas.Handler{
							"click": func(e gas.Event) {
								setMsg(msg() + fmt.Sprintf("%d", counter()))
							},
						},
						Attrs: func() gas.Map {
							return gas.Map{
								"class": "button outline",
							}
						},
					},
					"add",
				),
				msg(),
			),
		)
	})
}
