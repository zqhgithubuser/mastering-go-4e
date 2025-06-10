package main

import "fmt"

type IntA interface {
    foo()
}

type IntB interface {
    bar()
}

// IntC 需要实现 foo() 和 bar() 两个方法
type IntC interface {
    IntA
    IntB
}

type a struct {
    XX int
    YY int
}

type b struct {
    AA string
    XX int
}

type c struct {
    A a
    B b
}

func processA(s IntA) {
    fmt.Printf("%T\n", s)
}

// 满足 IntA 接口
func (varC c) foo() {
    fmt.Println("Foo Processing", varC)
}

// 满足 IntB 接口
func (varC c) bar() {
    fmt.Println("Bar Processing", varC)
}

type compose struct {
    field1 int
    a      // 匿名字段
}

func (A a) A() {
    fmt.Println("Function A() for A")
}

func (B b) A() {
    fmt.Println("Function A() for B")
}

func main() {
    var iC = c{
        a{120, 12},
        b{"-12", -12},
    }

    iC.A.A()     // Function A() for A
    iC.B.A()     // Function A() for B
    iC.bar()     // Bar Processing {{120 12} {-12 -12}}
    processA(iC) // main.c

    iComp := compose{
        123,
        a{456, 789},
    }
    // 直接访问匿名结构体 a 的字段
    fmt.Println(iComp.XX, iComp.YY, iComp.field1) // 456 789 123
}
