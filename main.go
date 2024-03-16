package main

import (
    "fmt"
    "os"
    "cid-local/internal/app/cid-local"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: cid-local <filename>")
        return
    }
    
    filename := os.Args[1]

    fileData, err := os.ReadFile(filename)
    if err != nil {panic(err)}

    cid := cid_local.Cid(fileData)
    fmt.Println(cid)
}

