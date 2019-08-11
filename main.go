package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/kouhin/envflag"

	"github.com/gascore/gasx"
	"github.com/gascore/gasx/acss"
	"github.com/gascore/gasx/html"
)

var (
	currentDir string
	currentSrv *http.Server
)

func main() {
	var (
		err            error
		lockFileName   = flag.String("lockfile", ".gaslock", "Lock file name")
		platform       = flag.String("platform", "wasm", "Platform")
		ignoreExternal = flag.Bool("ignoreExternal", true, "Ignore external *.gox files?")
		watch          = flag.Bool("watch", false, "Watch for file changings")
	)
	if err := envflag.Parse(); err != nil {
		panic(err)
	}

	currentDir, err = os.Getwd()
	gasx.Must(err)

	build(*lockFileName, *platform, *ignoreExternal)

	if *watch {
		serve()
		gasx.Must(gasx.StartWatcher(
			func(path string) {
				fmt.Println(color.GreenString("Triggered:"), strings.TrimPrefix(path, currentDir+"/app/"))
				build(*lockFileName, *platform, *ignoreExternal)

				currentSrv.Shutdown(context.TODO())

				serve()
			},
			[]string{"_gas.go"},
			[]string{},
			[]string{"app/"},
		))
	}
}

func build(lockFileName string, platform string, ignoreExternal bool) {
	lockFile, err := gasx.GetLockFile(lockFileName, ignoreExternal)
	gasx.Must(err)

	htmlCompiler := html.NewCompiler()

	acssGen := acss.Generator{
		LockFile: lockFile,
		BreakPoints: map[string]string{
			"lg": "@media(min-width:1200px)",
			"md": "@media(min-width:1000px)",
			"sm": "@media(min-width:750px)",
		},
		Custom: map[string]string{
			"b": "1px solid #dedede",

			// burn.css variables
			"color-primary":         "var(--color-primary)",
			"color-second-primary": "var(--color-second-primary)",
			"color-lightGrey":       "var(--color-lightGrey)",
			"color-grey":            "var(--color-grey)",
			"color-darkGrey":        "var(--color-darkGrey)",
			"color-error":           "var(--color-error)",
			"color-success":         "var(--color-success)",
			"main-bg":               "var(--main-bg)",
			"grid-maxWidth":         "var(--grid-maxWidth)",
			"grid-gutter":           "var(--grid-gutter)",
			"font-size":             "var(--font-size)",
			"-main-color":           "var(--main-color)",
		},
	}
	htmlCompiler.AddOnAttribute(acssGen.OnAttribute())

	builder := &gasx.Builder{
		LockFile: lockFile,
		BlockCompilers: []gasx.BlockCompiler{
			htmlCompiler.Block(),
		},
	}

	compileFiles(builder, lockFile.BuildExternal)

	gasx.ClearDir("dist")
	gasx.CopyDir("app/static", "dist")
	compileCode(builder, platform)
	gasx.NewFile("dist/acss.css", acssGen.GetStyles()) // save css styles generated by acss
	gasx.RunCommand("sass app/styles/main.scss dist/main.css")

	lockFile.Save()
	gasx.Log("Builded successfully!")
}

func compileFiles(builder *gasx.Builder, buildExternal bool) {
	gasx.Log("Compiling *.gox, *.gos")

	gasx.RunCommand("chmod 777 -R $GOPATH/pkg/mod;")

	files, err := gasx.GasFiles([]string{"gos"}, buildExternal) // get files
	gasx.Must(err)

	gasx.Must(builder.ParseFiles(files)) // compile them
}

func compileCode(builder *gasx.Builder, platform string) {
	gasx.Log("Compiling code")
	switch platform {
	case "gopherjs":
		gasx.RunCommand(fmt.Sprintf("cd app; gopherjs build -m -o ../dist/index.js; cd .."))
	case "wasm":
		gasx.RunCommand(fmt.Sprintf("cd app; GOOS=js GOARCH=wasm go build -o ../dist/main.wasm; cd .."))
		gasx.NewFile("dist/index.js", gasx.GetWASMExecScript())
	case "tinygo":
		gasx.ErrorMsg("comming soon(?)")
	default:
		gasx.ErrorMsg("invalid target platform")
	}
}

func serve() {
	currentSrv = &http.Server{Addr: ":8080", Handler: http.FileServer(http.Dir(currentDir + "/dist"))}
	go func() {
		currentSrv.ListenAndServe()
	}()
}
