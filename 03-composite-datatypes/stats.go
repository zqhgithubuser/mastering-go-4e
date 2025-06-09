package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func readFile(filepath string) ([]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	// [][]string -> []float64
	for _, line := range lines {
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}
	return values, nil
}

func stdDev(x []float64) (float64, float64) {
	sum := float64(0)
	for _, val := range x {
		sum += val
	}
	// 平均值
	meanValue := sum / float64(len(x))
	fmt.Printf("Mean value: %.5f\n", meanValue)
	// 标准差
	var squared float64
	for i := 0; i < len(x); i++ {
		squared += math.Pow(x[i]-meanValue, 2)
	}
	standardDeviation := math.Sqrt(squared / float64(len(x)))
	return meanValue, standardDeviation
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}
	return normalized
}

func main() {
	if len(os.Args) == 1 {
		log.Println("Need one argument!")
		return
	}
	file := os.Args[1]

	values, err := readFile(file)
	if err != nil {
		log.Println("Error reading:", file, err)
		os.Exit(0)
	}
	sort.Float64s(values)

	fmt.Println("Number of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	meanValue, standardDeviation := stdDev(values)
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	normalized := normalize(values, meanValue, standardDeviation)
	fmt.Println("Normalized:", normalized)
}
