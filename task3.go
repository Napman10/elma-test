package main

import "fmt"

type Node struct {
	ID   int
	Link *Node
}

func CreateListWithHead() *Node {
	nodes := []*Node{
		{
			ID: 1,
		},
	}

	for i := 1; i <= 10; i++ {
		n := &Node{ID: i + 1}

		nodes = append(nodes, n)
		nodes[i-1].Link = n
	}

	return nodes[0]
}

func SPrintNode(chain string, cur *Node) string {
	if cur == nil {
		return chain
	}

	return SPrintNode(fmt.Sprintf("%s->Node %d", chain, cur.ID), cur.Link)
}

func SPrintList(head *Node) string {
	return SPrintNode("", head)
}

func CollectByCur(chain []*Node, cur *Node) []*Node {
	if cur == nil {
		return chain
	}

	return CollectByCur(append(chain, cur), cur.Link)
}

func CollectByHead(head *Node) []*Node {
	return CollectByCur([]*Node{}, head)
}

func InvertList(head *Node) *Node {
	collected := CollectByHead(head)

	collected[0].Link = nil

	for i := 1; i < len(collected); i++ {
		collected[i].Link = collected[i-1]
	}

	return collected[len(collected)-1]
}

func ShowTask3() {
	head := CreateListWithHead()
	fmt.Println("new list from head")
	fmt.Println(SPrintList(head))

	inv := InvertList(head)
	fmt.Println("inverted list from tail")
	fmt.Println(SPrintList(inv))
	fmt.Println(SPrintList(head))

	head = InvertList(inv)
	fmt.Println("invert back")
	fmt.Println(SPrintList(inv))
	fmt.Println(SPrintList(head))
}
