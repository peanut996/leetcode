package main

// Return Max int
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Return min int
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**************************************
 * Functions for tree.
 */

// Create a tree with two nil son.
func createTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil}
}

//Function for creating a Complete Binary Tree from Slice
func createTreefromSlice(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	head := createTreeNode(nums[0])
	queue := []*TreeNode{}
	queue = append(queue, head)
	for i := 1; i < length; {

		//判断队列是否为空
		if len(queue) != 0 {

			//取出头部元素
			node := queue[0]
			queue = queue[1:]

			//判定是否为空结点
			if nums[i] != -1 {
				node.Left = createTreeNode(nums[i])
				queue = append(queue, node.Left)
			}
			i++

			// 手动加一次i<length判定
			if i == length {
				break
			}

			//判断是否为空结点
			if nums[i] != -1 {
				node.Right = createTreeNode(nums[i])
				queue = append(queue, node.Right)
			}
			i++
		} else {
			break
		}
		//取出头元素
	}
	return head
}
func getSliceFromTree(root *TreeNode) []int {
	nums := []int{}
	if root == nil {
		return nums
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		// 取值
		if node != nil {
			//结点入队列
			nums = append(nums, node.Val)
			//去除最外层
			if !(node.Left == nil && node.Right == nil) {
				queue = append(queue, node.Left, node.Right)
			}
		} else {
			nums = append(nums, -1)
		}

	}
	return nums
}

// Return the max height of a tree.
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := height(root.Left), height(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
