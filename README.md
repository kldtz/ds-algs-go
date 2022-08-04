# Data structures and algorithms in Go

Teaching myself some Go while refreshing data structures and algorithms.

## [Data Structures](ds/)

* [Binary heap](ds/binary_heap.go), array-backed
* [Hash table](ds/hash_table.go) with linear probing for collision resolution
* [Singly linked list](ds/singly_linked_list.go) with head and tail pointers
  * [Stack](ds/stack.go)
  * [Queue](ds/queue.go)

## [Algorithms](algs/)

### [Sorting](algs/sorts/)

* [Bubble sort](algs/sorts/bubble_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time, O(1) swaps in best case (sorted input), O(1) aux memory.
  * [Cocktail shaker sort](algs/sorts/bubble_sort.go): slight improvement by moving small elements from the end to the beginning faster in back passes.
  * [Comb sort](algs/sorts/bubble_sort.go): improves over bubble sort by eliminating turtles, small values at the end of the list.
* [Counting sort](algs/sorts/counting_sort.go)
* [Heapsort](ds/binary_heap.go)
* [Insertion sort](algs/sorts/insertion_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time and O(1) swaps in best case (sorted input), O(1) memory. 
  * [Shellsort](algs/sorts/insertion_sort.go)
* [Quicksort](algs/sorts/quicksort.go)
  * [Iterative quicksort](algs/sorts/quicksort.go)
  * [Tail-recursive quicksort](algs/sorts/quicksort.go)
* [Selection sort](algs/sorts/selection_sort.go)