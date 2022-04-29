func canFinish(numCourses int, prerequisites [][]int) bool {
    
    // count of incoming edges at a specific node
    // in this example, count of dependencies for a specific course ( where each course is the ith idx )
    indegrees := make([]int, numCourses)
    
    // { independent: [dependencies] }
    // { node: [edges] }
    // { prereq: [dependencies it solves ] }
    adjList := map[int][]int{}
    coursesTakenSuccessfully := 0
    
    for i := 0; i < len(prerequisites); i++ {
        preReq := prerequisites[i][1]
        course := prerequisites[i][0]
        adjList[preReq] = append(adjList[preReq], course)
        indegrees[course]++
    }
    
    // BFS
    q := []int{}
    // find ALL the independent nodes OR in other grap terms, find all the sources we can start from
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] == 0 {
            q = append(q, i)
        }
    }
    
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        coursesTakenSuccessfully++
        unlockedDependencies := adjList[dq]
        for _, dep := range unlockedDependencies {
            indegrees[dep]--
            if indegrees[dep] == 0  {
                // enqueue this source/independent node
                q = append(q, dep)
            }
        }
    }
    return coursesTakenSuccessfully == numCourses
}