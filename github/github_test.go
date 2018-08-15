package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitHub(t *testing.T) {
	gh := NewGitHub("chonla/cotton")

	assert.Equal(t, gh.URL, "https://github.com/chonla/cotton")
}
