package examples

import "github.com/gascore/gas"

// Example application #7
//
// 'hooks' shows how you can use component.Hooks
func Hooks() *gas.E {
	root := &HooksRoot{}
	c := &gas.C{Root: root}

	return c.Init()
}

type HooksRoot struct {
}

func (root *HooksRoot) Render() []interface{} {
	return gas.CL(
		getHooker(),
	)
}
