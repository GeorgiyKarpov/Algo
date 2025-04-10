package main

func deptSearch(root *TreeNode, res []*TreeNode) []*TreeNode {
	if root == nil {
		return append(res, nil)
	}
	res = deptSearch(root.Left, res)
	res = append(res, root)
	res = deptSearch(root.Right, res)
	return res
}

func isSymmetricDFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	data := deptSearch(root, []*TreeNode{})
	for i := 0; i < len(data)/2; i++ {
		if data[i] == nil && data[len(data)-i-1] == nil {
			continue
		}
		if data[i] == nil || data[len(data)-i-1] == nil {
			return false
		}
		if data[i].Val != data[len(data)-i-1].Val {
			return false
		}
	}
	return true
}
