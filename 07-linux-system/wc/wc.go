package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"os"
	"regexp"
)

func countByLine(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	lineCount := 0
	for s.Scan() {
		lineCount++
	}
	return lineCount, nil
}

func countByWord(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	re := regexp.MustCompile("\\S+")
	wordCount := 0
	for s.Scan() {
		words := re.FindAllString(s.Text(), -1)
		wordCount += len(words)
	}
	return wordCount, nil
}

func countByChar(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	charCount := 0
	for {
		_, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error reading file %s", err)
			return 0, err
		}
		charCount++
	}
	return charCount, nil
}

func process(file string, options map[string]bool) {
	output := ""
	if options["lines"] {
		lines, err := countByLine(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		output += fmt.Sprintf("%d ", lines)
	}

	if options["words"] {
		words, err := countByWord(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		output += fmt.Sprintf("%d ", words)
	}

	if options["chars"] {
		chars, err := countByChar(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		output += fmt.Sprintf("%d ", chars)
	}
	fmt.Printf("%s %s", output, file)
}

func main() {
	pflag.Usage = func() {
		fmt.Println("Usage: wc [OPTION]... [FILE]...")
		pflag.PrintDefaults()
	}

	pflag.BoolP("lines", "l", false, "print the newline counts")
	pflag.BoolP("words", "w", false, "print the word counts")
	pflag.BoolP("chars", "m", false, "print the character counts")

	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	if !viper.IsSet("lines") && !viper.IsSet("words") && !viper.IsSet("chars") {
		viper.Set("lines", true)
		viper.Set("words", true)
		viper.Set("chars", true)
	}

	args := pflag.Args()
	if len(args) == 0 {
		fmt.Println("Error: no file specified")
		os.Exit(1)
	}

	options := map[string]bool{
		"lines": viper.GetBool("lines"),
		"words": viper.GetBool("words"),
		"chars": viper.GetBool("chars"),
	}

	for _, file := range args {
		process(file, options)
	}
}
