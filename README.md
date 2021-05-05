# gogoboy

A GameBoy emulator written in Go

## About this project

This project is a proof of concept of building emulators with test driven development. I've chosen the GameBoy because from what I read it's architecture is quite simple, but I might be proven wrong. :)

## Current Status

- CPU: 40 opcodes implemented (out of 256)
- CPU: PC implemented
- CPU: 8 bit registers implemented: A, B, C, D, E, H, L
- CPU(flags): Z, N, H and C implemented
- CPU(stack): - SP, PUSH and POP implemented
- MRAM: can load a program
- MRAM: can read from address
- MRAM: can write to an address

## TODO

- Run Nintendo boot program