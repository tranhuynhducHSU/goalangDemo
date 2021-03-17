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

func addElement(nodeNew Node, indexParen int, nodes []Node) []Node {
	if nodeNew.firstChar == nodes[indexParen].firstChar {
		nodes[indexParen].arr = append(nodes[indexParen].arr, nodeNew.arr[0])
	} else if nodeNew.firstChar > nodes[indexParen].firstChar {
		if nodes[indexParen].right == 0 {
			nodes[indexParen].right = len(nodes)
			nodes = append(nodes, nodeNew)
		} else {
			nodes = addElement(nodeNew, nodes[indexParen].right, nodes)
		}

	} else if nodeNew.firstChar < nodes[indexParen].firstChar {
		if nodes[indexParen].left == 0 {
			nodes[indexParen].left = len(nodes)
			nodes = append(nodes, nodeNew)
		} else {
			nodes = addElement(nodeNew, nodes[indexParen].left, nodes)
		}
	}
	return nodes
}

func printTree(nodes []Node) {
	arr := make([][]int, 0)
	arrTemp := make([]int, 0)

	arrTemp = append(arrTemp, 0)
	arr = append(arr, arrTemp)

	for j := 0; j < len(nodes)/2; j++ {
		arrStr := make([]int, 0)
		end := true
		for i := 0; i < len(arr[j]); i++ {
			if arr[j][i] == -1 || nodes[arr[j][i]].left == 0 {
				arrStr = append(arrStr, -1)
			} else {
				arrStr = append(arrStr, nodes[arr[j][i]].left)
				end = false
			}
			if arr[j][i] == -1 || nodes[arr[j][i]].right == 0 {
				arrStr = append(arrStr, -1)
			} else {
				arrStr = append(arrStr, nodes[arr[j][i]].right)
				end = false
			}
		}
		if end == true {
			break
		}
		arr = append(arr, arrStr)
		
	}
	// for i:=len(arr);i>= len(arr);i-- {
	for i:=0;i< len(arr);i++{

		res:=""
		xx := i
		xy := i*2-1
		for j:=0;j<len(arr[i]);i++{
			res+=
		}
		fmt.Println()
	}
}

func main() {
	// str := ""
	// fmt.Print("Nhập mảng các từ [',']: ")
	// fmt.Scan(&str)
	// str := "h,d,q,t,sy,qm,aa,tq,da,af,hg,az,qq,zz,ah,wr,hd,ad"
	str := "5,55,7,4,8,9,3,6,1,2,5,3,6,9,7,45,0,7,4"

	arr := strings.Split(str, ",")

	//Phần tử đầu tiên
	firstChar := strings.ToUpper(string(arr[0]))
	nodes := make([]Node, 0)
	nodes = append(nodes, Node{firstChar, make([]string, 0), 0, 0})
	nodes[0].arr = append(nodes[0].arr, arr[0])

	for i := 1; i < len(arr); i++ {
		//khoi tao node add str vo arr
		nodeTemp := Node{"", make([]string, 0), 0, 0}
		nodeTemp.firstChar = strings.ToUpper(string(arr[i][0]))
		nodeTemp.arr = append(nodeTemp.arr, arr[i])
		nodes = addElement(nodeTemp, 0, nodes)
	}

	for i := 0; i < len(nodes); i++ {
		fmt.Println(i, nodes[i])
	}

	printTree(nodes)
}
