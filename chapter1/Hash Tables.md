# Hash Tables  
A hash table provides efficient insertion and lookup. It can be implemented over an array, BST, other.  
## Over Any Data Structure
* Compute the key's hash code. Generally an int or a long.
* Map the key's hash code to the underlying data structure.

## Over an array
* The map from the key's hash code is generally hashcode%array_len
* Each unit of the array is representative of some unit of collision space.
### Separately chained
* A __separately chained__ hash table will implement a __linked list__ on each unit of the array. An element whose hash code is represented by one of the linked lists is appended to the linked list.
* A good hash function will map evenly and chaotically.
* O(n) A(1)
### Openly Hashed
* An __openly hashed__ hash table will resolve collisions by adding an element to the next available hash bucket, called __linear probing__, or by computing a second hash to offset the element upon a collision, called __double hashing__.

## Over a Tree
* The hash code can either be a single number, which is added to the binary search tree, or the hash code is a binary sequence and the elements are added to the tree in a fashion similar to a __huffman tree__.
* BST: O(log(N))
* HT: O(M) where M is the length of the hash code.
