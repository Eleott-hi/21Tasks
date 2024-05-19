package main

import "fmt"

type TreeNode struct {
    HasToy bool
    Left   *TreeNode
    Right  *TreeNode
}

func unrollGarland(root *TreeNode) []bool {
    if root == nil {
        return []bool{}
    }

    result := []bool{}
    queue := []*TreeNode{root}
    leftToRight := true

    for len(queue) > 0 {
        levelSize := len(queue)
        levelValues := make([]bool, levelSize)

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            index := i
            if !leftToRight {
                index = levelSize - 1 - i
            }
            levelValues[index] = node.HasToy

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
			
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, levelValues...)
        leftToRight = !leftToRight
    }

    return result
}

func main() {
    root := &TreeNode{HasToy: true,
        Left:  &TreeNode{HasToy: true, Left: &TreeNode{HasToy: true}, Right: &TreeNode{HasToy: false}},
        Right: &TreeNode{HasToy: false, Left: &TreeNode{HasToy: true}, Right: &TreeNode{HasToy: true}},
    }
    fmt.Println(unrollGarland(root)) // Output: [true, true, false, true, false, true, true]
}
