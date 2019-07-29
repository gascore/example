package router

import "github.com/gascore/std/router"

var ctx *router.Ctx

func Ctx() *router.Ctx {
	return ctx
}

func SetCtx(newCtx *router.Ctx) {
	ctx = newCtx
} 
