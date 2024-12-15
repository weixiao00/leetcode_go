package algorithm

import (
	"fmt"
	"testing"
)

func destCity(paths [][]string) string {

	if len(paths) == 0 {
		return ""
	}

	destCity := make([]string, 0)
	mp := make(map[string]string)
	for _, path := range paths {
		mp[path[0]] = path[1]
		destCity = append(destCity, path[1])
	}

	for _, city := range destCity {
		_, ok := mp[city]
		if !ok {
			return city
		}
	}

	return ""
}

func Test1436(t *testing.T) {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	city := destCity(paths)
	fmt.Println(city)
}
