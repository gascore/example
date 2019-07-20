package examples

import "github.com/gascore/gas"

func FunctionalExample() *gas.E {
	f := gas.NFC(true)

	getCounter, setCounter := f.UseState(0)

	return f.Init(func()[]interface{} {return gas.CL(
		gas.NE(&gas.E{Tag: "button", Handlers: map[string]gas.Handler{"click": func(e gas.Object) { setCounter(getCounter().(int) + 1) }}}, "+"),
		getCounter(),
		gas.NE(&gas.E{Tag: "button", Handlers: map[string]gas.Handler{"click": func(e gas.Object) { setCounter(getCounter().(int) - 1) }}}, "-"),
	)})
}
