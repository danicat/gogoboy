# gogoboy

A GameBoy emulator written in Go

## About this project

This project is a proof of concept of building emulators with test driven development. I've chosen the GameBoy because from what I read it's architecture is quite simple, but I might be proven wrong. :)

## Current Status

- CPU: 31 opcodes implemented (out of 256)
- CPU: 8 bit registers implemented: A, C, D, E, H, L
- CPU: Flags implemented
- MRAM: can load a program
- MRAM: can read from address

## TODO

- Run Nintendo boot program