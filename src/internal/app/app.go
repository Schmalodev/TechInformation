package app

import (
	"TechInformation/src/internal/monitor"
	"fmt"
	"log"
	"time"

	"github.com/getlantern/systray"
)

func OnReady() {
	mQuit := systray.AddMenuItem("Beenden", "App Beenden")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	for {
		ramInformation := monitor.GetRamInformation()
		systray.SetTitle(fmt.Sprint("Total: ", ramInformation.Total, " Used: ", ramInformation.Used))
		time.Sleep(2 * time.Second)
	}
}

func OnExit() {
	log.Println("Exiting...")
}
