package examples

import "github.com/gascore/gas"

type dataForLi interface {
	MarkAsDone(int)
	Delete(int)
	Edit(int, string)
}

type listItem struct {
	c *gas.C

	data dataForLi

	listIndex int

	isEditing bool

	index int
	value string
}

func (root *listItem) Render() []interface{} {
	return gas.CL(
		func() interface{} {
			if root.listIndex != 0 {
				return nil
			}
			return gas.NE(
				&gas.E{
					Tag: "button",
					Handlers: map[string]gas.Handler{
						"click": func(e gas.Event) {
							root.data.MarkAsDone(root.index)
						},
					},
					Attrs: func() map[string]string {
						return map[string]string{
							"id": "submit",
						}
					},
				},
				gas.NE(
					&gas.E{
						Tag: "i",
						Attrs: func() map[string]string {
							return map[string]string{
								"class": "icon icon-check",
							}
						},
					},
				),
			)
		}(),

		func() interface{} {
			if root.isEditing {
				return gas.NE(
					&gas.E{
						Tag: "input",
						Attrs: func() map[string]string {
							return map[string]string{
								"style": "margin-right: 8px",
								"value": root.value,
								"class": "form-input",
							}
						},
						Handlers: map[string]gas.Handler{
							"input": func(event gas.Event) {
								root.value = event.Value()
								go root.c.Update()
							},
							"keyup.enter": func(e gas.Event) {
								root.isEditing = false
								root.data.Edit(root.index, root.value)
								go root.c.Update()
							},
						},
					},
				)
			} else {
				return gas.NE(
					&gas.E{
						Tag: "span",
						Handlers: map[string]gas.Handler{
							"dblclick": func(e gas.Event) {
								if root.listIndex != 0 {
									return
								}

								root.isEditing = true
								go root.c.Update()
							},
						},
					},
					root.value,
				)
			}
		}(),

		func() interface{} {
			if root.listIndex != 0 {
				return nil
			}
			return gas.NE(
				&gas.E{
					Tag: "button",
					Handlers: map[string]gas.Handler{
						"click": func(e gas.Event) {
							root.data.Delete(root.index)
						},
					},
					Attrs: func() map[string]string {
						return map[string]string{
							"id": "delete",
						}
					},
				},
				gas.NE(
					&gas.E{
						Tag: "i",
						Attrs: func() map[string]string {
							return map[string]string{
								"class": "icon icon-delete",
							}
						},
					},
				),
			)
		}(),
	)
}

func getLi(listIndex, index int, el string, data dataForLi) *gas.E {
	root := &listItem{
		listIndex: listIndex,

		index: index,
		value: el,

		data: data,
	}

	c := &gas.C{
		Root:       root,
		NotPointer: true,
	}
	root.c = c

	return c.Init()
}
