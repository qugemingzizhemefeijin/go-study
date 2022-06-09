package main

import "fmt"

//Hero ...
type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//PreOrder 前序遍历，先输出root节点，然后输出左子树，再输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//InfixOrder 先输出root的左子树，然后再输出root节点，最后输出右子树
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

//PostOrder 先输出root的左子树，再输出右子树，最后输出root节点
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

//二叉树，前序，中序和后序遍历
func main() {
	//构建二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}

	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2

	left11 := &Hero{
		No:   5,
		Name: "公孙胜",
	}
	left12 := &Hero{
		No:   6,
		Name: "关胜",
	}
	left1.Left = left11
	left1.Right = left12

	//PreOrder(root)
	//InfixOrder(root)
	PostOrder(root)
}
