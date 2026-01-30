package main

import (
	"TechInformation/src/internal/monitor"
	"fmt"
	"time"
)

func main() {
	for {
		info, err := monitor.GetRamInfo()
		if err != nil {
			fmt.Println("Fehler beim Auslesen: ", err)
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf(
			"Ram benutzt: %.2f%% (Used: %d MB / Total: %d MB)\n",
			info.UsedPreferenct,
			info.Used/1024/1024,
			info.Total/1024/1024,
		)

		time.Sleep(1 * time.Second)
	}
}
