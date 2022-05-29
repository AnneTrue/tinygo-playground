package main

import (
    "machine"
    "time"
)

func main() {
    go led1()
    go led2()
    led3()
}

func led1() {
    led := machine.D2
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.High()
        delay(500)
        led.Low()
        delay(500)
    }
}

func led2() {
    led := machine.D3
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.High()
        delay(400)
        led.Low()
        delay(400)
    }
}

func led3() {
    led := machine.D4
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.High()
        delay(600)
        led.Low()
        delay(600)
    }
}

func delay(t int64) {
    time.Sleep(time.Duration(1000000 * t))
}
