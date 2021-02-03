package main

import (
    "fmt"
    "strings"
)

func main() {
    phrase := "These are the times that try man's souls\n"
    words := strings.Split(phrase, " ")
    ch := make(chan string, len(phrase))
    //ch <- "hello"
    for _, word := range words {
        //fmt.Println(<-ch)
        ch <- word
    }

    // Just close send ing of the channel
    close(ch)

    /*
    for i := 0; i < len(words); i++ {
        fmt.Println(<-ch + " ")
    }
    */
    /*
    for  {
        if msg , ok := <- ch; ok {
            fmt.Println(msg + " ")
        } else {
            break
        }
    }
    */

    for msg := range ch {
        fmt.Println(msg + " ")
    }
    //ch <- "test"



}
