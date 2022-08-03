# Data structures and algorithms in Go

Teaching myself some Go while refreshing data structures and algorithms.

## [Data Structures](pkg/ds/)

* [Binary heap](pkg/ds/binary_heap.go), array-backed
* [Hash table](pkg/ds/hash_table.go) with linear probing for collision resolution
* [Singly linked list](pkg/ds/singly_linked_list.go) with head and tail pointers
  * [Stack](pkg/ds/stack.go)
  * [Queue](pkg/ds/queue.go)

## [Algorithms](pkg/algs/)

### Sorting

* [Bubble sort](pkg/algs/bubble_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time, O(1) swaps in best case (sorted input), O(1) aux memory.
* [Insertion sort](pkg/algs/insertion_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time and O(1) swaps in best case (sorted input), O(1) memory. 
  * [Shellsort](pkg/algs/insertion_sort.go)
* [Quicksort](pkg/algs/quicksort.go)
  * [Iterative quicksort](pkg/algs/quicksort.go)
  * [Tail-recursive quicksort](pkg/algs/quicksort.go)
* [Selection sort](pkg/algs/selection_sort.go)