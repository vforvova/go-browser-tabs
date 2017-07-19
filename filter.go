package main

import "os"
import "fmt"
import "strings"

import "./browser-tabs.go"

func main() {
    query := strings.TrimSpace(os.Args[1])
    fmt.Println(query)
}
