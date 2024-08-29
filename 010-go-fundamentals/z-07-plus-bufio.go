package main

import (
    "bufio"
    "fmt"
    "os"
)

func buffered() {
    file, err := os.Open("log.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    reader := bufio.NewReader(file)
    data := make([]byte, 1000)
    _, err = reader.Read(data)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(data))
}