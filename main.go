package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/pflag"

	"github.com/gascore/gasx"
	"github.com/gascore/gasx/acss"
	"github.com/gascore/gasx/html"
)

func main() {
	var (
		isDev     = pflag.BoolP("dev", "d", true, "Development mode")
		platform  = pflag.StringP("platform", "p", "wasm", "Target platform")
		watch     = pflag.BoolP("watch", "w", false, "Watch for file changings")
		serveDist = pflag.BoolP("serve", "s", true, "Serve ./dist directory")
	)
	pflag.Parse()

	build(*platform, *isDev)

	currentDir, err := os.Getwd()
	gasx.Must(err)

	if *watch {
		go serve(currentDir)
		gasx.Must(gasx.StartWatcher(
			func(path string) {
				fmt.Println(color.GreenString("Triggered:"), strings.TrimPrefix(path, currentDir+"/app/"))
				build(*platform, *isDev)
			},
			[]string{"_gas.go"},
			[]string{},
			[]string{"app/"},
		))
	}

	if *serveDist {
		serve(currentDir)
	}
}

func build(platform string, isDev bool) {
	htmlCompiler := html.NewCompiler()

	acssGen := acss.Generator{
		BreakPoints: map[string]string{
			"lg": "@media(min-width:1200px)",
			"md": "@media(min-width:1000px)",
			"sm": "@media(min-width:750px)",
		},
		Custom: map[string]string{
			"b": "1px solid #dedede",

			// burn.css variables
			"color-primary":        "var(--color-primary)",
			"color-second-primary": "var(--color-second-primary)",
			"color-lightGrey":      "var(--color-lightGrey)",
			"color-grey":           "var(--color-grey)",
			"color-darkGrey":       "var(--color-darkGrey)",
			"color-error":          "var(--color-error)",
			"color-success":        "var(--color-success)",
			"main-bg":              "var(--main-bg)",
			"grid-maxWidth":        "var(--grid-maxWidth)",
			"grid-gutter":          "var(--grid-gutter)",
			"font-size":            "var(--font-size)",
			"-main-color":          "var(--main-color)",
		},
	}
	acssGen.Init()
	htmlCompiler.AddOnElementInfo(acssGen.OnElementInfo())

	builder := &gasx.Builder{
		BlockCompilers: []gasx.BlockCompiler{
			htmlCompiler.Block(),
		},
	}

	compileFiles(builder)
	gasx.ClearDir("dist")
	gasx.CopyDir("app/static", "dist")
	chooseIndexHTML(isDev)
	compileCode(builder, platform, isDev)
	compileStyles(builder, acssGen, isDev)
	gasx.Log("Building finished")
}

func compileFiles(builder *gasx.Builder) {
	gasx.Log("Compiling *.gox, *.gos")

	gasx.RunCommand("chmod 777 -R $GOPATH/pkg/mod;")

	files, err := gasx.GasFiles([]string{"gos"}) // get files
	gasx.Must(err)

	gasx.Must(builder.ParseFiles(files)) // compile them
}

func chooseIndexHTML(isDev bool) {
	var choose string
	if isDev {
		choose = "index.dev.html"
	} else {
		choose = "index.prod.html"
	}

	gasx.CopyFile("dist/"+choose, "dist/index.html")

	if !isDev {
		gasx.DeleteFile("dist/index.dev.html")
		gasx.DeleteFile("dist/index.prod.html")
	}
}

func compileStyles(builder *gasx.Builder, acssGen acss.Generator, isDev bool) {
	gasx.Log("Compiling css")

	depsPaths, err := gasx.GrepStyles()
	depsFile, err := gasx.UniteFilesByPaths(depsPaths)
	gasx.Must(err)

	gasx.NewFile("dist/deps.css", depsFile)
	gasx.NewFile("dist/acss.css", acssGen.GetStyles()) // save css styles generated by acss

	if isDev {
		gasx.RunCommand("postcss app/styles/main.pcss -o dist/main.css")
		gasx.RunCommand("postcss dist/acss.css -r")
		gasx.RunCommand("postcss dist/deps.css -r")
	} else {
		gasx.RunCommand("postcss --env production app/styles/main.pcss -o dist/main.m.css")
		gasx.RunCommand("postcss --env production dist/acss.css -o dist/acss.m.css")
		gasx.RunCommand("postcss --env production dist/deps.css -o dist/deps.m.css")
	}
}

func compileCode(builder *gasx.Builder, platform string, isDev bool) {
	gasx.Log("Compiling code")
	switch platform {
	case "gopherjs":
		gasx.RunCommand(fmt.Sprintf("cd app; gopherjs build -m -o ../dist/index.js; cd .."))
		if !isDev {
			gasx.RunCommand("uglifyjs dist/index.js -o dist/index.m.js")
		}
	case "wasm":
		gasx.RunCommand(fmt.Sprintf("cd app; GOOS=js GOARCH=wasm go build -o ../dist/main.wasm; cd .."))

		indexJSFile := "index.js"
		if !isDev {
			indexJSFile = "index.m.js"
		}

		gasx.NewFile("dist/"+indexJSFile, gasx.GetWASMExecScript())
	case "tinygo":
		gasx.ErrorMsg("comming soon(?)")
	default:
		gasx.ErrorMsg("invalid target platform")
	}
}

func serve(currentDir string) {
	gasx.Log("Starting static server")
	currentSrv := &http.Server{Addr: ":8080", Handler: http.FileServer(http.Dir(currentDir + "/dist"))}
	currentSrv.ListenAndServe()
}
