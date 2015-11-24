# gofys
Go f\*ck your self -- yet another optimizing BrainF\*ck interpreter in Go

## tape storage
- static: 32k including 1k margin on left side (default)
- dynamic: allocates in 1k chunks as required on access

## code optimization
- [x] Contraction
- [x] Clear loops
- [x] Copy loops
- [x] Multiplication loops
- [x] Scan loops (kind of)
- [x] Operation offsets

Reference: http://calmerthanyouare.org/2015/01/07/optimizing-brainfuck.html

## some rough results

| Program     | with optimization | w/o optimization | speed gain |
| ----------- | -----------------:| ----------------:| ----------:|
| madelbrot.b | 20.7 s            | 1 min 54.3 sec   | 5.5x       |
| long.b      | 10.9 s            | 1 min 36.4 sec   | 8.8x       |
| hanoi.b     |  1.8 s            | 1 min 18.3 sec   | 43.5x      |
