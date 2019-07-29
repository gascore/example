package examples

import (
	"fmt"

	"github.com/gascore/gas"
)

// Example application #8
//
// 'to do' shows how you how to build basic to-do-mvc example
func TODO() *gas.E {
	root := &TODORoot{
		current: []string{},
		done:    []string{},
		deleted: []string{},
	}

	c := &gas.C{
		Root: root,
		Watchers: map[string]gas.Watcher{
			"newTask": func(val interface{}, e gas.Object) (string, error) {
				if val != nil {
					root.CurrentText = val.(string)
				}

				return root.CurrentText, nil
			},
		},
	}
	root.c = c

	return c.Init()
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
	root.c.UpdateWatchersValues("newTask", "")

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

func (root *TODORoot) Render() []interface{} {
	return gas.CL(
		gas.NE(
			&gas.E{
				Attrs: map[string]string{
					"id": "todo-wrap",
				},
			},
			gas.NE(
				&gas.E{
					Tag:   "style",
					Attrs: map[string]string{"type": "text/css"},
					HTML: gas.HTMLDirective{
						Render: func() string {
							return styles
						},
					},
				},
			),
			gas.NE(
				&gas.E{
					Attrs: map[string]string{
						"id": "todo-main",
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
					&gas.E{
						Watcher: "newTask",
						Tag:     "input",
						Handlers: map[string]gas.Handler{
							"keyup.enter": func(e gas.Object) {
								root.Add()
							},
						},
						Attrs: map[string]string{
							"id":          "todo-new",
							"placeholder": "New task",
						},
					},
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
					"Double-click to edit a task"),
				gas.NE(
					&gas.E{},
					"Created by",
					gas.NE(
						&gas.E{
							Tag: "a",
							Attrs: map[string]string{
								"href":   "https://noartem.github.io/",
								"target": "_blank",
							},
						},
						"Noskov Artem"),
					"with",
					gas.NE(
						&gas.E{
							Tag: "a",
							Attrs: map[string]string{
								"href":   "https://gascore.github.io",
								"target": "_blank",
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
			Attrs: map[string]string{
				"class": "list",
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
				"click": func(e gas.Object) {
					fmt.Println(e.GetInt("x"), e.GetInt("y"))
					root.ChangeCurrent(index)
				},
			},
			Binds: map[string]gas.Bind{
				"class": func() string {
					if current == index {
						return "active"
					}
					return ""
				},
			},
		},
		name)
}
