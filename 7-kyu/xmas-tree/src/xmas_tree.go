package src

import (
	"strings"
)

func XMasTree(height int) []string {
	if height == 0 {
		return []string{}
	}

	tree := make([][]string, height+2)
	for i := range tree {
		tree[i] = strings.Split(strings.Repeat("_", height*2-1), "")
	}

	for i := 0; i < len(tree); i++ {
		if i > height-1 {
			tree[i][height-1] = "#"
			continue
		}
		for j := height - 1 - i; j <= height-1+i; j++ {
			tree[i][j] = "#"
		}
	}

	result := []string{}
	for i := 0; i < len(tree); i++ {
		result = append(result, strings.Join(tree[i], ""))
	}

	return result
}
