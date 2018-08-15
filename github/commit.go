package github

import (
	"sort"
	"time"
)

// Commit contains commit information
type Commit struct {
	Ref  string
	Date time.Time
}

// Commits is a list of commit
type Commits []Commit

func (c Commits) Len() int {
	return len(c)
}

func (c Commits) Less(i, j int) bool {
	return c[i].Date.Before(c[j].Date)
}

func (c Commits) Swap(i, j int) {
	ct := c[i]
	c[i] = c[j]
	c[j] = ct
}

// RevertSort to sort in different order
func (c Commits) RevertSort() {
	sort.Sort(sort.Reverse(c))
}
