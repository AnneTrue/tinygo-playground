# Introduction

This is my playground for using TinyGo with my Arduino Uno v3.

# Flashing

For instance to flash the `blinky` module:

`tinygo flash -scheduler tasks -target arduino -port /dev/ttyACM0 /home/ant/Documents/tinygo-playground/blinky/main.go`

UART output/input can be interacted with using screen, with the port and baud rate:

`screen /dev/ttyACM0 9600`

Escaping screen can be done with ctrl+a to enter command mode, then typing `:quit`

UART can also be read using socat:

`socat stdio /dev/ttyACM0`
