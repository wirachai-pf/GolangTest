package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	api "github.com/wirachai.pf/GolangTest/api"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf("%s %s", "E", "ee"))
	a.Version("ee")
	a.HelpFlag.Short('h')

	// Start
	startCmd := a.Command("start", "Start server command.")

	startTime := time.Now()
	fmt.Println(startTime)
	switch kingpin.MustParse(a.Parse(os.Args[1:])) {
	case startCmd.FullCommand():
		e := echo.New()
		api.InitWebApiEndpoint(e)
		e.Logger.Fatal(e.Start(":1323"))
	}
}
