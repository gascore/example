package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

func FunctionalExample() *gas.E {
	f := &gas.F{}

	counter, setCounter := f.UseStateInt(0)
	msg, setMsg := f.UseStateString("")

	f.UseEffect(func() error {
		fmt.Println("UseEffect", counter(), msg())
		return nil
	})

	return f.Init(false, func() []interface{} {
		return gas.CL(
			gas.NE(
				&gas.E{},
				gas.NE(&gas.E{Tag: "button", Handlers: map[string]gas.Handler{"click": func(e gas.Object) { setCounter(counter() + 1) }}}, "+"),
				counter(),
				gas.NE(&gas.E{Tag: "button", Handlers: map[string]gas.Handler{"click": func(e gas.Object) { setCounter(counter() - 1) }}}, "-"),
			),
			gas.NE(
				&gas.E{},
				gas.NE(&gas.E{Tag: "button", Handlers: map[string]gas.Handler{"click": func(e gas.Object) {
					setMsg(msg() + fmt.Sprintf("%d", counter()))
				}}}, "add"),
				msg(),
			),
		)
	})
}
