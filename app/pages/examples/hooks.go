package examples

import "github.com/gascore/gas"

// Example application #7
//
// 'hooks' shows how you can use component.Hooks
func Hooks() *gas.C {
	root := &HooksRoot{}
	c := &gas.C{
		NotPointer: true,
		Root: root,
	}

	return c
}

type HooksRoot struct {
}

func (root *HooksRoot) Render() *gas.E {
	return gas.NE(&gas.E{}, getHooker())
}
