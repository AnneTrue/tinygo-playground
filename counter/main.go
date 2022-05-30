// Counts an int using goroutines and channels for a 7-segment display
package main

import (
    "github.com/AnneTrue/tinygo-playground/display"
    "time"
)

func main() {
    c := make(chan int64)
    display.InitDisplay()
    go countToChannel(c)
    displayFromChannel(c)
}

func displayFromChannel(c chan int64) () {
    for {
        display.DisplayInt(<- c)
    }
}

func countToChannel(c chan int64) () {
    var currentInt int64 = 0
    for {
        c <- currentInt
        delay(200)
        currentInt = currentInt + 1
    }
}

func delay(t int64) {
    time.Sleep(time.Duration(1000000 * t))
}
