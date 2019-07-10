package main

import (
	r "github.com/gascore/std/router"
	c "github.com/gascore/wow/components"
)

var ctx = &r.Ctx{
	Routes: []r.Route{
		{
			Name:      "home",
			Element:   c.Home,
			Path:      "/",
			Exact:     true,
		},
		{
			Name:      "about",
			Element:   c.About,
			Path:      "/about",
			Exact:     true,
		},
		{
			Name:      "components",
			Element:   c.Components,
			Path:      "/components",
			Exact:     true,
		},
		{
			Name:      "examples",
			Element:   c.Examples,
			Path:      "/examples",
			Exact:     true,
		},
		// {
		// 	Name:           "todo-redirect",
		// 	Exact:          true,
		// 	Path:           "/todo",
		// 	Redirect:       "todo-list",
		// 	RedirectParams: map[string]string{"type": "active"},
		// },
		// {
		// 	Name:     "todo-redirect",
		// 	Exact:    true,
		// 	Path:     "/todo/",
		// 	Redirect: "/todo/active",
		// },
		{
			Name:    "todo-list",
			Element: c.TodoList,
			Path:    "/todo/:type",
		},
	},

	Settings: r.Settings{
		HashMode: true,
		GetUserConfirmation: func() bool {
			return false
		},
		ForceRefresh: !r.SupportHistory(),
	},
}