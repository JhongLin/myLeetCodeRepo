//https://leetcode.com/problems/count-complete-tree-nodes/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func countNodes(root *TreeNode) int {
    if root == nil {    return 0    }
    var rightMost int
    rightMost = 0
    lack := 0
    done := false
    dfs(root, 0, &rightMost, true, &lack, &done)
    ans := int(math.Pow(2, float64(rightMost+2) ))-1
    ans -= lack
    return ans
}

func dfs(node *TreeNode, depth int, rightMost *int, rightStraight bool, lack *int, done *bool) {
    if *done {return}
    if node.Right !=nil{
        if  rightStraight {
            dfs(node.Right, depth+1, rightMost, rightStraight, lack, done)
        }else{
            if depth == *rightMost {
                *done = true;
                return
            }else{
                dfs(node.Right, depth+1, rightMost, false, lack, done)
            }
        }
    }else{ //no right Child
        if rightStraight {
            *rightMost = depth
            *lack += 1
        }else{
            *lack += 1
        }
    }
    
    if *done {return}
    
    if node.Left != nil{
        if depth == *rightMost {
            *done = true
            return
        }else{
            dfs(node.Left, depth+1, rightMost, false, lack, done)
        }
    }else { // no left child
        *lack += 1

    }

}