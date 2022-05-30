// Uses pins 2 to 8 for a seven-segment display.
// D2: Top right
// D3: Top
// D4: Top left
// D5: Bottom right
// D6: Middle
// D7: Bottom
// D8: Bottom left

package display

import (
    "machine"
)

func InitDisplay() {
    pins := [...]machine.Pin{
        machine.D2,
        machine.D3,
        machine.D4,
        machine.D5,
        machine.D6,
        machine.D7,
        machine.D8,
    }
    for _, singlePin := range pins {
        singlePin.Configure(machine.PinConfig{Mode: machine.PinOutput})
    }
}

func DisplayInt(digit int64) () {
    singleDigit := digit % 10
    switch {
    case singleDigit == 0:
        displayArray([7]bool{true, true, true, true, false, true, true})
    case singleDigit == 1:
        displayArray([7]bool{true, false, false, true, false, false, false})
    case singleDigit == 2:
        displayArray([7]bool{true, true, false, false, true, true, true})
    case singleDigit == 3:
        displayArray([7]bool{true, true, false, true, true, true, false})
    case singleDigit == 4:
        displayArray([7]bool{true, false, true, true, true, false, false})
    case singleDigit == 5:
        displayArray([7]bool{false, true, true, true, true, true, false})
    case singleDigit == 6:
        displayArray([7]bool{false, true, true, true, true, true, true})
    case singleDigit == 7:
        displayArray([7]bool{true, true, false, true, false, false, false})
    case singleDigit == 8:
        displayArray([7]bool{true, true, true, true, true, true, true})
    case singleDigit == 9:
        displayArray([7]bool{true, true, true, true, true, true, false})
    }
}

func displayArray(segments [7]bool) () {
    pins := [...]machine.Pin{
        machine.D2,
        machine.D3,
        machine.D4,
        machine.D5,
        machine.D6,
        machine.D7,
        machine.D8,
    }
    for idx, singlePin := range pins {
        if segments[idx] {
            singlePin.High()
            continue
        }
        singlePin.Low()
    }
}
