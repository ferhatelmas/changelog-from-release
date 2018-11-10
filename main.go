package main

import (
	"fmt"
	"os"
)

func help(code int) {
	fmt.Println(`Usage: changelog-from-release`)
	os.Exit(111)
}

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(111)
	}
	os.Exit(0)
}

func main() {
	if len(os.Args) != 1 {
		help(111)
	}

	git, err := NewGitForCwd()
	if err != nil {
		exit(err)
	}

	url, err := git.TrackingRemoteURL()
	if err != nil {
		exit(err)
	}

	gh, err := GitHubFromURL(url)
	if err != nil {
		exit(err)
	}

	rels, err := gh.Releases()
	if err != nil {
		exit(err)
	}

	cl, err := NewChangeLog(git.root, url)
	if err != nil {
		exit(err)
	}

	exit(cl.Generate(rels))
}