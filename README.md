# gofys
Go f\*ck your self -- yet another optimizing BrainF\*ck interpreter in Go

## tape storage
- static: 64k including 4k margin on left side (default)
- dynamic: allocates in 4k chunks as required on access

## code optimization
- [x] Contraction
- [x] Clear loops
- [x] Copy loops
- [x] Multiplication loops
- [x] Scan loops (kind of)
- [x] Operation offsets

Reference: http://calmerthanyouare.org/2015/01/07/optimizing-brainfuck.html

## some rough results

| Program     | with optimization | w/o optimization | gain |
| ----------- | -----------------:| ----------------:| ----:|
| madelbrot.b | 38 s              | 2 min 35 sec     | 4x   |
| long.b      | 16 s              | 4 min            | 15x  |
| hanoi.b     |  3 s              | 3 min 50 sec     | 77x  |
