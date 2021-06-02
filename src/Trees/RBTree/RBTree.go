package RBTree

func rbtInsert(tree *RBTree, node *Node) {
    currNode := nil
    subTreeRoot := tree.root//each iteration down the branches treats the current node as the root of a subtree

    for subTreeRoot != nil {//interate until leaves
        currNode = subTreeRoot
        if node.currNode.key < subTreeRoot.currNode.key {
            subTreeRoot = subTreeRoot.leftChild
        } else {
            subTreeRoot = subTreeRoot.rightChild
        }
    }

    node.parent = currNode

    if currNode == nil {
        tree.root = node
    } else if node.key < currNode.key {
        currNode.leftChild = node
    } else {
        currNode.rightChild = node
    }
    
    node.leftChild = nil
    node.rightChild = nil
    node.colour = RED

    rbtInsertFix(tree, node)
}

//fix RBT property violations after insertion
func rbtInsertFix(tree *RBTree, node *Node) {
    currNode := nil
    for node.parent.colour == RED {
        if node.parent == node.parent.parent.leftChild {
            currNode = node.parent.parent.rightChild
            if currNode.colour == RED {  //CASE 1
                node.parent.colour = BLACK
                currNode.colour = BLACK
                node.parent.parent.colour == RED
                node = node.parent.parent
            } else if node == node.parent.rightChild { //CASE 2
                node = node.parent
                rotateLeft(tree, node)
                node.parent.colour = BLACK
                node.parent.parent.colour = BLACK
                rotateRight(tree, node.parent.parent)
            } else { //CASE 3
                node = node.parent
                rotateRight(tree, node.parent.parent)
                node.parent.colour = BLACK
                node.parent.parent.colour = BLACK
                rotateLeft(tree, node)
            } 
        }
    }
    tree.root.colour = BLACK
}

func rbtDelete(tree *RBTree, node *Node) {

    
    rbtDeleteFix(tree, node)
}

func rbtDeleteFix(tree *RBTree, node *Node) {
    
}

func rotateLeft(tree *rbTree, node *Node) {
    currNode := node.rightChild
    node.rightChild = currNode.leftChild

    if currNode.leftChild != nil {
        currNode.leftChild.parent = node
    }
    currNode.parent = node.parent//link parents of curren node and new node
    if node.parent == nil {
        tree.root = currNode
    } else if node == node.parent.leftChild{
        node.parent.leftChild = currNode
    } else {
        node.parent.rightChild = currNode
    }
    currNode.leftChild = node
    node.parent = currNode
}

func rotateRight() {
    currNode := node.leftChild
    node.leftChild = currNode.rightChild

    if currNode.rightChild != nil {
        currNode.rightChild.parent = node
    }
    currNode.parent = node.parent
    if node.parent == nil {
        tree.root = currNode
    } else if node == node.parent.rightChild {
        node.parent.rightChild = currNode
    } else {
        node.parent.leftChild = currNode
    }
    currNode.rightChild = node
    node.parent = currNode
}


//recurse down all subtrees in inorder fashion
func inOrderTraversal(root *Node) {
    if root == nil {
        return
    }
    inOrderTraversal(root.leftChild)
    fmt.Println(root.data)
    inOrderTraversal(root.rightChild)
}


