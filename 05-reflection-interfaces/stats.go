package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
)

type DataFile struct {
	Filename string
	Len      int
	Minimum  float64
	Maximum  float64
	Mean     float64
	StdDev   float64
}

type DFSlice []DataFile

// ------------------------------------------------------------
// DFSlice 实现 sort.interface 接口

func (a DFSlice) Len() int {
	return len(a)
}

func (a DFSlice) Less(i, j int) bool {
	return a[i].Mean < a[j].Mean
}

func (a DFSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ------------------------------------------------------------

func readCSVFile(filepath string) ([]float64, error) {
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
			fmt.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}
	return values, nil
}

func stdDev(x []float64) (float64, float64) {
	sum := 0.0
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
		fmt.Println("Need one or more file paths!")
		return
	}
	files := DFSlice{}

	for i := 1; i < len(os.Args); i++ {
		file := os.Args[i]
		currentFile := DataFile{}
		currentFile.Filename = file

		values, err := readCSVFile(file)
		if err != nil {
			fmt.Println("Error reading:", file, err)
			os.Exit(0)
		}

		currentFile.Len = len(values)
		currentFile.Minimum = slices.Min(values)
		currentFile.Maximum = slices.Max(values)
		meanValue, standardDeviation := stdDev(values)
		currentFile.Mean = meanValue
		currentFile.StdDev = standardDeviation

		files = append(files, currentFile)
	}

	// 排序
	sort.Sort(files)
	for _, val := range files {
		f := val.Filename
		fmt.Println(f, ":", val.Len, val.Mean, val.Maximum, val.Minimum)
	}
}
