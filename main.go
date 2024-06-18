package main

import (
	"flag"
	"micro-api/routers"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload"
)

var (
	title   = ""
	version = ""
)

func init() {
	flag.StringVar(&title, "t", "title", "set title for API")
	flag.StringVar(&version, "v", "v0.0.1", "set version")
	flag.Parse()

	switch runtime.GOOS {
	case "darwin", "linux":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		log.Info().Msg("Running on linux")

	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("Running on windows")
	}
}

// @title Micro API Server
// @version 0.0.1
// @description This is a sample server and postman server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	routers.New(e)

	e.Logger.Fatal(e.Start(":8080"))
}
