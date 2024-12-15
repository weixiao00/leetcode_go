package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// hash+并查集
func accountsMerge(accounts [][]string) [][]string {
	if len(accounts) == 0 {
		return accounts
	}

	mailIndexMap := make(map[string]int)
	mailNameMap := make(map[string]string)
	// 每个邮箱对应着一个唯一索引
	count := 0
	for _, accountMails := range accounts {
		name := accountMails[0]
		for i := 1; i < len(accountMails); i++ {
			mailNameMap[accountMails[i]] = name
			_, ok := mailIndexMap[accountMails[i]]
			if ok {
				continue
			}
			mailIndexMap[accountMails[i]] = count
			count++
		}
	}

	parent := make([]int, count)
	// 每个节点的父节点初始化为自己
	for i := 0; i < count; i++ {
		parent[i] = i
	}
	var findParent func(index int) int
	findParent = func(index int) int {
		if index != parent[index] {
			parent[index] = findParent(parent[index])
		}
		return parent[index]
	}
	var union func(idx1, idx2 int)
	union = func(idx1, idx2 int) {
		parent[findParent(idx2)] = findParent(idx1)
	}

	for _, accountMails := range accounts {
		if len(accountMails) < 2 {
			continue
		}
		firstIndex := mailIndexMap[accountMails[1]]
		for i := 2; i < len(accountMails); i++ {
			// 并查集关联
			nextIndex := mailIndexMap[accountMails[i]]
			union(firstIndex, nextIndex)
		}
	}

	indexMailsMap := make(map[int][]string)
	for mail := range mailIndexMap {
		index := mailIndexMap[mail]
		parentIndex := findParent(index)
		mails, ok := indexMailsMap[parentIndex]
		if !ok {
			mails = make([]string, 0)
		}
		mails = append(mails, mail)
		indexMailsMap[parentIndex] = mails
	}

	res := make([][]string, 0)
	for _, mails := range indexMailsMap {
		mailArr := make([]string, 0)
		name := mailNameMap[mails[0]]
		mailArr = append(mailArr, name)
		sort.Strings(mails)
		for _, mail := range mails {
			mailArr = append(mailArr, mail)
		}
		res = append(res, mailArr)
	}

	return res
}

func Test721(t *testing.T) {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"}}
	res := accountsMerge(accounts)
	//[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
	fmt.Println(res)
}
