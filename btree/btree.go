package btree

import "strings"

type node struct {
	Leftnode  *node
	Value     string
	Rightnode *node
}

func Init(treeString string) node {

	btree := createBtree(treeString)
	return btree
}

func createBtree(treeString string) node {

	//Alway take left node/tree string starting point as 0
	//Always take Right node/tree ending point as len()-1
	//fmt.Println("String received: ", treeString)
	processingLeftTree := true
	parenCounter := 0
	var rootnodeString, leftTreeString, rightTreeString string
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
					rootnodeString = strings.Trim(treeString[index+1:nextParenIndex], " ")
					index = nextParenIndex - 1
				case false:
					rightTreeString = treeString[nextParenIndex:]
				}

			}
		}
	}
	var lnode, rnode, rootnode node

	//Check if each left string and right string has only one left '(' and one ')'
	//This will be recursion stopper
	if strings.Count(leftTreeString, "(") == 1 && strings.Count(leftTreeString, ")") == 1 {
		leftnodeValue := strings.Trim(leftTreeString[1:len(leftTreeString)-1], " ")
		//Create node with this value
		lnode = node{nil, leftnodeValue, nil}
	} else {

		//If the string represents a branch instead of a node, pass the string to createBtree again
		//When passing make sure you strip left most '(' and right most ')'
		lnode = createBtree(leftTreeString[1 : len(leftTreeString)-1])
	}

	if strings.Count(rightTreeString, "(") == 1 && strings.Count(rightTreeString, ")") == 1 {
		rightnodeValue := strings.Trim(rightTreeString[1:len(rightTreeString)-1], " ")
		//Create node with this value
		rnode = node{nil, rightnodeValue, nil}
	} else {

		//If the string represents a branch instead of a node, pass the string to createBtree again
		//When passing make sure you strip left most '(' and right most ')'
		rnode = createBtree(rightTreeString[1 : len(rightTreeString)-1])
	}

	//Create root node and attach left and right nodes, and return root node address
	rootnode = node{&lnode, rootnodeString, &rnode}

	return rootnode
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
