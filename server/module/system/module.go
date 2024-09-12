package system

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type Info struct {
	Basic    Basic    `json:"basic"`
	Hardware Hardware `json:"hardware"`
}

type Basic struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"buildTime"`
	Hostname  string `json:"hostname"`
}

type Hardware struct {
	Cpu  []cpu.InfoStat        `json:"cpu"`
	Mem  mem.VirtualMemoryStat `json:"mem"`
	Host host.InfoStat         `json:"host"`
}

type HolidayAPI struct {
	Name      string                  `json:"Name"`
	Version   string                  `json:"Version"`
	Generated string                  `json:"Generated"`
	Timezone  string                  `json:"Timezone"`
	Author    string                  `json:"Author"`
	URL       string                  `json:"URL"`
	Years     map[string][]YearEntity `json:"Years"`
}

type YearEntity struct {
	Name      string        `json:"Name"`
	StartDate string        `json:"StartDate"`
	EndDate   string        `json:"EndDate"`
	Duration  int64         `json:"Duration"`
	CompDays  []interface{} `json:"CompDays"`
	URL       string        `json:"URL"`
	Memo      string        `json:"Memo"`
}
