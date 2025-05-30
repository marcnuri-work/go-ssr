package web

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"rogchap.com/v8go"
)

func StartServer() {
	engine := gin.Default()
	engine.GET("/", main)
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}

func main(c *gin.Context) {
	js := esBuild()
	// Render using V8 (c++ bindings)
	dom, err := v8Eval(js)
	// Render using Goja (Go native JS)
	dom, err = evalJavaScript(js)
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Go SSR</title>
		</head>
		<body>
			<div id="root">%s</div>
			<script type="module">
				window.goSsr = {render: () => {
				}};
			</script>
			<script type="module">
				%s
			</script>
		</body>
		</html>
	`, dom, js)
	if err != nil {
		fmt.Printf("Error evaluating JavaScript: %v\n", err)
		c.String(500, "Internal Server Error")
		return
	} else {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html)
	}
}

func esBuild() string {
	res := api.Build(api.BuildOptions{
		EntryPoints: []string{path.Join("assets", "entry-point.jsx")},
		Bundle:      true,
		Outdir:      "out-ignored",
		Write:       false,
		Platform:    api.PlatformBrowser,
		Define: map[string]string{
			"process.env.NODE_ENV": `"production"`,
		},
		Inject: []string{
			path.Join("assets", "polyfills.js"),
		},
	})
	return string(res.OutputFiles[0].Contents)
}

func v8Eval(js string) (string, error) {
	isolate := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(isolate)
	goSsr := v8go.NewObjectTemplate(isolate)
	global.Set("goSsr", goSsr)
	renderedHtml := ""
	render := v8go.NewFunctionTemplate(isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		if len(info.Args()) > 0 {
			renderedHtml = info.Args()[0].String()
		}
		return nil
	})
	goSsr.Set("render", render)
	context := v8go.NewContext(isolate, global)
	f, err := os.CreateTemp("", "script.js")
	fmt.Printf("Failed to create temporary file: %v", err)
	defer os.Remove(f.Name())
	_, err = context.RunScript(js, path.Dir(f.Name()))
	return renderedHtml, err
}

func evalJavaScript(js string) (string, error) {
	vm := goja.New()
	renderedHtml := ""
	err := vm.Set("goSsr", map[string]interface{}{
		"render": func(renderedString string) {
			renderedHtml = renderedString
		},
	})
	if err != nil {
		return "", err
	}
	_, err = vm.RunString(js)
	return renderedHtml, err
}
