# Solution for Advent of Code 24 day 2 

## Task 1
To check that the task properties holds, for ascending, one simply traverse the list from the 1st index. For descending lists, one traverse the list from last index. 
This will make the list ascending and one can apply the same logic from the first traversal.

## Task 2
Sum up the differences between each pair in a list. If that sum is between x and x\*3, where x is the length of the list minus 1, the list is valid. 
To allow for tolerance, one simply keeps tracks of how many invalid pairs seen in the list so far. Once it's above the boundary, set the lists sum > x\*3 and end the traversal.

## Deliverable
| | Solution | Benchmark (ns/op) | Complexity | 
| - | :------: | :-------------: | :-----------: |
|Task 1 | 631 &#10003; | 0.0000230  | O(*n*\**m*) |
|Task 2 | 665 &#10003; | 0.0000427  | O(*n*\**m*) |

*n*=number of inputs; *m*=length of a input;
