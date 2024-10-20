package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
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

	// Mapa final de dados
	dados := make(map[string]Measurement)
	var wg sync.WaitGroup

	lines := make(chan string, 100000) // Buffer aumentado para 100 mil linhas
	results := make(chan map[string]Measurement, 100)

	// Worker Goroutines (12 threads, igual ao número de threads do processador)
	numWorkers := 12
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			localData := make(map[string]Measurement)
			batchSize := 1000 // Processa 1000 linhas por vez em cada Goroutine
			batch := make([]string, 0, batchSize)

			for rawData := range lines {
				batch = append(batch, rawData)

				// Processa o lote quando atingir o tamanho especificado
				if len(batch) >= batchSize {
					processBatch(batch, localData)
					batch = batch[:0] // Limpa o lote
				}
			}

			// Processa as linhas restantes no final
			if len(batch) > 0 {
				processBatch(batch, localData)
			}
			results <- localData
		}()
	}

	// Goroutine para ler o arquivo
	go func() {
		scanner := bufio.NewScanner(measurements)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)

		// Verificação de erro ao ler o arquivo
		if err := scanner.Err(); err != nil {
			fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		}
	}()

	// Goroutine para fechar o canal de resultados após a conclusão dos workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Coleta dos resultados dos workers e merge final
	for localData := range results {
		for location, measurement := range localData {
			finalMeasurement, ok := dados[location]
			if !ok {
				dados[location] = measurement
			} else {
				finalMeasurement.Min = min(finalMeasurement.Min, measurement.Min)
				finalMeasurement.Max = max(finalMeasurement.Max, measurement.Max)
				finalMeasurement.Sum += measurement.Sum
				finalMeasurement.Count += measurement.Count
				dados[location] = finalMeasurement
			}
		}
	}

	// Ordena as localizações e imprime os resultados
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

func processBatch(batch []string, localData map[string]Measurement) {
	for _, rawData := range batch {
		semicolon := strings.Index(rawData, ";")
		if semicolon == -1 {
			// Ignora linha inválida
			continue
		}
		location := rawData[:semicolon]
		rawTemperature := rawData[semicolon+1:]

		// Suprimindo verificação de erro para desempenho
		temperature, _ := strconv.ParseFloat(rawTemperature, 64)

		measurement, ok := localData[location]
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
		localData[location] = measurement
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
