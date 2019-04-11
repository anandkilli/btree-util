package main

import (
	"github.com/anandkilli/btree-util/btree"
)

func main() {

	treeString := "(((2)10(12))20(22))25(((28)30(32))36(40))"
	bt := btree.Init(treeString)
	btree.BtreeToHtml(bt)
	_ = deleteNode(&bt, "25")
	_ = deleteNode(&bt, "36")
	btree.HPrintln("Tree after deleting 25 and 36")
	btree.BtreeToHtml(bt)
	btree.DrawBtree()
}

func minValueNode(srcNode *btree.Node) *btree.Node {

	current := srcNode

	/* loop down to find the leftmost leaf */
	for current.LeftNode != nil {
		current = current.LeftNode
	}

	return current
}

/* Given a binary search tree and a key, this function deletes the key
   and returns the new root */
func deleteNode(root *btree.Node, key string) *btree.Node {
	// base case
	if root == nil {
		return root
	}

	// If the key to be deleted is smaller than the root's key,
	// then it lies in left subtree
	if key < root.Value {
		root.LeftNode = deleteNode(root.LeftNode, key)
	} else if key > root.Value {
		// If the key to be deleted is greater than the root's key,
		// then it lies in right subtree
		root.RightNode = deleteNode(root.RightNode, key)
	} else {
		// if key is same as root's key, then This is the node
		// to be deleted
		// node with only one child or no child
		if root.LeftNode == nil {

			temp := root.RightNode
			root.RightNode = nil
			return temp
		} else if root.RightNode == nil {

			temp := root.RightNode
			root.RightNode = nil
			return temp
		}

		// node with two children: Get the inorder successor (smallest
		// in the right subtree)
		temp := minValueNode(root.RightNode)

		// Copy the inorder successor's content to this node
		root.Value = temp.Value

		// Delete the inorder successor
		root.RightNode = deleteNode(root.RightNode, temp.Value)
	}
	return root
}
