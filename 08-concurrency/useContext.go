package main

import (
    "context"
    "fmt"
    "os"
    "strconv"
    "time"
)

func func1(t int) {
    c1 := context.Background()
    // 创建一个可取消的子上下文 c1，并返回一个取消函数 cancel
    c1, cancel := context.WithCancel(c1)
    defer cancel()

    select {
    case <-c1.Done():
        fmt.Println("f1() Done:", c1.Err())
        return
    case r := <-time.After(time.Duration(t) * time.Second):
        fmt.Println("f1():", r)
    }
    return
}

func func2(t int) {
    c2 := context.Background()
    // 创建一个带有超时的上下文 c2，超时设置为 t 秒
    c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
    defer cancel()

    select {
    case <-c2.Done():
        fmt.Println("f2() Done:", c2.Err())
        return
    case r := <-time.After(time.Duration(2*t) * time.Second):
        fmt.Println("f2():", r)
    }
    return
}

func func3(t int) {
    c3 := context.Background()
    // 创建一个带有固定截止时间（deadline）的上下文 c3
    deadline := time.Now().Add(time.Duration(t) * time.Second)
    c3, cancel := context.WithDeadline(c3, deadline)
    defer cancel()

    select {
    case <-c3.Done():
        fmt.Println("f3() Done:", c3.Err())
        return
    case r := <-time.After(time.Duration(2*t) * time.Second):
        fmt.Println("f3():", r)
    }
    return
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Need a delay!")
        return
    }

    delay, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Delay:", delay)
    func1(delay)
    func2(delay)
    func3(delay)
}
