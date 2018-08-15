/*
Package github provides a wrapper to github.
*/
package github

import (
	"fmt"
	"sort"

	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// GitHub handler
type GitHub struct {
	URL             string
	repo            *git.Repository
	commits         Commits
	invertedCommits map[string]int
	tags            Tags
}

// NewGitHub to create a new github handler from given repo
func NewGitHub(repo string) *GitHub {
	url := fmt.Sprintf("https://github.com/%s", repo)
	gh := &GitHub{
		URL:             url,
		commits:         []Commit{},
		invertedCommits: map[string]int{},
		tags:            []Tag{},
	}
	gh.loadRepository()
	return gh
}

func (gh *GitHub) loadRepository() error {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: gh.URL,
	})
	if err != nil {
		return err
	}
	gh.repo = r

	ref, err := r.Head()

	// Load commits
	commitIter, err := gh.repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return err
	}

	err = commitIter.ForEach(func(c *object.Commit) error {
		gh.commits = append(gh.commits, Commit{
			Ref:  c.Hash.String(),
			Date: c.Author.When,
		})
		return nil
	})

	sort.Stable(gh.commits)
	for i, c := range gh.commits {
		gh.invertedCommits[c.Ref] = i
	}

	if err != nil {
		return err
	}

	// Load tags
	iter, err := gh.repo.Tags()
	if err != nil {
		return err
	}

	err = iter.ForEach(func(ref *plumbing.Reference) error {
		gh.tags = append(gh.tags, Tag{
			Name:        ref.Name().String(),
			Ref:         ref.Hash().String(),
			CommitIndex: gh.invertedCommits[ref.Hash().String()],
		})
		return nil
	})

	gh.tags.RevertSort()

	return nil
}

// Commits return list of commits ordered by date descendingly
func (gh *GitHub) Commits() Commits {
	return gh.commits
}

// Tags return all tags
func (gh *GitHub) Tags() Tags {
	return gh.tags
}
