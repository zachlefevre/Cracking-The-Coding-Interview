### Array List
An ArrayList is a resizable array. Generally elements are added to an array, and when the capacity is reached, the size of the array is doubled.
* When the array is doubled, insertion takes O(n)
* When the array is not at capacity, insertion takes O(1).
* O(n) A(1)

## String Builder
If a sequence of strings of length *i* are concatenated together into a single array, the algorithm would iterate through the entire array, each time a new string is added. Therefore the complexity is O(xn^2), or just O(n^2).  
1 + 2 + 3 + ... + n = n(n+1)/2, which reduces to a complexity of O(n^2)
