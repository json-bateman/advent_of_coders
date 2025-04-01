package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	finalArr = []int{}
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	finalArr := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		newRow := []int{}
		l := len(q)
		for i := 0; i < l; i++ {
			curr := q[0]
			q = q[1:]

			if curr != nil {
				newRow = append(newRow, curr.Val)

				if curr.Left != nil {
					q = append(q, curr.Left)
				}
				if curr.Right != nil {
					q = append(q, curr.Right)
				}
			}
		}
		finalArr = append(finalArr, newRow)
	}
	return finalArr
}
