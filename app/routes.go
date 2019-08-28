package main

import (
	"fmt"

	"github.com/gascore/dom"
	"github.com/gascore/std/router"

	"github.com/gascore/example/app/pages"
)

func InitRouter() *router.Ctx {
	p := pages.NewPages()

	ctx := &router.Ctx{
		Routes: []router.Route{
			{
				Name:    "home",
				Element: p.Home,
				Path:    "/",
				Exact:   true,
			},
			{
				Name:    "about",
				Element: p.About,
				Path:    "/about",
				Exact:   true,
			},
			{
				Name:    "markdown",
				Element: p.Markdown,
				Path:    "/markdown",
				Exact:   true,
			},
			{
				Name:     "md-redirect",
				Path:     "/md",
				Exact:    true,
				Redirect: "/markdown",
			},
			{
				Name:     "links-redirect",
				Exact:    true,
				Path:     "/link",
				Redirect: "/links",
			},
			{
				Name:     "link-redirect",
				Exact:    true,
				Path:     "/links/",
				Redirect: "/links",
			},
			{
				Name:    "link",
				Element: p.Link,
				Path:    "/links/:name",
			},
			{
				Name:    "links",
				Element: p.Links,
				Path:    "/links",
				Exact:   true,
			},
			{
				Name:    "examples",
				Element: p.Examples,
				Path:    "/examples",
				Exact:   true,
				Childes: []router.Route{
					{
						Name:    "components",
						Element: p.Components,
						Path:    "/components",
						Exact:   true,
						Before: func(info *router.MiddlewareInfo) (bool, error) {
							fmt.Println("Router middleware: Before components", info)
							return false, nil
						},
					},
					{
						Name:         "c",
						Path:         "/c",
						RedirectName: "components",
					},
				},
				Before: func(info *router.MiddlewareInfo) (bool, error) {
					fmt.Println("Router middleware: Before examples", info)
					return false, nil
				},
			},
			{
				Name:           "todo-redirect",
				Exact:          true,
				Path:           "/todo",
				RedirectName:   "todo-list",
				RedirectParams: map[string]string{"type": "active"},
			},
			{
				Name:     "todo-redirect",
				Exact:    true,
				Path:     "/todo/",
				Redirect: "/todo/active",
			},
			{
				Name:    "todo-list",
				Element: p.TodoList,
				Path:    "/todo/:type",
			},
		},

		After: func(to, from *router.RouteInfo) error {
			dom.GetWindow().JSValue().Call("scrollTo", 0, 0)
			return nil
		},

		Settings: router.Settings{
			NotFound: p.NotFound,
			HashMode: true,
			GetUserConfirmation: func() bool {
				return false
			},
			ForceRefresh: !router.SupportHistory(),
		},
	}
	ctx.Init()

	p.RCtx = ctx
	return ctx
}
