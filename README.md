# OpGame

## Overview

You will write a Go program to solve a simple mathematical puzzle.  The program must efficiently find the binary operators that complete the equation.  For this problem the only operators you need to consider are addition, subtraction, multiplication, and division (when the divisor is non-zero and the remainder is zero).  For example, given `1 O 2 = 3` the only operator that works to put in the circle is "+".  So `1 + 2 = 3` solves the problem.

For simplicity we will not use the normal operator precedence.  Instead there will be no precedence thus expressions are evaluated strictly from left to right.  For example, given `1 O 2 O 3 = 1`,  your goal is to find two operators that satisfy the equation.  The solution would be `1 + 2 / 3 = 1`.  Notice that this is computed as `(1+2) / 3 = 1` and not `1 + (2/3)`.

Generating your own example problems is easy.  Write an equation (being careful to evaluate it strictly left to right) and then remove the operators and supply that to your program.

## Example Output

```shell
$ echo "3 1 2" | go run .
3-1=2

$ echo "5 4 2 22" | go run .
5*4+2=22

$ (echo "3 1 2" && echo "9 2 18") | go run .
3-1=2
9*2=18

$ echo "6 2 3 4" | go run .
6*2/3=4

$ go run . < testdata/basic.txt
3 - 1 = 2

9 + 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6
```

With multiple solutions it looks like:

```shell
$ echo "7 3 3 7" | go run . -all
7+3-3=7, 7-3+3=7, 7*3/3=7

$ go run . -all < testdata/basic.txt
3 - 1 = 2

9 + 0 = 9, 9 - 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6, 2 * 3 * 1 = 6, 2 * 3 / 1 = 6
```

## Hints

- Commit and push often so you do not loose any of your work.
- Use the `Makefile`.  For example, you can run tests with `make test`.  You can make sure you code is properly formatted with `make fix`.
