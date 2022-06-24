// Control the pulsing of an LED using a potentiometer and analog input on an Uno
package main

import (
    "machine"
    "time"
)

func main() {
    c := make(chan int64, 1)
    go getInputForever(c)
    blinkForever(c)
}

func blinkForever(c chan int64) () {
    led := machine.D11
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    var normal int64 = 0
    var rawADCValue int64 = 0
    for {
        rawADCValue = <- c
        normal = 2500 * rawADCValue / 0xffff // range from 0 to 5 sec period
        if normal > 2500 {
            normal = 2500
        } else if normal < 50 {
            normal = 50
        }
        led.High()
        delayMilliseconds(normal)
        led.Low()
        delayMilliseconds(normal)
    }
}

func getInputForever(c chan int64) () {
    machine.InitADC() // init the machine's ADC subsystem
    input0 := machine.ADC{machine.ADC0}
    var rawADCValue uint16 = 0
    for {
        rawADCValue = input0.Get()
        c <- int64(rawADCValue)
    }
}

func delayMilliseconds(t int64) {
    time.Sleep(time.Duration(1000000 * t))
}
