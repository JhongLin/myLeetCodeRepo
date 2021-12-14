//https://leetcode.com/problems/all-paths-from-source-to-target/

func allPathsSourceTarget(graph [][]int) [][]int {
    n:=len(graph)
    ans := make([][]int, 0)
    stack := make([]int, 0)
    numOfStack := 0
    stackRoutes := make([][]int, 0)
    for i:=0 ; i<len(graph[0]) ; i++ {
        temp := make([]int, 1)
        temp[0] = 0
        stackRoutes = append(stackRoutes, temp)
        stack = append(stack, graph[0][i])
        numOfStack++
    }
    for numOfStack>0 {
        curNum := stack[numOfStack-1]
        curRoute := stackRoutes[numOfStack-1]
        numOfStack--
        stack = stack[0:numOfStack]
        stackRoutes = stackRoutes[0:numOfStack]
        curRoute = append(curRoute, curNum)
        if curNum == n-1 {
            ans = append(ans, curRoute)
            continue
        }
        for i:=0 ; i<len(graph[curNum]) ; i++ {
            //target := graph[curNum][i]
            length := len(curRoute)
            copyRoute := make([]int, length)
            for j:=0 ; j<length ; j++ {copyRoute[j]=curRoute[j]}
            stackRoutes = append(stackRoutes, copyRoute)
            numOfStack++
            stack = append(stack, graph[curNum][i])
        }
    }
    return ans
}