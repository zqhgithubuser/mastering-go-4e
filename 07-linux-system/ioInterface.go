package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1   S1
	text []byte
}

func (s *S1) Read(p []byte) (n int, err error) {
	fmt.Print("Give me your name: ")
	fmt.Scanln(&p)
	s.F2 = string(p)
	return len(p), nil
}

func (s *S1) Write(p []byte) (n int, err error) {
	if s.F1 < 0 {
		return -1, nil
	}
	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s ", p)
	}
	fmt.Println()
	return s.F1, nil
}

func (s *S2) eof() bool {
	return len(s.text) == 0
}

// 逐字节读取
func (s *S2) readByte() byte {
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return 0, err
	}
	l := len(p)
	if l > 0 {
		for n < l {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return n, nil
}

func main() {
	s1var := S1{4, "Hello"}
	fmt.Println(s1var) // {4 Hello}
	buf := make([]byte, 2)
	_, err := s1var.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read:", s1var.F2)
	_, _ = s1var.Write([]byte("Hello There!"))

	s2var := S2{
		F1:   s1var,
		text: []byte("Hello world!!"),
	}
	r := bufio.NewReader(&s2var)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", string(buf[:n]))
	}
}
