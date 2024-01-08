package gol_datastructures

type SllNode[T comparable] struct {
	next *SllNode[T]
	val  T
}

type DllNode[T comparable] struct {
	prev *DllNode[T]
	next *DllNode[T]
	val  T
}

type BSTNode[T comparable] struct {
	left *BSTNode[T]
	right *BSTNode[T]
	val T
}

type TreeNode[T comparable] struct {
	children []*TreeNode[T]
	val T
}