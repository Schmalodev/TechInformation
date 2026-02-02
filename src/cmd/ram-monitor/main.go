package main

import (
	"TechInformation/src/internal/app"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(app.OnReady, app.OnExit)
}
