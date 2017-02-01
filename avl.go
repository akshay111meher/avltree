package main

import (
	"fmt"
)

type avlNode struct {
	Key            int
	Height         int
	Lchild, Rchild *avlNode
}

func leftRotate(root *avlNode) *avlNode {
	node := root.Rchild
	// fmt.Println(node.Key)
	root.Rchild = node.Lchild
	node.Lchild = root

	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	node.Height = max(height(node.Rchild), height(node.Lchild)) + 1
	return node
}

func leftRigthRotate(root *avlNode) *avlNode {
	root.Lchild = leftRotate(root.Lchild)
	root = rightRotate(root)
	return  root
}

func rightRotate(root *avlNode) *avlNode {
	node := root.Lchild
	root.Lchild = node.Rchild
	node.Rchild = root
	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	node.Height = max(height(node.Lchild), height(node.Rchild)) + 1
	return node
}

func rightLeftRotate(root *avlNode) *avlNode {
	root.Rchild = rightRotate(root.Rchild)
	root = leftRotate(root)
	return  root
}

func height(root *avlNode) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func hasElement(root *avlNode, value int) bool  {
  if root == nil{
    return false
  }else if root.Key==value {
    return true
  }else if value > root.Key {
    return hasElement(root.Rchild,value)
  }else{
    return hasElement(root.Lchild,value)
  }
}

func sizeOfTree(root *avlNode) int {
  if root == nil{
    return 0
  }
  return 1 + sizeOfTree(root.Lchild) + sizeOfTree(root.Rchild)
}

func kthsmallestNode(root *avlNode, k int) int {
  if root ==nil {
    return -1
  }
  no_left := sizeOfTree(root.Lchild)
  if no_left == k-1 {
    return root.Key
  }else if no_left > k-1{
    return kthsmallestNode(root.Lchild,k)
  }else{
    return kthsmallestNode(root.Rchild,k - no_left - 1)
  }
}
func kthlargestNode(root *avlNode, k int) int {
  if root == nil {
    return -1
  }
  no_right := sizeOfTree(root.Rchild)
  if no_right == k-1 {
    return root.Key
  }else if no_right > k-1 {
    return kthlargestNode(root.Rchild,k)
  }else{
    return kthlargestNode(root.Lchild,k - no_right - 1)
  }
}

func maximum(root *avlNode) int  {
  if root.Rchild == nil {
    return root.Key
  }else{
    return maximum(root.Rchild)
  }
}
func minimum(root *avlNode) int  {
  if root.Lchild == nil {
    return root.Key
  }else{
    return minimum(root.Lchild)
  }
}
func insert(root *avlNode, key int) *avlNode {
	if root == nil {
		root = &avlNode{key, 0, nil, nil}
		root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
		return root
	}

	if key < root.Key {
		root.Lchild = insert(root.Lchild, key)
		if height(root.Lchild)-height(root.Rchild) == 2 {
			if key < root.Lchild.Key {
				root = rightRotate(root)
			} else {
				root = leftRigthRotate(root)
			}
		}
	}

	if key > root.Key {
		root.Rchild = insert(root.Rchild, key)
		if height(root.Rchild)-height(root.Lchild) == 2 {
			if key > root.Rchild.Key {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root)
			}
		}
	}

	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	return root
}


type action func(node *avlNode)

func inOrder(root *avlNode, action action) {
	if root == nil {
		return
	}

	inOrder(root.Lchild, action)
	action(root)
	inOrder(root.Rchild, action)
}

func preOrder(root *avlNode, action action) {
	if root == nil {
		return
	}

  action(root)
	preOrder(root.Lchild, action)
	preOrder(root.Rchild, action)
}
func postOrder(root *avlNode, action action) {
	if root == nil {
		return
	}

	postOrder(root.Lchild, action)
  postOrder(root.Rchild, action)
	action(root)
}

func main() {
	var root *avlNode
	var T,temp int
  fmt.Println("Start")
  fmt.Scan(&T)
  for i := 0; i < T; i++ {
    fmt.Scan(&temp)
    root = insert(root,temp)
  }
  fmt.Println("postOrder traversal")
  postOrder(root,func (root *avlNode)  {
    fmt.Println(root.Key)
  })

  fmt.Println("inOrder traversal")
  inOrder(root,func (root *avlNode)  {
    fmt.Println(root.Key)
  })

  fmt.Println("preOrder traversal")
  preOrder(root,func (root *avlNode)  {
    fmt.Println(root.Key)
  })

  fmt.Println(kthlargestNode(root,2),kthsmallestNode(root,2))

}
