package main

func main() {
	
}

// 简单
// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。

// TreeNode 树节点定义
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 自己实现，遍历所有结点，得出各个子节点的最长深度，逐步返回到顶级
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	res := max(maxLeft, maxRight) + 1
	return res
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

// 官方解法1：递归
// 更加简洁
func maxDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 官方解法2：广度优先搜索
// 每次拓展下一层的时候，不同于广度优先搜索的每次只从队列里拿出一个节点，我们需要将队列里的所有节点都拿出来进行拓展，这样能保证每次拓展完的时候队列里存放的是当前层的所有节点，即我们是一层一层地进行拓展，最后我们用一个变量 ans 来维护拓展的次数，该二叉树的最大深度即为 ans
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue := append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue := queue[1:]
			if node.Left != nil {
				queue = append(queue, node.left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}