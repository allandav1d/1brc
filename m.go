package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func main() {
	start := time.Now()
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemperature := rawData[semicolon+1:]

		temperature, _ := strconv.ParseFloat(rawTemperature, 64)

		measurement, ok := dados[location]
		if !ok {
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Max, temperature)
			measurement.Sum += temperature
			measurement.Count++
		}

		dados[location] = measurement
	}

	locations := make([]string, 0, len(dados))
	for location := range dados {
		locations = append(locations, location)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, location := range locations {
		measurement := dados[location]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",
			location,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")

	fmt.Printf("Tempo de execução: %v\n", time.Since(start))
}
