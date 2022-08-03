# Data structures and algorithms in Go

Teaching myself some Go while refreshing data structures and algorithms.

## [Data Structures](ds/)

* [Binary heap](ds/binary_heap.go), array-backed
* [Hash table](ds/hash_table.go) with linear probing for collision resolution
* [Singly linked list](ds/singly_linked_list.go) with head and tail pointers
  * [Stack](ds/stack.go)
  * [Queue](ds/queue.go)

## [Algorithms](algs/)

### Sorting

* [Bubble sort](algs/bubble_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time, O(1) swaps in best case (sorted input), O(1) aux memory.
* [Insertion sort](algs/insertion_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time and O(1) swaps in best case (sorted input), O(1) memory. 
  * [Shellsort](algs/insertion_sort.go)
* [Quicksort](algs/quicksort.go)
  * [Iterative quicksort](algs/quicksort.go)
  * [Tail-recursive quicksort](algs/quicksort.go)
* [Selection sort](algs/selection_sort.go)