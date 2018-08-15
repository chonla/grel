package github

import "sort"

// Tag contains tag information
type Tag struct {
	Name        string
	Ref         string
	CommitIndex int
}

// Tags is a list of tag
type Tags []Tag

func (t Tags) Len() int {
	return len(t)
}

func (t Tags) Less(i, j int) bool {
	return t[i].CommitIndex < t[j].CommitIndex
}

func (t Tags) Swap(i, j int) {
	tt := t[i]
	t[i] = t[j]
	t[j] = tt
}

// Sort is convenient tool to sort Tags
func (t Tags) Sort() {
	sort.Sort(t)
}

// RevertSort to sort in different order
func (t Tags) RevertSort() {
	sort.Sort(sort.Reverse(t))
}
