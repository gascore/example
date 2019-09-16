package examples

import "github.com/gascore/gas"

// Example application #8
//
// 'to do' shows how you how to build basic to-do-mvc example
func TODO() *gas.C {
	root := &TODORoot{
		current: []string{},
		done:    []string{},
		deleted: []string{},
	}

	c := &gas.C{
		Root: root,
		NotPointer: true,
	}
	root.c = c

	return c
}

type TODORoot struct {
	c *gas.C

	CurrentList int
	CurrentText string

	current []string
	done    []string
	deleted []string
}

func (root *TODORoot) Delete(i int) {
	deletedItem := root.current[i]
	root.current = append(root.current[:i], root.current[i+1:]...)

	root.deleted = append(root.deleted, deletedItem)

	go root.c.Update()
}

func (root *TODORoot) Add() {
	if root.CurrentText == "" {
		return
	}

	root.current = append(root.current, root.CurrentText)
	root.CurrentList = 0

	root.CurrentText = ""

	go root.c.Update()
}

func (root *TODORoot) MarkAsDone(i int) {
	doneItem := root.current[i]
	root.current = append(root.current[:i], root.current[i+1:]...)

	root.done = append(root.done, doneItem)

	go root.c.Update()
}

func (root *TODORoot) Edit(i int, newValue string) {
	root.current[i] = newValue
	go root.c.Update()
}

func (root *TODORoot) ChangeCurrent(newCurrent int) {
	root.CurrentList = newCurrent
	go root.c.Update()
}

func (root *TODORoot) Render() *gas.E {
	return gas.NE(
		&gas.E{},
		gas.NE(
			&gas.E{
				Attrs: func() gas.Map {
					return gas.Map{
						"id": "todo-wrap",
					}
				},
			},
			gas.NE(
				&gas.E{
					Tag: "style",
					Attrs: func() gas.Map {
						return gas.Map{"type": "text/css"}
					},
					HTML: func() string {
						return styles
					},
				},
			),
			gas.NE(
				&gas.E{
					Attrs: func() gas.Map {
						return gas.Map{
							"id": "todo-main",
						}
					},
				},
				gas.NE(
					&gas.E{
						Tag: "nav",
					},
					getNavEl(0, root.CurrentList, "Current", root),
					getNavEl(1, root.CurrentList, "Completed", root),
					getNavEl(2, root.CurrentList, "Deleted", root),
				),
				gas.NE(
					&gas.E{},
					gas.NE(
						&gas.E{
							Tag: "input",
							Handlers: map[string]gas.Handler{
								"keyup.enter": func(event gas.Event) {
									root.Add()
								},
								"input": func(event gas.Event) {
									root.CurrentText = event.Value()
									go root.c.Update()
								},
							},
							Attrs: func() gas.Map {
								return gas.Map{
									"id":          "todo-new",
									"placeholder": "New task",
									"value":       root.CurrentText,
									"class":       "form-input todo-input",
								}
							},
						},
					),
					gas.NE(
						&gas.E{
							Tag: "button",
							Handlers: map[string]gas.Handler{
								"click": func(event gas.Event) {
									root.Add()
								},
							},
							Attrs: func() gas.Map {
								return gas.Map{
									"class": "button outline",
								}
							},
						},
						"Add",
					),
					func() interface{} {
						if root.CurrentList == 0 {
							return gas.NE(
								&gas.E{
									Tag: "i",
									Attrs: func() gas.Map {
										return gas.Map{
											"style": "color: gray;font-size: 12px;margin-left:6px;",
										}
									},
								},
								"Double-click to edit a task",
							)
						}

						return nil
					}(),
				),
				func() interface{} {
					switch root.CurrentList {
					case 0:
						return getList(0, root.current, root)
					case 1:
						return getList(1, root.done, root)
					case 2:
						return getList(2, root.deleted, root)
					default:
						return nil
					}
				}(),
			),
			gas.NE(
				&gas.E{
					Tag: "footer",
				},
				gas.NE(
					&gas.E{},
					"Created by",
					gas.NE(
						&gas.E{
							Tag: "a",
							Attrs: func() gas.Map {
								return gas.Map{
									"href":   "https://noartem.github.io/",
									"target": "_blank",
								}
							},
						},
						"Noskov Artem"),
					"with",
					gas.NE(
						&gas.E{
							Tag: "a",
							Attrs: func() gas.Map {
								return gas.Map{
									"href":   "https://gascore.github.io",
									"target": "_blank",
								}
							},
						},
						"GAS framework"),
					"and love",
				),
			),
		),
	)
}

func getList(index int, list []string, root dataForLi) interface{} {
	return gas.NE(
		&gas.E{
			Tag: "ul",
			Attrs: func() gas.Map {
				return gas.Map{
					"class": "list",
				}
			},
		},
		func() []interface{} {
			var elements []interface{}
			for i, el := range list {
				elements = append(elements, gas.NE(
					&gas.E{
						Tag: "li",
					},
					getLi(index, i, el, root),
				))
			}
			return elements
		}(),
	)
}

func getNavEl(index, current int, name string, root interface{ ChangeCurrent(int) }) interface{} {
	return gas.NE(
		&gas.E{
			Tag: "button",
			Handlers: map[string]gas.Handler{
				"click": func(e gas.Event) {
					root.ChangeCurrent(index)
				},
			},
			Attrs: func() gas.Map {
				return gas.Map{
					"class": func() string {
						if current == index {
							return "active button outline"
						}
						return "button outline"
					}(),
				}
			},
		},
		name)
}
