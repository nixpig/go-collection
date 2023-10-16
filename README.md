# Data Structures The Fun Way

Implementing the data structures and algorithms (presented in pseudocode) from _Data Structures The Fun Way_ in **TypeScript**, **Go** and **Rust**.

I've never written Go or Rust before this, so trying to translate the Python-esque pseudocode presented in the book into two completely new languages (and attempt to do so idomatically, albeit , probably failing) is going to be a learning exercise _(...which is kinda the point)_!

## Chapter 1: Information In Memory

### Insertion Sort

- TypeScript - [test](typescript/src/insertion-sort.test.ts) - [solution](typescript/src/insertion-sort.ts)
- Go - [test](go/insertion_sort/insertion_sort_test.go) - [solution](go/insertion_sort/insertion_sort.go)
- Rust - [test](rust/insertion_sort/insertion_sort.rs#L22) - [solution](rust/insertion_sort/insertion_sort.rs)

## Chapter 2: Binary Search

### Binary Search

- TypeScript - [test](typescript/src/binary-search.test.ts) - [solution](typescript/src/binary-search.ts)
- Go - [test](go/binary_search/binary_search_test.go) - [solution](go/binary_search/binary_search.go)
- Rust - [test](rust/binary_search/binary_search.rs#L27) - [solution](rust/binary_search/binary_search.rs)

## Chapter 3: Dynamic Data Structures

### Linked Lists

- Go - [test]() - [solution]()

## Chapter 4: Stacks And Queues

### Stacks

#### Array Stack

- Go - [test](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/array-stack/array_stack_test.go) - [solution](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/array-stack/array_stack.go)

### Queues

#### Array List Queue

- Go - [test](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/array-queue/array_queue_test.go) - [solution](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/array-queue/array_queue.go)

#### Array Ring Buffer Queue

**Note:** this one actually _isn't_ covered in the book, short of stating that _"Wrapping is a better solution, though it does require us to carefully handle indices being incremented pas the end of the array..."_

- Go - [test](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/ring-buffer/ring_buffer_test.go) - [solution](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/array-queue/array_queue_test.go)

#### Linked List Queue

- Go - [test]() - [solution]()

## Chapter 5: Binary Search Trees

- Go - [test](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/binary-search-tree/binary_search_tree_test.go) - [solution](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/binary-search-tree/binary_search_tree.go)

## Chapter 6: Tries And Adapting Data Structures

- Go - [test](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/trie/trie_test.go) - [solution](https://github.com/nixpig/data-structures-the-fun-way/blob/main/go/trie/trie.go)
