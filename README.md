list of problem: https://mica-wineberry-d80.notion.site/Code-Challenge-9f10e4d75199422082ea5e503f25067e

Case 1: Unearthing's the Philosopher Stone

for n = 1, the diagonal will be 1

for n = 2, the diagonal will be 1, 2, 3, 4

for n = 3, the diagonal will be 1, 3, 5, 7, 9

| 7 | 8 | 9 |
|---|---|---|
| 6 | 1 | 2 |
| 5 | 4 | 3 |

for n = 4, the diagonal will be 1, 2, 3, 4, 7, 10, 13, 16

| 7  | 8  | 9  | 10 |
|----|----|----|----|
| 6  | 1  | 2  | 11 |
| 5  | 4  | 3  | 12 |
| 16 | 15 | 14 | 13 |

let's take diagonal from n = 2 and n = 4, notice that all the diagonal in n = 2 is also shown in n = 4

| n | diagonal  | 
|----|----|
| 2  | **1, 2, 3, 4** |
| 4  | **1, 2, 3, 4**, 7, 10, 13, 16 |

if we continue to n = 6, we will notice the similar pattern as well, where all diagonal number in n = 4 also shown in diagonal number in n = 6

| 21 | 22 | 23 | 24 | 25 | 26 |
|----|----|----|----|----|----|
| 20 | 7  | 8  | 9  | 10 | 27 |
| 19 | 6  | 1  | 2  | 11 | 28 |
| 18 | 5  | 4  | 3  | 12 | 29 |
| 17 | 16 | 15 | 14 | 13 | 30 |
| 36 | 35 | 34 | 33 | 32 | 31 |

| n | diagonal  | 
|----|----|
| 4  | **1, 2, 3, 4, 7, 10, 13, 16** |
| 6  | **1, 2, 3, 4, 7, 10, 13, 16**, 21, 26, 31, 36 |

also notice, for the new diagonal number for each N, the difference for each new diagonal number is always the same. the smallest new diagonal number also has the same difference with the last element from N - 2

for n = 4, the difference in new diagonal number (compared to diagonal number in n = 2) is 3 (7, 10, 13, 16), and the difference from the last element of N - 2 is also 3 (4 and 7)

while for n = 6, the difference in new diagonal number is 5 (21, 26, 31, 36)

based on this finding we can represent sum new diagonal number for N as this

$sumNewDiagonalNumber(N) = (N-2)^2 + (N - 1) + (N - 2) ^ 2 + 2(N - 1) +  (N - 2) ^ 2 + 3(N - 1) +  (N - 2) ^ 2 + 4(N - 1)$

$sumNewDiagonalNumber(N) = 4*(N - 2) ^ 2 + 10(N-1)$

so to determine what is the sum of diagonal for n, if we know the sum of diagonal number for n - 2

$sumDiagonal(N) = sumDiagonal(N - 2) + sumNewDiagonalNumber(N)$

$sumDiagonal(N) = sumDiagonal(N - 2) + 4*(N - 2) ^ 2 + 10(N-1)$
