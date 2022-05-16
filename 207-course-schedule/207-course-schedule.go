func canFinish(numCourses int, prerequisites [][]int) bool {
    
    indegrees := make([]int, numCourses)
    adjList := map[int][]int{}
    coursesTakenSuccessfully := 0
    
    for _, ele := range prerequisites {
        course := ele[0]
        prereq := ele[1]
        adjList[prereq] = append(adjList[prereq] , course)
        indegrees[course]++
    }
    
    q := []int{}
    for idx, count := range indegrees {
        if count == 0 {
            q = append(q, idx)
        }
    }
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        coursesTakenSuccessfully++
        dependencies := adjList[dq]
        for _,dep := range dependencies {
            indegrees[dep]--
            if indegrees[dep] == 0 {
                q = append(q, dep)
            }
        }
    }
    
    
    return coursesTakenSuccessfully == numCourses
    
}