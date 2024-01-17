package datahuman

import (
	"math"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
)

func UptimeToDays() {

}

// func FmtFloat()

func CpuPercent(prev cpu.TimesStat, curr cpu.TimesStat) float64 {

	return calculateBusy(prev, curr)
}

func getAllBusy(t cpu.TimesStat) (float64, float64) {
	tot := t.Total()
	if runtime.GOOS == "linux" {
		tot -= t.Guest     // Linux 2.6.24+
		tot -= t.GuestNice // Linux 3.2.0+
	}

	busy := tot - t.Idle - t.Iowait

	return tot, busy
}

func getTimeStatTotal(c *cpu.TimesStat) float64 {
	total := c.User + c.System + c.Idle + c.Nice + c.Iowait + c.Irq +
		c.Softirq + c.Steal + c.Guest + c.GuestNice

	return total
}

func calculateBusy(t1, t2 cpu.TimesStat) float64 {
	t1All, t1Busy := getAllBusy(t1)
	t2All, t2Busy := getAllBusy(t2)

	if t2Busy <= t1Busy {
		return 0
	}
	if t2All <= t1All {
		return 100
	}
	return math.Min(100, math.Max(0, (t2Busy-t1Busy)/(t2All-t1All)*100))
}

func calculateAllBusy(t1, t2 []cpu.TimesStat) []float64 {

	if len(t1) != len(t2) {
		return nil
	}

	ret := make([]float64, len(t1))
	for i, t := range t2 {
		ret[i] = calculateBusy(t1[i], t)
	}
	return ret
}
