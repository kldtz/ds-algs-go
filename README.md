# Data structures and algorithms in Go

Teaching myself some Go while refreshing data structures and algorithms.

## [Data Structures](ds/)

* [Binary heap](ds/binary_heap.go), array-backed
* [Binary tree](ds/binary_tree.go)
  * [Red-black tree](ds/red_black_tree.go)
* [Hash table](ds/hash_table.go) with linear probing for collision resolution
* [Singly linked list](ds/singly_linked_list.go) with head and tail pointers
  * [Stack](ds/stack.go)
  * [Queue](ds/queue.go)
* [Suffix tree](ds/suffix_tree.go) constructed in linear time following Ukkonen (1995)

## [Algorithms](algs/)

### [Sorting](algs/sorting/)

* [Bubble sort](algs/sorting/bubble_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time, O(1) swaps in best case (sorted input), O(1) aux memory.
  * [Cocktail shaker sort](algs/sorting/bubble_sort.go): slight improvement by moving small elements from the end to the beginning faster in back passes.
  * [Comb sort](algs/sorting/bubble_sort.go): improves over bubble sort by eliminating turtles, small values at the end of the list.
* [Counting sort](algs/sorting/counting_sort.go)
* [Heapsort](ds/binary_heap.go)
* [Insertion sort](algs/sorting/insertion_sort.go): O(n²) comparisons and swaps in worst and average case, O(n) time and O(1) swaps in best case (sorted input), O(1) memory. 
  * [Shellsort](algs/sorting/insertion_sort.go)
  * [Stack-based insertion sort](algs/sorting/insertion_sort.go)
* [Merge sort](algs/sorting/merge_sort.go)
* [Quicksort](algs/sorting/quicksort.go)
  * [Iterative quicksort](algs/sorting/quicksort.go)
  * [Tail-recursive quicksort](algs/sorting/quicksort.go)
* [Selection sort](algs/sorting/selection_sort.go)


### [String](algs/string/)

* [Longest common substring](algs/string/lcs.go)

---

Ukkonen, E. (1995). On-line construction of suffix trees. Algorithmica, 14(3), 249-260.