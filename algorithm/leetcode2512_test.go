package algorithm

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {

	if len(student_id) == 0 || k == 0 {
		return []int{}
	}

	words := map[string]int{}
	for _, w := range positive_feedback {
		words[w] = 3
	}
	for _, w := range negative_feedback {
		words[w] = -1
	}

	scoreIds := make([][]int, 0)
	for i := range student_id {
		reportArr := strings.Split(report[i], " ")
		score := 0
		for _, r := range reportArr {
			score += words[r]
		}
		scoreIds = append(scoreIds, []int{score, student_id[i]})
	}

	sort.Slice(scoreIds, func(i, j int) bool {
		return scoreIds[i][0] > scoreIds[j][0] || scoreIds[i][0] == scoreIds[j][0] && scoreIds[i][1] < scoreIds[j][1]
	})

	res := make([]int, 0)
	for _, ids := range scoreIds {
		res = append(res, ids[1])
		if len(res) == k {
			return res
		}
	}

	return res
}

func Test2512(t *testing.T) {
	positive_feedback := []string{"smart", "brilliant", "studious"}
	negative_feedback := []string{"not"}
	report := []string{"this student is not studious", "the student is smart"}
	student_id := []int{1, 2}
	k := 2
	res := topStudents(positive_feedback, negative_feedback, report, student_id, k)
	fmt.Println(res)
}
