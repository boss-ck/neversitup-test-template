package main

import (
	"slices"
)

type Permute struct {
	nums    []rune
	current []rune
	res     *[]string
	used    []bool
}

func Manipulate(text string) []string {
	// TODO : start your code here
	var res []string
	rword := []rune(text)

	// sorting for clean dup
	slices.Sort(rword)

	// backtrack
	permute(&Permute{
		nums:    rword,
		current: []rune{},
		res:     &res,
		used:    make([]bool, len(rword)),
	})

	return res
}

func permute(p *Permute) {
	// base case
	if len(p.current) == len(p.nums) {
		*p.res = append(*p.res, string(p.current))
		return
	}

	for i := 0; i < len(p.nums); i++ {
		// skip used or dup
		if p.used[i] || (i > 0 && p.nums[i] == p.nums[i-1] && !p.used[i-1]) {
			continue
		}

		// set used
		p.used[i] = true
		p.current = append(p.current, p.nums[i])

		// backtrack
		permute(p)

		// reset used
		p.used[i] = false
		p.current = p.current[:len(p.current)-1]
	}
}
