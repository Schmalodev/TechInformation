package monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

type RamInformations struct {
	Total int
	Used  int
}

func GetRamInformation() RamInformations {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("its not Reading")
	}

	ram := RamInformations{
		Total: int(v.Total),
		Used:  int(v.Used),
	}

	return ram
}
