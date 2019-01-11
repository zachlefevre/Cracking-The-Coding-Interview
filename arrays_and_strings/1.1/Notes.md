# Implement an algorithm to determine if a string has all unique characters.

## Additional Possible Constraints:
1. What if you cannot use additional data structures?
2. What if your solution must be functional?


## Ideas:
* The string characters or array elements can be added to a {binary array, hash table}, any collisions would be obvious on insertion. **O(n)**

### If _no other data structures are allowed_:
* The array elements could be sorted and then checked linearly. **O(n\*log(n) + n) = O(n*log(n))**
* The array elements can be checked against all array elements ahead of it. **O(n^2)**

### If _no side-effects are allowed_:
* The elements could be checked to be unique against every character after it. **O(n^2)**