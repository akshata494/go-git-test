// Akshata's test code to start using the go-git package.
// Also, test code for git pulling non fast forward updates.

package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	fmt.Println("Testing pull with go-git.")

	r, err := git.PlainOpen("C:\\Users\\akrishnan\\go\\src\\github.com\\akshata494\\go-merge-test\\git-merge-test")
	if err != nil {
		fmt.Printf("PlainOpen error : %s.\n", err)
		os.Exit(1)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		fmt.Printf("Retrieving worktree error : %s.\n", err)
		os.Exit(1)
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	fmt.Println("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		fmt.Printf("Pull error : %s.\n", err)
		os.Exit(1)
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		fmt.Printf("Fetching HEAD error : %s.\n", err)
		os.Exit(1)
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		fmt.Printf("Retrieving commit object error : %s.\n", err)
		os.Exit(1)
	}

	fmt.Println(commit)
}
