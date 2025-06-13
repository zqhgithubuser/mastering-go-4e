package main

import (
    "context"
    "fmt"
    "golang.org/x/sync/semaphore"
    "os"
    "strconv"
    "time"
)

var Workers = 4
var sem = semaphore.NewWeighted(int64(Workers))

func worker(n int) int {
    square := n * n
    time.Sleep(time.Second)
    return square
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Need #jobs!")
        return
    }

    nJobs, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    var results = make([]int, nJobs)

    ctx := context.TODO()

    for i := range results {
        err = sem.Acquire(ctx, 1)
        if err != nil {
            fmt.Println("Cannot acquire semaphore:", err)
            break
        }

        go func(i int) {
            defer sem.Release(1)
            temp := worker(i)
            results[i] = temp
        }(i)
    }

    // 等待所有 goroutines 完成任务
    err = sem.Acquire(ctx, int64(Workers))
    if err != nil {
        fmt.Println(err)
    }

    for k, v := range results {
        fmt.Println(k, "->", v)
    }
}
