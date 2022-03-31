/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */


type sol struct {
    empMap map[int]*Employee
}

func getImportance(employees []*Employee, id int) int {
    
    sol := &sol{empMap : map[int]*Employee{}}
    for _, emp := range employees {
        sol.empMap[emp.Id] = emp
    }    
    
    return sol.dfs(id)
}

func (s *sol) dfs(id int) int {
    
    // base
    
    
    // logic
    emp := s.empMap[id]
    imp := emp.Importance
    
    for _, sub := range emp.Subordinates {
        imp += s.dfs(s.empMap[sub].Id)
    }
    return imp
    
}