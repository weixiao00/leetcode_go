package algorithm

import "testing"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// dfs
func getImportance(employees []*Employee, id int) int {

	idMap := make(map[int]*Employee)

	for _, employ := range employees {
		idMap[employ.Id] = employ
	}

	total := 0
	var dfs func(id int)
	dfs = func(id int) {
		employee := idMap[id]
		total += employee.Importance
		for _, subId := range employee.Subordinates {
			dfs(subId)
		}
	}

	dfs(id)
	return total
}

// bfs
func getImportance1(employees []*Employee, id int) int {

	idMap := make(map[int]*Employee)

	for _, employ := range employees {
		idMap[employ.Id] = employ
	}

	total := 0
	queue := make([]*Employee, 0)
	queue = append(queue, idMap[id])
	var bfs func(id int)
	bfs = func(id int) {
		for len(queue) > 0 {
			employ := queue[0]
			queue = queue[1:]
			total += employ.Importance
			for _, subId := range employ.Subordinates {
				queue = append(queue, idMap[subId])
			}
		}
	}

	bfs(id)
	return total
}

func Test690(t *testing.T) {

}
