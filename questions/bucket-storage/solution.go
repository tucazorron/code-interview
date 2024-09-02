package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type TreeNode struct {
	address  string
	children []*TreeNode
	isFile   bool
	size     int
}

type Tree struct {
	root *TreeNode
}

type File struct {
	address string
	size    int
}

// Solution 1: Collect all files and sort by size.
func solution1(root *TreeNode) []File {
	files := []File{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.isFile {
			files = append(files, File{
				address: node.address,
				size:    node.size,
			})
			return
		}
		for _, c := range node.children {
			dfs(c)
		}
	}
	dfs(root)
	sort.Slice(files, func(i, j int) bool {
		return files[i].size > files[j].size
	})
	if len(files) > 10 {
		return files[:10]
	}
	return files
}

type FileHeap []File

func (h FileHeap) Len() int {
	return len(h)
}

func (h FileHeap) Less(i, j int) bool {
	return h[i].size < h[j].size
}

func (h FileHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FileHeap) Push(x interface{}) {
	*h = append(*h, x.(File))
}

func (h *FileHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func solution2(root *TreeNode) []File {
	h := &FileHeap{}
	heap.Init(h)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.isFile {
			heap.Push(h, File{
				address: node.address,
				size:    node.size,
			})
			if h.Len() > 10 {
				heap.Pop(h)
			}
		} else {
			for _, c := range node.children {
				dfs(c)
			}
		}
	}
	dfs(root)
	files := make([]File, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		files[i] = heap.Pop(h).(File)
	}
	return files
}

func main() {
	root := &TreeNode{
		address: "/",
		children: []*TreeNode{
			{
				address: "/file1.txt",
				isFile:  true,
				size:    100,
			},
			{
				address: "/dir1",
				children: []*TreeNode{
					{
						address: "/dir1/file2.txt",
						isFile:  true,
						size:    200,
					},
					{
						address: "/dir1/file3.txt",
						isFile:  true,
						size:    150,
					},
					{
						address: "/dir1/dir2",
						children: []*TreeNode{
							{
								address: "/dir1/dir2/file4.txt",
								isFile:  true,
								size:    300,
							},
							{
								address: "/dir1/dir2/file5.txt",
								isFile:  true,
								size:    250,
							},
						},
					},
				},
			},
			{
				address: "/dir3",
				children: []*TreeNode{
					{
						address: "/dir3/file6.txt",
						isFile:  true,
						size:    50,
					},
					{
						address: "/dir3/file7.txt",
						isFile:  true,
						size:    400,
					},
					{
						address: "/dir3/dir4",
						children: []*TreeNode{
							{
								address: "/dir3/dir4/file8.txt",
								isFile:  true,
								size:    350,
							},
							{
								address: "/dir3/dir4/file9.txt",
								isFile:  true,
								size:    450,
							},
							{
								address: "/dir3/dir4/dir5",
								children: []*TreeNode{
									{
										address: "/dir3/dir4/dir5/file10.txt",
										isFile:  true,
										size:    500,
									},
									{
										address: "/dir3/dir4/dir5/file11.txt",
										isFile:  true,
										size:    600,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(solution1(root))
	fmt.Println(solution2(root))
}
