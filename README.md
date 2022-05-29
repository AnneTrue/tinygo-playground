# Introduction

This is my playground for using TinyGo with my Arduino Uno v3.

# Flashing

For instance to flash the `blinky` module:

`tinygo flash -scheduler tasks -target arduino -port /dev/ttyACM0 /home/ant/Documents/tinygo-playground/blinky/main.go`
