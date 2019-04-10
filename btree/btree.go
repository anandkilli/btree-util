package btree

import "strings"

type Node struct {
	LeftNode  *Node
	Value     string
	RightNode *Node
}

func Init(treeString string) Node {

	btree := createBtree(treeString)
	return *btree
}

func createBtree(treeString string) *Node {

	//Alway take left Node/tree string starting point as 0
	//Always take Right Node/tree ending point as len()-1
	//fmt.Println("String received: ", treeString)
	processingLeftTree := true
	parenCounter := 0
	var rootNodeString, leftTreeString, rightTreeString string
	var nextParenIndex int
	for index := 0; index < len(treeString); index++ {

		if strings.Compare(string(treeString[index]), "(") == 0 {

			// Increate the conter for left Paranthesis '('
			parenCounter++

		} else if strings.Compare(string(treeString[index]), ")") == 0 {

			// Decrease the conter for right Paranthesis ')'
			parenCounter--
			//When it's zero, you found a branch like "((3) 4 (6))"
			if parenCounter == 0 {

				switch processingLeftTree {
				case true:
					leftTreeString = treeString[0 : index+1]
					processingLeftTree = false
					nextParenIndex = getIndexAfter(treeString, "(", index)
					rootNodeString = strings.Trim(treeString[index+1:nextParenIndex], " ")
					index = nextParenIndex - 1
				case false:
					rightTreeString = treeString[nextParenIndex:]
				}

			}
		}
	}
	var lNode, rNode, rootNode *Node

	//Check if each left string and right string has only one left '(' and one ')'
	//This will be recursion stopper
	if strings.Count(leftTreeString, "(") == 1 && strings.Count(leftTreeString, ")") == 1 {
		leftNodeValue := strings.Trim(leftTreeString[1:len(leftTreeString)-1], " ")
		//if leftNodeValue is a blank string, make lNode nil
		if len(leftNodeValue) == 0 {
			lNode = nil
		} else {
			//Create Node with this value
			lNode = &Node{nil, leftNodeValue, nil}
		}
	} else {

		//If the string represents a branch instead of a Node, pass the string to createBtree again
		//When passing make sure you strip left most '(' and right most ')'
		lNode = createBtree(leftTreeString[1 : len(leftTreeString)-1])
	}

	if strings.Count(rightTreeString, "(") == 1 && strings.Count(rightTreeString, ")") == 1 {
		rightNodeValue := strings.Trim(rightTreeString[1:len(rightTreeString)-1], " ")
		//if rightNodeValue is a blank string, make lNode nil
		if len(rightNodeValue) == 0 {
			rNode = nil
		} else {
			//Create Node with this value
			rNode = &Node{nil, rightNodeValue, nil}
		}
	} else {

		//If the string represents a branch instead of a Node, pass the string to createBtree again
		//When passing make sure you strip left most '(' and right most ')'
		rNode = createBtree(rightTreeString[1 : len(rightTreeString)-1])
	}

	//Create root Node and attach left and right Nodes, and return root Node address
	rootNode = &Node{lNode, rootNodeString, rNode}

	return rootNode
}

func getIndexAfter(str string, match string, index int) int {

	temp := str[:index]
	str = str[index:]

	newIndex := strings.Index(str, match)

	if newIndex == -1 {
		return newIndex
	} else {
		return len(temp) + newIndex
	}
}
