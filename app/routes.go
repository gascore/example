package main

import (
	"fmt"

	"github.com/gascore/dom"
	"github.com/gascore/std/router"
	
	c "github.com/gascore/example/app/components"
	appRouter "github.com/gascore/example/app/router"
)

func InitRouter() {
	ctx := &router.Ctx{
		Routes: []router.Route{
			{
				Name:    "home",
				Element: c.Home,
				Path:    "/",
				Exact:   true,
			},
			{
				Name:    "about",
				Element: c.About,
				Path:    "/about",
				Exact:   true,
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
				Element: c.Link,
				Path:    "/links/:name",
			},
			{
				Name:    "links",
				Element: c.Links,
				Path:    "/links",
				Exact:   true,
			},
			{
				Name:    "examples",
				Element: c.Examples,
				Path:    "/examples",
				Exact:   true,
				Childes: []router.Route{
					{
						Name:    "components",
						Element: c.Components,
						Path:    "/components",
						Exact:   true,
						Before: func(info *router.MiddlewareInfo) (bool, error) {
							fmt.Println("Before components", info)
							return false, nil
						},
					},
					{
						Name: "c",
						Path: "/c",
						RedirectName: "components",
					},
				},
				Before: func(info *router.MiddlewareInfo) (bool, error) {
					fmt.Println("Before examples", info)
					return false, nil
				},
			},
			{
				Name:           "todo-redirect",
				Exact:          true,
				Path:           "/todo",
				RedirectName: 	"todo-list",
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
				Element: c.TodoList,
				Path:    "/todo/:type",
			},
		},

		After: func(to, from *router.RouteInfo) error {
			dom.GetWindow().JSValue().Call("scrollTo", 0, 0)
			return nil
		},
	
		Settings: router.Settings{
			HashMode: true,
			GetUserConfirmation: func() bool {
				return false
			},
			ForceRefresh: !router.SupportHistory(),
		},
	}
	ctx.Init()
	appRouter.SetCtx(ctx)
}