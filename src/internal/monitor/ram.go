package monitor

import "github.com/shirou/gopsutil/v3/mem"

type RamInfo struct {
	Total          uint64
	Used           uint64
	Free           uint64
	UsedPreferenct float64
}

func GetRamInfo() (*RamInfo, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &RamInfo{
		Total:          v.Total,
		Used:           v.Used,
		Free:           v.Free,
		UsedPreferenct: v.UsedPercent,
	}, nil
}
