package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := [][]*Node{[]*Node{root}}
	for len(queue) > 0 {
		row := queue[0]
		queue = queue[1:]

		newRow := []*Node{}

		var node *Node
		for i := 0; i < len(row)-1; i++ {
			node = row[i]
			node.Next = row[i+1]
			if node.Left != nil {
				newRow = append(newRow, node.Left)
			}
			if node.Right != nil {
				newRow = append(newRow, node.Right)
			}
		}
		node = row[len(row)-1]
		node.Next = nil
		if node.Left != nil {
			newRow = append(newRow, node.Left)
		}
		if node.Right != nil {
			newRow = append(newRow, node.Right)
		}

		if len(newRow) > 0 {
			queue = append(queue, newRow)
		}
	}

	return root
}
