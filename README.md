# Go Binary Search Algorithm

A simple console game implementation using binary search algorithm and GO.

## Binary Search

In computer science, binary search, also known as half-interval search, logarithmic search, or binary chop, is a search algorithm that finds the position of a target value within a sorted array.

![Binary Search](https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Binary_Search_Depiction.svg/470px-Binary_Search_Depiction.svg.png)

## Searching Algorithm

**Iteration 1**
1. low = 0
2. high = 16
3. middle = (0+16)/2 = 16/2 = 8
4. 7 (target at index 4) < 14 (middle at index 8)
5. high = middle - 1 = 7

**Iteration 2**
1. low = 0
2. high = 7
3. middle = (0+7)/2 = 7/2 = 3 R 1
4. 7 (target at index 4) > 6 (middle at index 3)
5. low = middle + 1 = 4

**Iteration 3**
1. low = 4
2. high = 7
3. middle = (4+7)/2 = 11/2 = 5 R 1
4. 7 (target at index 4) < 8 (middle at index 5)
5. high = middle - 1 = 4

**Iteration 4**
1. low = 4
2. high = 4
3. middle = (4+4)/2 = 8/2 = 4
4. 7 (target at index 4) == 7 (middle at index 4)

