# gofys
Go f\*ck your self -- yet another optimizing BrainF\*ck interpreter in Go

## Installation
    go get -u github.com/dim13/gofys

## Usage
```
Usage of gofys:
  -debug
    	Enable debugging
  -dump
    	Dump AST
  -file string
    	Source file (required)
  -noopt
    	Disable optimization
  -out string
    	Output file or /dev/null
  -profile string
    	Write CPU profile to file
  -tape string
    	Tape type: static or dynamic (default "static")
```

## Tape storage type
- static: 32k cells including 1k margin on the lower end (used by default)
- dynamic: allocates in 1k chunks as required on access

Reference: http://calmerthanyouare.org/2015/01/07/optimizing-brainfuck.html

# Code optimization
- [x] Contraction
- [x] Clear loops
- [x] Copy loops
- [x] Multiplication loops
- [x] Scan loops (kind of)
- [x] Operation offsets

## Some rough results

| Program     | with optimization | w/o optimization | speed gain |
| ----------- | -----------------:| ----------------:| ----------:|
| madelbrot.b | 20.7 sec          | 1 min 54.3 sec   | 5.5x       |
| long.b      | 10.9 sec          | 1 min 36.4 sec   | 8.8x       |
| hanoi.b     |  1.8 sec          | 1 min 18.3 sec   | 43.5x      |

## CPU profiles

### mandelbrod.b
![mandelbrod profile](https://raw.githubusercontent.com/dim13/gofys/master/profiles/mandelbrod.gif)

### long.b
![long profile](https://raw.githubusercontent.com/dim13/gofys/master/profiles/long.gif)

### hanoi.b
![hanoi profile](https://raw.githubusercontent.com/dim13/gofys/master/profiles/hanoi.gif)
