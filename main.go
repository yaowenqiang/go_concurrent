package main

import (
    "fmt"
    "time"
    "runtime"
)

func main() {
    godur, _ := time.ParseDuration("10ms")
    runtime.GOMAXPROCS(2)

    go func() {
        for i := 0; i <= 100; i++ {
            fmt.Println("hello")
            time.Sleep(godur)
        }
    }()
    go func() {
        for i := 0; i <= 100; i++ {
            fmt.Println("go")
            time.Sleep(godur)
        }
    }()

    dur, _ := time.ParseDuration("1s")
    time.Sleep(dur)
}
