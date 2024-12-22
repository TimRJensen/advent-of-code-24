# Solution for Advent of Code 24 day 4 

## Task 1
Asumming that each line of the input is the same length, the lines are concatenated and the length is used to validate a sequence when one reaches an 'X' in the list.
## Task 2
Same as task 1, except validation occurs when one reaches a 'A' in the list. Then it's simply a matter of testing whether the diagonals are valid sequences.

## Deliverable
| | Solution | Benchmark (ns/op) |
| - | :------: | :-------------: |
|Task 1 | 2562 &#10003; | 0.001175  |
|Task 2 | 1902 &#10003; | 0.001154  |
