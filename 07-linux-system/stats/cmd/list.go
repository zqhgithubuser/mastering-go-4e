/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "io"
    "log/slog"
    "os"
    "sort"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "A brief description of your command",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {
        list()
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}

func list() {
    sort.Sort(data)
    text, err := PrettyPrintJSONStream(data)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(text)

    logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
    if disableLogging == false {
        logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
    }

    slog.SetDefault(logger)
    s := fmt.Sprintf("%d records in total.", len(data))
    logger.Info(s)
}

func PrettyPrintJSONStream(data interface{}) (string, error) {
    buffer := new(bytes.Buffer)
    encoder := json.NewEncoder(buffer)
    encoder.SetIndent("", "\t")

    err := encoder.Encode(data)
    if err != nil {
        return "", err
    }
    return buffer.String(), nil
}

// -----------------------------------------------------------
// Implement sort.Interface

func (a DFSlice) Len() int {
    return len(a)
}

func (a DFSlice) Less(i, j int) bool {
    if a[i].Mean == a[j].Mean {
        return a[i].StdDev < a[j].StdDev
    }
    return a[i].Mean < a[j].Mean
}

func (a DFSlice) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

// -----------------------------------------------------------
