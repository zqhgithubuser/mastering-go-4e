package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	re := regexp.MustCompile("\\S+")
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			// 最后一行非空
			if len(line) != 0 {
				words := re.FindAllString(line, -1)
				for i := 0; i < len(words); i++ {
					fmt.Println(words[i])
				}
			}
			break
		}

		if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		words := re.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
