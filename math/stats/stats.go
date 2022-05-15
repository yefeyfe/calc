package stats

import "fmt"

func init() {
	fmt.Println("initialized package stats (local).")
}

func Avg(vals []float64) float64 {
	total := 0.0
	for _, val := range vals {
		total += val
	}
	return total / float64(len(vals))
}