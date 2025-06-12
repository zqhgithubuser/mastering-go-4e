/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

type Entry struct {
	Filename string  `json:"filename"`
	Len      int     `json:"length"`
	Minimum  float64 `json:"minimum"`
	Maximum  float64 `json:"maximum"`
	Mean     float64 `json:"mean"`
	StdDev   float64 `json:"stddev"`
}

var logger *slog.Logger
var JSONFILE = "./data.json"

type DFSlice []Entry

var data = DFSlice{}
var index map[string]int
var disableLogging bool

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = Serialize(&data, f)
	return err
}

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// json -> Go Structure
	err = DeSerialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Filename
		index[key] = i
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stats",
	Short: "Statistics application",
	Long:  `The statistics application.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	if disableLogging == false {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	slog.SetDefault(logger)
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&disableLogging, "log", "l", false, "Disable logging output")

	err := readJSONFile(JSONFILE)
	if err != nil && err != io.EOF {
		return
	}
	createIndex()
}
