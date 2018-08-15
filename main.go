package main

import (
	"github.com/chonla/grel/github"
	"github.com/kr/pretty"
)

func main() {

	gh := github.NewGitHub("chonla/cotton")
	tags := gh.Tags()
	pretty.Println(tags)

}
