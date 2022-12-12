package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.New()

	app.Get("/demo/{name}", demo)

	return app
}

func demo(ctx iris.Context) {

	auth := ctx.GetHeader("Authorization")

	b64 := strings.Split(auth, " ")[1]
	clear, err := base64.RawStdEncoding.DecodeString(b64)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
	split := strings.Split(string(clear), ":")

	fmt.Printf("username %s password %s\n", split[0], split[1])

	name := ctx.Params().Get("name")
	ctx.WriteString("Hello world " + name)

}

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			println()
			println("end program " + sig.String())
			os.Exit(0)
		}
	}()

	app := newApp()
	app.Listen(":8080")

}
