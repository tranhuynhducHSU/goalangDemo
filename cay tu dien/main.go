package main

import (
	"fmt"
	"strings"
)

type Node struct {
	firstChar string
	arr       []string
	left      int
	right     int
}

func addNode(elem string, nodes []Node) []Node {
	//st:=0
	firstChar := strings.ToUpper(string(elem[0]))
	node := Node{"", make([]string, 0), 0, 0}
	index := 0
	if len(nodes) == 0 {
		node.firstChar = firstChar
		node.arr = append(node.arr, elem)
		nodes = append(nodes, node)
	} else {
		if firstChar == nodes[index].firstChar {
			nodes[index].arr = append(nodes[index].arr, elem)
		} else if firstChar > nodes[index].firstChar {
			if nodes[index].right == 0 {
				nodes[index].right = index
				nodes[index].arr = append(nodes[index].arr, elem)
			}

		} else if firstChar < nodes[index].firstChar {
			if nodes[index].left == 0 {
				nodes[index].left = index
				nodes[index].arr = append(nodes[index].arr, elem)
			}

		}
	}

	return nodes
}

func findNodeNil(node Node, nodes []Node) (viTri int) {
	for {
		if 

		if true {
			break
		}
	}
	return
}

func main() {
	dictionary := make([]Node, 0)
	dictionary = addNode("duc", dictionary)
	dictionary = addNode("dung", dictionary)

	fmt.Print(dictionary)
}
