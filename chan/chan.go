package main

import (
    "fmt"
    "time"
)

func main() {
    messages := make(chan string)

    go boring(messages)

    msg := <-messages
    fmt.Println(msg)

    time.Sleep(100)
}

func boring(messages chan string) {
    messages <- "yep"
}

