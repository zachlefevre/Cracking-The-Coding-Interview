# Given two strings, write a method to decide if one is a permutation of the other

## Additional Possible Constraints:
1. What if no modifications can be made to the arrays?
2. What if no extra data structures are allowed to be used?
3. What if only an approximate solution is needed?

## Ideas:
* [Assuming modifications are allowed, and extra data structures are not allowed] The first string can be iterated over. For every element of the first string, iterate over the entire second second string. If the iteration character of the second string matches the interation character of the first string, then remove it.
1. If the string is not "" when the end of the first string is reached then it is not a permutation.
2. If the string is "" before the end of the first string is reached, then it is not a permutation
    
* [Assuming modifications are allowed and extra data structures are not allowed] The two strings could be sorted, and if the sorted strings are equivalent to one another, the second string was a permutation of the first.

* [Assuming modifications are allowed and extra data structures are not allowed] The two strings could be sorted **out of place**, and if the sorted strings are equivalent to one another, the second string was a permutation of the first.


* [Assuming modifications are not allowed, and extra data structures are allowed] A hash table could be created which maps an element of the array to its frequency. For each element of the first array, increment its corresponding frequency in the map. Then the second string would be iterated over. For each element of the second array, decrement its corresponding frequency in the map
1. If every element of the int array is 0, then the second string is a permutation of the first. Otherwise it is not.

* [Assuming modifications are not allowed, and extra data structures are allowed and only an approximate solution is required] A Count Min Sketch data structure could be used to track frequency of the elements in the first array via lossy compression. For each element of the first array, increment its corresponding frequency in the CMS. Then the second string would be iterated over. For each element of the second array, decrement its corresponding frequency in the map
1. If every element of the int array is 0, then the second string is a permutation of the first. Otherwise it is not

