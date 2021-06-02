package RBTree

import (
    "fmt"
    "math/rand"
)

type Colour int

const (
    RED = iota//0
    BLACK//1
)

type Node struct {

    leftChild *Node
    rightChild *Node
    parent *Node

    colour Colour
    key int//key is the compared label as well as the data
}

type RBTree struct {
    root *Node
}

func initNode(newNode *Node, data *int) {
    newNode.leftChild = nil
    newNode.rightChild = nil
    newNode.parent = nil

    newNode.colour = RED
    newNode.key = int
}

func main() {
    var rbt RBTree

    for i := 0; i < 10; i++ {
        data := rand.Intn(100)
        newNode := Node{nil, nil, nil, BLACK, data}
        rbtInsert(rbt, newNode)
    }

    inOrderTraversal(rbt.root)
}
