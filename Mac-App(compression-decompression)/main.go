package main

import (
    "github.com/andlabs/ui"
    "fmt"
)

func main() {
	err := ui.Main(func() {
        Show("main_window")
    })
    if err != nil {
        fmt.Println(err)
    }
    for{}
	
}
