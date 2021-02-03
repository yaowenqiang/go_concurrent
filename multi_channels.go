package main

import (
    "fmt"
)

type Message struct {
    To []string
    From string
    Content string
}

type FailedMessage struct {
    ErrorMessage string
    OriginalMessage Message
}

func main() {
    msgCh := make(chan Message, 1)
    errCh := make(chan FailedMessage, 1)

    msg := Message{
        To: []string{"hello@jack.com"},
        From: "jack@jack.jack",
        Content: "he ther",
    }

    //msgCh <- msg

    failedMessage := FailedMessage{
        ErrorMessage:" Message interceptd by black rider",
        OriginalMessage: Message{},
    }

    errCh <- failedMessage
    msgCh <- msg

    /*
    fmt.Println(<-msgCh)
    fmt.Println(<-errCh)
    */

    select {
    case receivedMsg := <- msgCh:
        fmt.Println(receivedMsg)
    case receivedError := <- errCh:
        fmt.Println(receivedError)
    default:
        fmt.Println("No message received")
    }
}

