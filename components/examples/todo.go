package examples

import (
	"errors"
	"fmt"

	"github.com/gascore/gas"
)

// TODO Example application #11
//
// 'to do' shows how you how to build basic to-do-mvc example
func TODO() *gas.C {
	return gas.NC(
		&gas.C{
			Data: map[string]interface{}{
				"currentList": "0",
				"currentText": "",

				"current": []interface{}{},
				"done":    []interface{}{},
				"deleted": []interface{}{},
			},
			Methods: map[string]gas.Method{
				"delete": func(this *gas.C, values ...interface{}) (interface{}, error) {
					i, ok := values[0].(int)
					if !ok {
						return nil, errors.New("invalid index")
					}

					appendToDeleted, ok := values[1].(bool)
					if !ok {
						return nil, errors.New("invalid appendToDeleted")
					}

					list, ok := this.Get("current").([]interface{})
					if !ok {
						return nil, errors.New("invalid current list")
					}
					removedItem := list[i]

					err := this.DataDeleteFromArray("current", i)
					if err != nil {
						return nil, err
					}

					this.SetValue("currentText", "")

					if appendToDeleted {
						_, err = this.MethodSafely("append", "deleted", removedItem)
						if err != nil {
							return nil, err
						}
					}

					return nil, nil
				},
				"append": func(this *gas.C, values ...interface{}) (interface{}, error) {
					listTypeS, ok := values[0].(string)
					if !ok {
						return nil, errors.New("invalid list type")
					}

					newTask, ok := values[1].(string)
					if !ok {
						return nil, errors.New("invalid task")
					}

					err := this.DataAddToArray(listTypeS, newTask)
					if err != nil {
						return nil, err
					}

					if listTypeS == "current" {
						this.SetValue("currentText", "")
					}

					return nil, nil
				},
				"markAsDone": func(this *gas.C, values ...interface{}) (interface{}, error) {
					i, ok := values[0].(int)
					if !ok {
						return nil, errors.New("invalid index")
					}

					list := this.Get("current").([]interface{})

					item := list[i]

					_, err := this.MethodSafely("append", "done", item)
					if err != nil {
						return nil, err
					}

					_, err = this.MethodSafely("delete", i, false)
					if err != nil {
						return nil, err
					}

					return nil, nil
				},
				"edit": func(this *gas.C, values ...interface{}) (interface{}, error) {
					i, ok := values[0].(int)
					if !ok {
						return nil, errors.New("invalid index")
					}

					newValue, ok := values[1].(string)
					if !ok {
						return nil, errors.New("invalid new value")
					}

					err := this.DataEditArray("current", i, newValue)
					if err != nil {
						return nil, err
					}

					return nil, nil
				},
			},
		},
		func(this *gas.C) []interface{} {
			return gas.CL(
				todoGetStyleEl(),
				gas.NE(
					&gas.C{
						Tag: "div",
						Attrs: map[string]string{
							"id": "todo",
						},
					},
					gas.NE(
						&gas.C{
							Tag: "nav",
						},
						todoGetNavEl(this, "0", "Current"),
						todoGetNavEl(this, "1", "Completed"),
						todoGetNavEl(this, "2", "Deleted")),
					gas.NE(
						&gas.C{
							If: func(p *gas.C) bool {
								return this.Get("currentList").(string) == "0"
							},
							Model: gas.ModelDirective{
								Data:      "currentText",
								Component: this,
							},
							Tag: "input",
							Handlers: map[string]gas.Handler{
								"keyup.enter": func(p *gas.C, e gas.Object) {
									currentText := this.Get("currentText").(string)
									if len(currentText) == 0 {
										return
									}

									this.Method("append", "current", currentText)
								},
							},
							Attrs: map[string]string{
								"id":          "new",
								"placeholder": "New task",
							},
						},
					),
					gas.NE(
						&gas.C{},
						todoGetList(this, 0),
						todoGetList(this, 1),
						todoGetList(this, 2))),
				gas.NE(
					&gas.C{
						Attrs: map[string]string{
							"id": "todo-footer",
						},
					},
					"Double-click to edit a task"))
		})
}

func todoGetList(pThis *gas.C, index int) interface{} {
	return gas.NE(
		&gas.C{
			Show: func(p *gas.C) bool {
				return pThis.Get("currentList") == fmt.Sprintf("%d", index)
			},
			Tag: "ul",
			Attrs: map[string]string{
				"id":    "list__current",
				"class": "list",
			},
		},
		gas.CL(todoGetLi(pThis, index)...))
}

func todoGetLi(pThis *gas.C, listType int) []interface{} {
	// listType: 0 - current, 1 - done, 2 - deleted tasks
	var listTypeS string
	switch listType {
	case 0:
		listTypeS = "current"
	case 1:
		listTypeS = "done"
	case 2:
		listTypeS = "deleted"
	}

	return gas.NewFor(listTypeS, pThis, func(i interface{}, el interface{}) interface{} {
		return gas.NC(
			&gas.C{
				Tag: "li",
				Data: map[string]interface{}{
					"isEditing": false,
					"newValue":  "no",
				},
			},
			func(this *gas.C) []interface{} {
				return gas.CL(
					gas.NE(
						&gas.C{
							Tag: "button",
							If: func(p *gas.C) bool {
								return listType == 0
							},
							Handlers: map[string]gas.Handler{
								"click": func(this5 *gas.C, e gas.Object) {
									pThis.Method("markAsDone", i)
								},
							},
							Attrs: map[string]string{
								"id": "submit",
							},
						},
						gas.NE(
							&gas.C{
								Tag: "i",
								Attrs: map[string]string{
									"class": "icon icon-check",
								},
							})),
					gas.NE(
						&gas.C{
							Tag: "span",
							If: func(p *gas.C) bool {
								return !this.Get("isEditing").(bool)
							},
							Handlers: map[string]gas.Handler{
								"dblclick": func(p *gas.C, e gas.Object) {
									if listType != 0 {
										return
									}

									this.Set(map[string]interface{}{
										"newValue":  el,
										"isEditing": true,
									})
								},
							},
						},
						fmt.Sprintf("%s", el)),
					gas.NE(
						&gas.C{
							Tag: "input",
							Attrs: map[string]string{
								"style": "margin-right: 8px",
							},
							If: func(p *gas.C) bool {
								return this.Get("isEditing").(bool)
							},
							Model: gas.ModelDirective{
								Component: this,
								Data:      "newValue",
							},
							Handlers: map[string]gas.Handler{
								"keyup.enter": func(p *gas.C, e gas.Object) {
									newValue := this.Get("newValue")

									this.SetValue("isEditing", false)
									pThis.Method("edit", i, newValue)
									el = newValue
								},
							},
						},
						fmt.Sprintf("%s", el)),
					gas.NE(
						&gas.C{
							Tag: "button",
							If: func(p *gas.C) bool {
								return listType == 0
							},
							Handlers: map[string]gas.Handler{
								"click": func(this5 *gas.C, e gas.Object) {
									pThis.Method("delete", i, true)
								},
							},
							Attrs: map[string]string{
								"id": "delete",
							},
						},
						gas.NE(
							&gas.C{
								Tag: "i",
								Attrs: map[string]string{
									"class": "icon icon-delete",
								},
							})))
			})
	})
}

func todoGetStyleEl() interface{} {
	return gas.NE(
		&gas.C{
			Tag:   "style",
			Attrs: map[string]string{"type": "text/css"},
			HTML: gas.HTMLDirective{
				Render: func(this2 *gas.C) string {
					return `
#todo {
	border: 1px solid #dedede;
	border-radius: 4px;
	padding: 0px 0px 4px 0px;
}

#todo ul {
	padding: 0 16px;
	list-style-type: none;
	margin-left: 0;
}

#todo ul li {
	display: flex;
	padding: 4px 8px;
	border-bottom: 1px solid #dedede;

	font-size: 18px;
}

#todo ul li button {
	border: 0;
	padding: 0;
	background-color: inherit;
	cursor: pointer;
}
#todo ul li button#submit:hover, button#submit:focus {
	color: #009966;
}
#todo ul li button#delete:hover, button#delete:focus {
	color: #ff0033;
}

#todo ul li button#submit {
	margin: 0 12px 0 0;
}

#todo ul li button#delete {
	margin: 0 0 0 auto;
}

#todo nav {
	padding: 6px 16px;
	margin-bottom: 8px;
	border-bottom: 1px solid #dedede;
	background-color: #f1f1f1;
}

#todo nav button {
	margin-right: 6px;
	border: 0;
	padding: 0;
	color: #009966;
	background-color: inherit;
	cursor: pointer;
}
#todo nav button:focus, nav button:hover {
	color: #00CC99;
}
#todo nav button.active {
	text-decoration: underline;
}

#todo #new {
	margin: 0 16px;
}

#todo-footer {
	margin-top: 18px;
	color: gray;
	font-size: 12px;
	text-align: center;
}

`
				},
			},
		})
}

func todoGetNavEl(this *gas.C, index, name string) interface{} {
	return gas.NE(
		&gas.C{
			Tag: "button",
			Handlers: map[string]gas.Handler{
				"click": func(p *gas.C, e gas.Object) {
					this.ConsoleLog(e.GetInt("x"), e.GetInt("y"))
					this.SetValue("currentList", index)
				},
			},
			Binds: map[string]gas.Bind{
				"class": func() string {
					if this.Get("currentList").(string) == index {
						return "active"
					}
					return ""
				},
			},
		},
		name)
}
