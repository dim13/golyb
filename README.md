# golyb
Go love your brain - yet another optimizing BrainF\*ck interpreter in Go

## Installation
    go get -u github.com/dim13/golyb

## Usage
```
Usage of golyb:
  -debug
        Enable debugging
  -dump
        Dump AST
  -file string
        Source file (required)
  -in string
        Input file
  -noopt
        Disable optimization
  -out string
        Output file or /dev/null
  -profile string
        Write CPU profile to file
  -show int
        Dump # tape cells around last position
  -tape string
        Tape type: static or dynamic (default "static")
```

## Tape storage type
- static: 32k cells including 1k margin on the lower end (used by default)
- dynamic: allocates in 1k chunks as required on access

# Code optimization
- [x] Contraction
- [x] Clear loops
- [x] Copy loops
- [x] Multiplication loops
- [x] Scan loops (kind of)
- [x] Operation offsets

Reference: http://calmerthanyouare.org/2015/01/07/optimizing-brainfuck.html

## Some rough results

| Program     | w/o optimization | with optimization | speed gain |
| -----------:| ----------------:| -----------------:| ----------:|
| madelbrot.b |   1 min 25.1 sec |          15.4 sec |       5.5x |
| long.b      |   1 min  9.5 sec |           7.6 sec |       9.0x |
| hanoi.b     |         58.3 sec |           1.3 sec |      44.8x |

## CPU profiles

### mandelbrot.b
![mandelbrot profile](https://raw.githubusercontent.com/dim13/golyb/master/profiles/mandelbrot.gif)

### long.b
![long profile](https://raw.githubusercontent.com/dim13/golyb/master/profiles/long.gif)

### hanoi.b
![hanoi profile](https://raw.githubusercontent.com/dim13/golyb/master/profiles/hanoi.gif)
