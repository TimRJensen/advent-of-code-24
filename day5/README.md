# Solution for Advent of Code 24 day 5 

## Task 1
Append all the rules for a digit in a lookup table and simply traverse the list of inputs, while making sure that i+1 element is in i's ruleset.

## Task 2
Ignore all the valid inputs from task one and simply sort the rest with a slightly custom quicksearch. This solution uses Hoare's partition.

## Deliverable
| | Solution | Benchmark (ns/op) | Complexity |
| - | :------: | :-------------: | :-----------: |
|Task 1 | 5391 &#10003; | 0.0000372  | O(*n*\**m*\*log*k*) |
|Task 2 | 6142 &#10003; | 0.0000477  | O(*n*\**m*\*log*k*+*n*\**m*\*log*m*) |

*n*=number of inputs; *m*=length of a input; *k*=length of a rule;
